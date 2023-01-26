package service

import (
	"golang_vast/internal/model"
	"golang_vast/internal/repository"
)

type Vast interface {
	PutVastInfo(ad model.VastModel) (string, error)
	Generate_Vast(id string) error
}

type Service struct {
	Vast
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Vast: NewVastService(repo.VastRepository),
	}
}
