package repository

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepository struct {
	db *gorm.DB
}

func (b *BaseRepository) getDB(ctx context.Context) *gorm.DB {
	return b.db.WithContext(ctx)
}
