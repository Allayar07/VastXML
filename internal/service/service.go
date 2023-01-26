package service

import (
	"golang_vast/internal/repository"
)

type Service struct {
	Vast *VastService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Vast: NewVastService(repo.Vast),
	}
}
