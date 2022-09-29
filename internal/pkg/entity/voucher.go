package entity

import (
	voucherApi "github.com/h-varmazyar/voucher/api/proto"
	"github.com/h-varmazyar/voucher/internal/pkg/db"
	"time"
)

type Voucher struct {
	db.UniversalModel
	Code           string                 `json:"code"`
	Description    string                 `json:"description"`
	MaxAmount      int64                  `json:"max_amount"`
	UsageLimit     int64                  `json:"usage_limit"`
	Discount       int32                  `json:"discount"`
	StartTime      time.Time              `json:"start_time"`
	ExpirationTime time.Time              `json:"expiration_time"`
	Amount         int64                  `json:"amount"`
	Type           voucherApi.VoucherType `json:"type"`
}
