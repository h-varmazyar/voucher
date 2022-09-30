package entity

import (
	voucherApi "github.com/h-varmazyar/voucher/api/proto"
	"github.com/h-varmazyar/voucher/internal/pkg/db"
	"time"
)

type Voucher struct {
	db.UniversalModel
	Code           string                 `json:"code" gorm:"not null"`
	Description    string                 `json:"description" gorm:"not null"`
	Amount         int64                  `json:"amount"`
	MaxAmount      int64                  `json:"max_amount"`
	UsageLimit     int64                  `json:"usage_limit"`
	Discount       int32                  `json:"discount"`
	StartTime      time.Time              `json:"start_time" gorm:"not null"`
	ExpirationTime time.Time              `json:"expiration_time" gorm:"not null"`
	Type           voucherApi.VoucherType `json:"type"  gorm:"type:varchar(25);not null"`
}
