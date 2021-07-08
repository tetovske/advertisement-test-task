package repository

import (
	"database/sql"
	"github.com/tetovske/advertisement-service/pkg/models"
	"github.com/tetovske/advertisement-service/pkg/repository/postgres"
)

type AdvertisementRepository interface {
	CreateAdvertisement(ad models.Advertisement) (int, error)
	GetAdvertisement(id int) (models.Advertisement, error)
	GetAdvertisementList(id int) (models.Advertisement, error)
}

type PhotoRepository interface {
	CreatePhoto(pic models.Photo) (int, error)
	GetPhoto(id int) (models.Photo, error)
	GetPhotoList(id int) ([]models.Photo, error)
}

type Repository struct {
	AdvertisementRepository
	PhotoRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		AdvertisementRepository: postgres.NewAdvertisementPSQL(db),
		PhotoRepository: postgres.NewPhotoPSQL(db),
	}
}
