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

func (r *Vast_Repos) AdVast(ad model.VastModel) (string, error) {
	var id string

	query := fmt.Sprintf("insert into vast ( title, is_Skipable, skip_time, ads_height, ads_width, adsDuration, uri) values ($1, $2, $3, $4, $5, $6, $7) returning id")

	row := r.db.QueryRow(query, ad.Title, ad.IsSkipable, ad.SkipTime, ad.AdsHeight, ad.AdsWidth, ad.AdsDuration, ad.MediaURI)

	if err := row.Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (r *Vast_Repos) GetById(id string) (model.VastModel, error) {

	var vast model.VastModel
	rows, err := r.db.Query("select title, is_Skipable, skip_time, ads_height, ads_width, adsDuration, uri from vast where id=$1", id)
	if err != nil {
		return model.VastModel{}, err
	}

	for rows.Next() {
		if err := rows.Scan(&vast.Title, &vast.IsSkipable, &vast.SkipTime, &vast.AdsHeight, &vast.AdsWidth, &vast.AdsDuration, &vast.MediaURI); err != nil {
			return model.VastModel{}, err
		}
	}

	defer rows.Close()

	return vast, nil
}
