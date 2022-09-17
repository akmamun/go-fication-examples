package repository

import (
	"go-fication-examples/infra/database"
)

type GormRepository struct {
	db *database.DB
}

func NewGormRepository(db *database.DB) *GormRepository {
	return &GormRepository{
		db: db,
	}
}
