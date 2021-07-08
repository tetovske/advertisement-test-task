package services

import "github.com/tetovske/advertisement-service/pkg/repository"

type AdvertisementService struct {
	repo repository.AdvertisementRepository
}

func NewAdvertisementService(repo repository.AdvertisementRepository) *AdvertisementService {
	return &AdvertisementService{repo: repo}
}

func (r *AdvertisementService) CreateAdvertisement() (int, error) {
	return 0, nil
}
