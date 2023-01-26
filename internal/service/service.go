package service

import (
	"golang_vast/internal/model"
	"golang_vast/internal/repository"
)

type Vast interface {
	PutVastInfo(ad model.VastModel) (model.VastModel, error)
	Generate_Vast(ad model.VastModel) error
	GenerateVMAP() error
}

type Service struct {
	Vast
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Vast: NewVastService(repo.VastRepository),
	}
}
