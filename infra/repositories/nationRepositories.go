package repositories

import (
	"database/sql"
	"rachao/internal/core/domain"
)

type NationRepository struct {
	DB *sql.DB
}

const GetNationAllQuery = `SELECT * FROM nation;`

func (repo *NationRepository) GetAll() ([]domain.Nation, error) {
	rows, err := repo.DB.Query(GetNationAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nations []domain.Nation
	for rows.Next() {
		var nation domain.Nation
		if err := rows.Scan(&nation.ID, &nation.Name, &nation.Acronym); err != nil {
			return nil, err
		}
		nations = append(nations, nation)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return nations, nil
}

const GetNationByIDQuery = `SELECT * FROM nation WHERE id = $1;`

func (repo *NationRepository) GetByID(id int) (domain.Nation, error) {
	row := repo.DB.QueryRow(GetNationByIDQuery, id)
	var nation domain.Nation
	if err := row.Scan(&nation.ID, &nation.Name, &nation.Acronym); err != nil {
		if err == sql.ErrNoRows {
			return domain.Nation{}, nil
		}
		return domain.Nation{}, err
	}
	return nation, nil
}

const CreateNationQuery = `INSERT INTO nation (name, acronym) VALUES ($1, $2) RETURNING id;`

func (repo *NationRepository) Create(nation domain.CreateNationRequest) (int, error) {
	var id int
	err := repo.DB.QueryRow(CreateNationQuery, nation.Name, nation.Acronym).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

const UpdateNationQuery = `UPDATE nation SET name = $1, acronym = $2 WHERE id = $3;`

func (repo *NationRepository) Update(id int, nation domain.Nation) error {
	_, err := repo.DB.Exec(UpdateNationQuery, nation.Name, nation.Acronym, id)
	if err != nil {
		return err
	}
	return nil
}
