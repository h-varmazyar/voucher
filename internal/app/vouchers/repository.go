package vouchers

import (
	"context"
	"github.com/google/uuid"
	"github.com/h-varmazyar/voucher/internal/pkg/db"
	"github.com/h-varmazyar/voucher/internal/pkg/entity"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, transaction *entity.Voucher) (*entity.Voucher, error)
	ReturnByID(ctx context.Context, voucherID uuid.UUID) (*entity.Voucher, error)
	ReturnByCode(ctx context.Context, code string) (*entity.Voucher, error)
	IsUsedBefore(ctx context.Context, voucherID uuid.UUID, phoneNumber string) (bool, error)
	UsageCount(ctx context.Context, voucherID uuid.UUID) (int64, error)
	UsageList(ctx context.Context, voucherID uuid.UUID) ([]*entity.VoucherUsage, error)
	AddUsage(ctx context.Context, usage *entity.VoucherUsage) (*entity.VoucherUsage, error)
	DeleteUsage(ctx context.Context, usageID uuid.UUID, hardDelete bool) error
}

type repository struct {
	logger *log.Logger
	db     *db.DB
}

func NewRepository(db *db.DB, logger *log.Logger) Repository {
	return &repository{logger, db}
}

func (r *repository) Create(_ context.Context, voucher *entity.Voucher) (*entity.Voucher, error) {
	//TODO: validate voucher
	if err := r.db.Save(voucher).Error; err != nil {
		return nil, err
	}
	return voucher, nil
}

func (r *repository) ReturnByID(_ context.Context, voucherID uuid.UUID) (*entity.Voucher, error) {
	voucher := new(entity.Voucher)
	if err := r.db.
		Model(new(entity.Voucher)).
		Where("id = ?", voucherID).
		First(voucher).Error; err != nil {
		return nil, err
	}
	return voucher, nil
}

func (r *repository) ReturnByCode(_ context.Context, code string) (*entity.Voucher, error) {
	voucher := new(entity.Voucher)
	if err := r.db.
		Model(new(entity.Voucher)).
		Where("code = ?", code).
		First(voucher).Error; err != nil {
		return nil, err
	}
	return voucher, nil
}

func (r *repository) IsUsedBefore(_ context.Context, voucherID uuid.UUID, phoneNumber string) (bool, error) {
	count := int64(0)
	if err := r.db.
		Model(new(entity.VoucherUsage)).
		Where("voucher_id = ?", voucherID).
		Where("phone_number = ?", phoneNumber).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repository) UsageCount(_ context.Context, voucherID uuid.UUID) (int64, error) {
	count := int64(0)
	if err := r.db.
		Model(new(entity.VoucherUsage)).
		Where("voucher_id = ?", voucherID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *repository) AddUsage(_ context.Context, usage *entity.VoucherUsage) (*entity.VoucherUsage, error) {
	if err := r.db.Save(usage).Error; err != nil {
		return nil, err
	}
	return usage, nil
}

func (r *repository) DeleteUsage(ctx context.Context, usageID uuid.UUID, hardDelete bool) error {
	var tx *gorm.DB
	if hardDelete {
		tx = r.db.Unscoped()
	} else {
		tx = r.db.DB
	}
	if err := tx.Model(new(entity.VoucherUsage)).Delete(new(entity.VoucherUsage), usageID).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) UsageList(_ context.Context, voucherID uuid.UUID) ([]*entity.VoucherUsage, error) {
	transactions := make([]*entity.VoucherUsage, 0)
	if err := r.db.
		Model(new(entity.VoucherUsage)).
		Where("voucher_id = ?", voucherID).
		Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
