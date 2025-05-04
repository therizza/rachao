package repositories

import (
	"database/sql"
	"rachao/internal/core/domain"

	"github.com/google/uuid"
)

type OverallRepository struct {
	DB *sql.DB
}

const ExistsOverallQuery = `SELECT EXISTS(SELECT 1 FROM overall WHERE id_play = $1);`

func (repo *OverallRepository) Exists(idPlay uuid.UUID) (bool, error) {
	var exists bool
	err := repo.DB.QueryRow(ExistsOverallQuery, idPlay).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

const GetByIDPlayQuery = `SELECT * FROM overall WHERE id_play = $1;`

func (repo *OverallRepository) GetByIDPlay(idUser uuid.UUID) (overall domain.Overall, err error) {
	rows, err := repo.DB.Query(GetByIDPlayQuery, idUser)
	if err != nil {
		return overall, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&overall.ID,
			&overall.IDPlay,
			&overall.Overall,
		); err != nil {
			return overall, err
		}
	}
	if err := rows.Err(); err != nil {
		return overall, err
	}

	return overall, nil
}

const CreateOverallQuery = `INSERT INTO overall (id, id_play, overall) VALUES ($1, $2, $3) RETURNING id;`

func (repo *OverallRepository) Create(overall domain.OverallRequest) (uuid.UUID, error) {
	id := uuid.New()
	err := repo.DB.QueryRow(CreateOverallQuery, id, overall.IDPlay, overall.Overall).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

const UpdateOverallQuery = `UPDATE overall SET overall = $1 WHERE id_play = $2;`

func (repo *OverallRepository) Update(overall domain.OverallRequest, idUser uuid.UUID) error {
	_, err := repo.DB.Exec(UpdateOverallQuery, overall.Overall, idUser)
	if err != nil {
		return err
	}
	return nil
}

const DeleteOverallQuery = `DELETE FROM overall WHERE id_user = $1;`

func (repo *OverallRepository) Delete(idUser uuid.UUID) error {
	_, err := repo.DB.Exec(DeleteOverallQuery, idUser)
	if err != nil {
		return err
	}
	return nil
}
