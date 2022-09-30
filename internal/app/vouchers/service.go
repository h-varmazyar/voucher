package vouchers

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	voucherApi "github.com/h-varmazyar/voucher/api/proto"
	"github.com/h-varmazyar/voucher/internal/pkg/db"
	"github.com/h-varmazyar/voucher/internal/pkg/entity"
	"github.com/h-varmazyar/voucher/pkg/grpcext"
	"github.com/h-varmazyar/voucher/pkg/mapper"
	walletApi "github.com/h-varmazyar/wallet/api/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"strings"
)

type Service struct {
	voucherApi.UnimplementedVoucherServiceServer
	repository    Repository
	walletService walletApi.WalletServiceClient
	configs       *Configs
	logger        *log.Logger
}

func NewService(configs *Configs, db *db.DB, log *log.Logger) *Service {
	walletConn := grpcext.NewConnection(configs.WalletServiceAddress)
	service := &Service{
		repository:    NewRepository(db, log),
		walletService: walletApi.NewWalletServiceClient(walletConn),
		configs:       configs,
		logger:        log,
	}

	return service
}

func (s *Service) RegisterServer(server *grpc.Server) {
	voucherApi.RegisterVoucherServiceServer(server, s)
}

func (s *Service) Create(ctx context.Context, req *voucherApi.VoucherCreateReq) (*voucherApi.Voucher, error) {
	voucher := new(entity.Voucher)
	mapper.Struct(req, voucher)

	if tx, err := s.repository.Create(ctx, voucher); err != nil {
		return nil, err
	} else {
		response := new(voucherApi.Voucher)
		mapper.Struct(tx, response)
		return response, nil
	}
}

func (s *Service) Apply(ctx context.Context, req *voucherApi.VoucherApplyReq) (*voucherApi.Void, error) {
	var (
		err     error
		voucher *entity.Voucher
	)

	if voucher, err = s.repository.ReturnByCode(ctx, req.Code); err != nil {
		return nil, err
	}

	if err = s.checkVoucherUsage(ctx, voucher, req.PhoneNumber); err != nil {
		return nil, err
	}

	switch voucher.Type {
	case voucherApi.Voucher_Voucher:
		return nil, errors.New("unimplemented")
	case voucherApi.Voucher_Credit:
		err = s.creditAllocation(ctx, voucher, req.PhoneNumber)
	}
	if err != nil {
		return nil, err
	}

	return new(voucherApi.Void), nil
}

func (s *Service) Usage(ctx context.Context, req *voucherApi.VoucherUsageReq) (*voucherApi.Usages, error) {
	voucher, err := s.prepareVoucherForUsage(ctx, req)
	if err != nil {
		return nil, err
	}

	var usages []*entity.VoucherUsage
	if usages, err = s.repository.UsageList(ctx, voucher.ID); err != nil {
		return nil, err
	}

	response := new(voucherApi.Usages)
	phoneNumber := make([]string, 0)
	for _, usage := range usages {
		phoneNumber = append(phoneNumber, usage.PhoneNumber)
	}
	mapper.Slice(phoneNumber, &response.PhoneNumbers)

	respVoucher := new(voucherApi.Voucher)
	mapper.Struct(voucher, respVoucher)
	response.Voucher = respVoucher
	response.Count = int64(len(phoneNumber))

	return response, nil
}

func (s *Service) checkVoucherUsage(ctx context.Context, voucher *entity.Voucher, phoneNumber string) error {
	var (
		used       bool
		usageCount int64
		err        error
	)

	if usageCount, err = s.repository.UsageCount(ctx, voucher.ID); err != nil {
		return err
	}
	if usageCount >= voucher.UsageLimit {
		return errors.New("usage limit exceed")
	}

	if used, err = s.repository.IsUsedBefore(ctx, voucher.ID, phoneNumber); err != nil {
		return err
	}
	if used {
		return errors.New("voucher code used before")
	}
	return nil
}

func (s *Service) creditAllocation(ctx context.Context, voucher *entity.Voucher, phoneNumber string) error {
	var (
		err    error
		wallet *walletApi.Wallet
		usage  *entity.VoucherUsage
	)
	if wallet, err = s.fetchWallet(ctx, phoneNumber); err != nil {
		return err
	}

	if usage, err = s.repository.AddUsage(ctx, &entity.VoucherUsage{
		VoucherID:   voucher.ID,
		PhoneNumber: phoneNumber,
	}); err != nil {
		return err
	}

	if err = s.chargeWallet(ctx, voucher, wallet.ID, usage.ID); err != nil {
		return err
	}

	return nil
}

func (s *Service) fetchWallet(ctx context.Context, phoneNumber string) (*walletApi.Wallet, error) {
	var (
		err    error
		wallet *walletApi.Wallet
	)
	if wallet, err = s.walletService.ReturnByPhoneNumber(ctx, &walletApi.WalletReturnByPhoneReq{
		PhoneNumber: phoneNumber,
	}); err != nil {
		if strings.Contains(err.Error(), "record not found") {
			if wallet, err = s.walletService.Create(ctx, &walletApi.WalletCreateReq{
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

func (s *Service) chargeWallet(ctx context.Context, voucher *entity.Voucher, walletID string, usageID uuid.UUID) error {
	if _, err := s.walletService.Deposit(ctx, &walletApi.NewTransaction{
		WalletID:    walletID,
		Amount:      voucher.Amount,
		RefID:       usageID.String(),
		Description: fmt.Sprintf("add credit of voucher %v", voucher.Code),
	}); err != nil {
		s.logger.WithError(err).Errorf("failed to deposite wallet %v for voucher %v", walletID, voucher.Code)
		if rollbackErr := s.repository.DeleteUsage(ctx, usageID, true); rollbackErr != nil {
			s.logger.WithError(rollbackErr).Errorf("failed to rollback uage: %v", usageID)
		}
		return err
	}
	return nil
}

func (s *Service) prepareVoucherForUsage(ctx context.Context, req *voucherApi.VoucherUsageReq) (*entity.Voucher, error) {
	var (
		voucher *entity.Voucher
		err     error
	)

	switch req.Identifier.(type) {
	case *voucherApi.VoucherUsageReq_Code:
		voucher, err = s.repository.ReturnByCode(ctx, req.GetCode())
	case *voucherApi.VoucherUsageReq_ID:
		var voucherID uuid.UUID
		voucherID, err = uuid.Parse(req.GetID())
		if err != nil {
			return nil, err
		}
		voucher, err = s.repository.ReturnByID(ctx, voucherID)
	default:
		err = errors.New("invalid voucher identifier")
	}
	if err != nil {
		return nil, err
	}
	return voucher, nil
}
