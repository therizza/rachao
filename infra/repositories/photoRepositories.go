package repositories

import (
	"database/sql"
	"rachao/internal/core/domain"

	"github.com/google/uuid"
)

type PhotoRepository struct {
	DB *sql.DB
}

const CreatePhotoQuery = `INSERT INTO photo (id_play, photo) VALUES ($1, $2) RETURNING id;`

func (repo *PhotoRepository) Create(idPlay uuid.UUID, photo []byte) (uuid.UUID, error) {
	var id uuid.UUID
	err := repo.DB.QueryRow(CreatePhotoQuery, idPlay, photo).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

const GetPhotoByIDPlayQuery = `SELECT * FROM photo WHERE id_play = $1;`

func (repo *PhotoRepository) GetByIDPlay(idPlay uuid.UUID) ([]domain.Photo, error) {
	rows, err := repo.DB.Query(GetPhotoByIDPlayQuery, idPlay)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []domain.Photo
	for rows.Next() {
		var photo domain.Photo
		if err := rows.Scan(&photo.ID, &photo.IDPlay, &photo.Photo); err != nil {
			return nil, err
		}
		photos = append(photos, photo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return photos, nil
}

const UpdatePhotoQuery = `UPDATE photo SET photo = $1 WHERE id_play = $2;`

func (repo *PhotoRepository) Update(idPlay uuid.UUID, photo []byte) error {
	_, err := repo.DB.Exec(UpdatePhotoQuery, photo, idPlay)
	if err != nil {
		return err
	}
	return nil
}

const DeletePhotoQuery = `DELETE FROM photo WHERE id_play = $1;`

func (repo *PhotoRepository) Delete(idPlay uuid.UUID) error {
	_, err := repo.DB.Exec(DeletePhotoQuery, idPlay)
	if err != nil {
		return err
	}
	return nil
}
