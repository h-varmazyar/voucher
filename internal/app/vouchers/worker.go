package vouchers

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	voucherApi "github.com/h-varmazyar/voucher/api/proto"
	"github.com/h-varmazyar/voucher/internal/pkg/entity"
	"github.com/h-varmazyar/voucher/pkg/grpcext"
	walletApi "github.com/h-varmazyar/wallet/api/proto"
	log "github.com/sirupsen/logrus"
	"strings"
)

type ApplyWorker struct {
	dataChan      chan *workerSeed
	repository    Repository
	walletService walletApi.WalletServiceClient
	logger        *log.Logger
}

type workerSeed struct {
	ctx         context.Context
	voucher     *entity.Voucher
	phoneNumber string
	respChan    chan *WorkerResp
}

type WorkerResp struct {
	isDone bool
	Error  error
}

func NewWorker(configs *Configs, r Repository) *ApplyWorker {
	worker := new(ApplyWorker)

	walletConn := grpcext.NewConnection(configs.WalletServiceAddress)
	worker.walletService = walletApi.NewWalletServiceClient(walletConn)
	worker.repository = r
	worker.dataChan = make(chan *workerSeed)

	return worker
}

func (w *ApplyWorker) Start() {
	for {
		seed, ok := <-w.dataChan
		if !ok {
			break
		}

		response := w.consumeSeed(seed)

		select {
		case <-seed.ctx.Done():
			return
		default:
			seed.respChan <- response
		}
	}
}

func (w *ApplyWorker) consumeSeed(seed *workerSeed) *WorkerResp {
	var (
		err error
	)

	if err = w.checkForUsage(seed.ctx, seed.voucher); err != nil {
		return &WorkerResp{
			Error: err,
		}
	}

	switch seed.voucher.Type {
	case voucherApi.Voucher_Voucher:
		err = errors.New("unimplemented")
	case voucherApi.Voucher_Credit:
		err = w.creditAllocation(seed.ctx, seed.voucher, seed.phoneNumber)
	}
	if err != nil {
		return &WorkerResp{
			Error: err,
		}
	}

	return &WorkerResp{
		isDone: true,
	}
}

func (w *ApplyWorker) checkForUsage(ctx context.Context, voucher *entity.Voucher) error {
	var (
		err        error
		usageCount int64
	)

	if usageCount, err = w.repository.UsageCount(ctx, voucher.ID); err != nil {
		return err
	}
	if usageCount >= voucher.UsageLimit {
		return errors.New("usage limit exceed")
	}
	return nil
}

func (w *ApplyWorker) creditAllocation(ctx context.Context, voucher *entity.Voucher, phoneNumber string) error {
	var (
		err    error
		wallet *walletApi.Wallet
		usage  *entity.VoucherUsage
	)
	if wallet, err = w.fetchWallet(ctx, phoneNumber); err != nil {
		return err
	}

	if usage, err = w.repository.AddUsage(ctx, &entity.VoucherUsage{
		VoucherID:   voucher.ID,
		PhoneNumber: phoneNumber,
	}); err != nil {
		return err
	}

	if err = w.chargeWallet(ctx, voucher, wallet.ID, usage.ID); err != nil {
		return err
	}

	return nil
}

func (w *ApplyWorker) fetchWallet(ctx context.Context, phoneNumber string) (*walletApi.Wallet, error) {
	var (
		err    error
		wallet *walletApi.Wallet
	)
	if wallet, err = w.walletService.ReturnByPhoneNumber(ctx, &walletApi.WalletReturnByPhoneReq{
		PhoneNumber: phoneNumber,
	}); err != nil {
		if strings.Contains(err.Error(), "record not found") {
			if wallet, err = w.walletService.Create(ctx, &walletApi.WalletCreateReq{
				PhoneNumber: phoneNumber,
				Amount:      0,
			}); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return wallet, nil
}

func (w *ApplyWorker) chargeWallet(ctx context.Context, voucher *entity.Voucher, walletID string, usageID uuid.UUID) error {
	if _, err := w.walletService.Deposit(ctx, &walletApi.NewTransaction{
		WalletID:    walletID,
		Amount:      voucher.Amount,
		RefID:       usageID.String(),
		Description: fmt.Sprintf("add credit of voucher %v", voucher.Code),
	}); err != nil {
		w.logger.WithError(err).Errorf("failed to deposite wallet %v for voucher %v", walletID, voucher.Code)
		if rollbackErr := w.repository.DeleteUsage(ctx, usageID, true); rollbackErr != nil {
			w.logger.WithError(rollbackErr).Errorf("failed to rollback uage: %v", usageID)
		}
		return err
	}
	return nil
}
