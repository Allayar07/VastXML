package repository

import (
	"github.com/jmoiron/sqlx"
	"golang_vast/internal/model"
)

type VastRepository interface {
	AdVast(ad model.VastModel) (model.VastModel, error)
	//GetById(id string) (model.VastModel, error)
}

type Repository struct {
	VastRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		VastRepository: NewVast_Repos(db),
	}
}
