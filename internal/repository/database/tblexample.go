package database

import (
	"context"
	"github.com/cynx-io/micro-name/internal/model/entity"
	"gorm.io/gorm"
)

type TblExample struct {
	DB *gorm.DB
}

func NewTblExample(DB *gorm.DB) *TblExample {
	return &TblExample{
		DB: DB,
	}
}

func (r *TblExample) GetExample(ctx context.Context, id int32) (*entity.TblExample, error) {
	return nil, nil
}
