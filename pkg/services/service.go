package services

import (
	"github.com/tetovske/advertisement-service/pkg/models"
	"github.com/tetovske/advertisement-service/pkg/repository"
)

type Advertisement interface {
	CreateAdvertisement(ad models.Advertisement) (int, error)
	GetAdvertisement(id int, fields []string) (map[string]interface{}, error)
	GetAdvertisements(sort string) ([]map[string]interface{}, error)
}

type Service struct {
	Advertisement
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Advertisement: NewAdvertisementService(repo),
	}
}
