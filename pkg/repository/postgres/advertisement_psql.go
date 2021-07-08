package postgres

import (
	"database/sql"
	"github.com/tetovske/advertisement-service/pkg/models"
)

type AdvertisementPSQL struct {
	conn *sql.DB
}

func NewAdvertisementPSQL(conn *sql.DB) *AdvertisementPSQL {
	return &AdvertisementPSQL{conn: conn}
}

func (r *AdvertisementPSQL) CreateAdvertisement(ad models.Advertisement) (int, error) {
	return 0, nil
}

func (r *AdvertisementPSQL) GetAdvertisement(id int) (models.Advertisement, error) {
	return models.Advertisement{
		Id:              0,
		Title:          "a",
		Description: 	"sd",
	}, nil
}

func (r *AdvertisementPSQL) GetAdvertisementList(id int) (models.Advertisement, error) {
	return models.Advertisement{
		Id:          	0,
		Title:       	"a",
		Description: 	"sd",
	}, nil
}
