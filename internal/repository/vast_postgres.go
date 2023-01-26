package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"golang_vast/internal/model"
)

type VastPostgres struct {
	db *sqlx.DB
}

func NewVastPostgres(db *sqlx.DB) *VastPostgres {
	return &VastPostgres{
		db: db,
	}
}

func (r *VastPostgres) GetAd(id int) (model.VastModel, error) {
	var ad model.VastModel
	query := fmt.Sprintf(`SELECT id, title, video_path, ads_duration, ads_height, ads_width, is_skipable, skip_time FROM %s WHERE id = $1`, AdsTable)
	err := r.db.QueryRow(query, id).Scan(&ad.Id, &ad.Title, &ad.MediaURI, &ad.AdsDuration, &ad.AdsHeight, &ad.AdsWidth, &ad.IsSkippable, &ad.SkipTime)
	if err != nil {
		return ad, err
	}
	return ad, nil
}
