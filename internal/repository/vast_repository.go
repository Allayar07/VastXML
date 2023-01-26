package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"golang_vast/internal/model"
)

type Vast_Repos struct {
	db *sqlx.DB
}

func NewVast_Repos(db *sqlx.DB) *Vast_Repos {
	return &Vast_Repos{
		db: db,
	}
}

func (r *Vast_Repos) AdVast(ad model.VastModel) (model.VastModel, error) {

	query := fmt.Sprintf("insert into vast ( title, is_Skipable, skip_time, ads_height, ads_width, adsDuration, uri) values ($1, $2, $3, $4, $5, $6, $7) returning id, title, is_Skipable, skip_time, ads_height, ads_width, adsDuration, uri ")

	rows, err := r.db.Query(query, ad.Title, ad.IsSkipable, ad.SkipTime, ad.AdsHeight, ad.AdsWidth, ad.AdsDuration, ad.MediaURI)
	if err != nil {
		return model.VastModel{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&ad.ID, &ad.Title, &ad.IsSkipable, &ad.SkipTime, &ad.AdsHeight, &ad.AdsWidth, &ad.AdsDuration, &ad.MediaURI); err != nil {
			return model.VastModel{}, err
		}
	}

	defer rows.Close()

	return ad, nil
}
