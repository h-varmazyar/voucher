package entity

import (
	"github.com/google/uuid"
	"github.com/h-varmazyar/voucher/internal/pkg/db"
)

type VoucherUsage struct {
	db.UniversalModel
	VoucherID   uuid.UUID
	PhoneNumber string
}
