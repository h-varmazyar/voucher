package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func New(dsn string) (*DB, error) {
	dialector := postgres.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{FullSaveAssociations: false})
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
