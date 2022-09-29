package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UniversalModel struct {
	ID        uuid.UUID `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
