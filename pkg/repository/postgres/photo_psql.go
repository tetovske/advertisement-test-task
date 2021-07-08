package postgres

import (
	"database/sql"
	"github.com/tetovske/advertisement-service/pkg/models"
)

type PhotoPSQL struct {
	conn *sql.DB
}

func NewPhotoPSQL(conn *sql.DB) *PhotoPSQL {
	return &PhotoPSQL{conn: conn}
}

func (r *PhotoPSQL) CreatePhoto(pic models.Photo) (int, error) {
	return 0, nil
}

func (r *PhotoPSQL) GetPhoto(id int) (models.Photo, error) {
	return models.Photo{
		Id:   0,
		Link: "a",
		Tag:  0,
	}, nil
}

func (r *PhotoPSQL) GetPhotoList(id int) ([]models.Photo, error) {
	return []models.Photo{
		{
			Id:   0,
			Link: "a",
			Tag:  0,
		},
	}, nil
}
