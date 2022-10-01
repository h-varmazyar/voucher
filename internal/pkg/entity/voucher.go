package entity

import (
	voucherApi "github.com/h-varmazyar/voucher/api/proto"
	"github.com/h-varmazyar/voucher/internal/pkg/db"
	"time"
)

type Voucher struct {
	db.UniversalModel
	Code           string                 `json:"code" gorm:"not null, unique" validate:"required"`
	Description    string                 `json:"description" gorm:"not null"`
	Amount         int64                  `json:"amount" validate:"gte=0"`
	MaxAmount      int64                  `json:"max_amount" validate:"gte=0"`
	UsageLimit     int64                  `json:"usage_limit" validate:"required,gt=0"`
	Discount       int32                  `json:"discount" validate:"gte=0,lte=100"`
	StartTime      time.Time              `json:"start_time" gorm:"not null" validate:"required"`
	ExpirationTime time.Time              `json:"expiration_time" gorm:"not null" validate:"required"`
	Type           voucherApi.VoucherType `json:"type"  gorm:"type:varchar(25);not null" validate:"required,enum"`
}
