package repositories

import (
	"database/sql"
	"rachao/internal/core/domain"
)

type PositionRepository struct {
	DB *sql.DB
}

const GetPositionAllQuery = `SELECT * FROM position ORDER BY name DESC;`

func (repo *PositionRepository) GetAll() ([]domain.Position, error) {
	rows, err := repo.DB.Query(GetPositionAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var positions []domain.Position
	for rows.Next() {
		var position domain.Position
		if err := rows.Scan(&position.ID, &position.Name, &position.Acronym); err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return positions, nil
}

const CreatePositionQuery = `INSERT INTO position (name, acronym) VALUES ($1, $2) RETURNING id;`

func (repo *PositionRepository) Create(position domain.CreatePositionRequest) (int, error) {
	var id int
	err := repo.DB.QueryRow(CreatePositionQuery, position.Name, position.Acronym).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

const GetPositionByIDQuery = `SELECT * FROM position WHERE id = $1;`

func (repo *PositionRepository) GetByID(id int) (domain.Position, error) {
	row := repo.DB.QueryRow(GetPositionByIDQuery, id)
	var position domain.Position
	if err := row.Scan(&position.ID, &position.Name, &position.Acronym); err != nil {
		if err == sql.ErrNoRows {
			return position, nil
		}
		return position, err
	}
	return position, nil
}

const UpdatePositionQuery = `UPDATE position SET name = $1, acronym = $2 WHERE id = $3;`

func (repo *PositionRepository) Update(id int, position domain.Position) error {
	_, err := repo.DB.Exec(UpdatePositionQuery, position.Name, position.Acronym, id)
	if err != nil {
		return err
	}
	return nil
}

const DeletePositionQuery = `DELETE FROM position WHERE id = $1;`

func (repo *PositionRepository) Delete(id int) error {
	_, err := repo.DB.Exec(DeletePositionQuery, id)
	if err != nil {
		return err
	}
	return nil
}
