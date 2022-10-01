package entity

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/voucher/internal/pkg/db"
)

type VoucherUsage struct {
	db.UniversalModel
	VoucherID   uuid.UUID `json:"voucher_id" gorm:"type:uuid REFERENCES vouchers(id)" validate:"required"`
	PhoneNumber string    `json:"phone_number" validate:"required,mobile"`
}
