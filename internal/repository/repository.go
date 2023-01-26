package repository

import (
	"github.com/jmoiron/sqlx"
)

//type VastRepository interface {
//	AdVast(ad model.VastModel) (model.VastModel, error)
//	//GetById(id string) (model.VastModel, error)
//}

type Repository struct {
	Vast *VastPostgres
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Vast: NewVastPostgres(db),
	}
}
