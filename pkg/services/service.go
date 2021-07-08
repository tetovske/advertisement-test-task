package services

import "github.com/tetovske/advertisement-service/pkg/repository"

type Advertisement interface {
	CreateAdvertisement() (int, error)
}

type Service struct {
	Advertisement
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Advertisement: NewAdvertisementService(repo.AdvertisementRepository),
	}
}
