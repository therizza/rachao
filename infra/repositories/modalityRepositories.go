package repositories

import (
	"database/sql"
	"rachao/internal/core/domain"
)

type ModalityRepository struct {
	DB *sql.DB
}

const GetModalityAllQuery = `SELECT * FROM modality WHERE active = true;`

func (repo *ModalityRepository) GetAll() ([]domain.Modality, error) {
	rows, err := repo.DB.Query(GetModalityAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modalities []domain.Modality
	for rows.Next() {
		var modality domain.Modality
		if err := rows.Scan(&modality.ID, &modality.Name, &modality.Amount_play, &modality.Active); err != nil {
			return nil, err
		}
		modalities = append(modalities, modality)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return modalities, nil
}

const GetModalityAllByInactiveQuery = `SELECT * FROM modality WHERE active = false;`

func (repo *ModalityRepository) GetAllByInactive() ([]domain.Modality, error) {
	rows, err := repo.DB.Query(GetModalityAllByInactiveQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var modalities []domain.Modality
	for rows.Next() {
		var modality domain.Modality
		if err := rows.Scan(&modality.ID, &modality.Name, &modality.Amount_play, &modality.Active); err != nil {
			return nil, err
		}
		modalities = append(modalities, modality)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return modalities, nil
}

const GetModalityByIDQuery = `SELECT * FROM modality WHERE id = $1;`

func (repo *ModalityRepository) GetByID(id int) (domain.Modality, error) {
	row := repo.DB.QueryRow(GetModalityByIDQuery, id)
	var modality domain.Modality
	if err := row.Scan(&modality.ID, &modality.Name, &modality.Amount_play, &modality.Active); err != nil {
		return modality, err
	}
	return modality, nil
}

const CreateModalityQuery = `INSERT INTO modality (name, amount_play, active) VALUES ($1, $2, $3) RETURNING id;`

func (repo *ModalityRepository) Create(modality domain.CreateModalityRequest) (int, error) {
	var id int
	err := repo.DB.QueryRow(CreateModalityQuery, modality.Name, modality.Amount_play, modality.Active).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

const UpdateModalityQuery = `UPDATE modality SET name = $1, amount_play = $2, active = $3 WHERE id = $4;`

func (repo *ModalityRepository) Update(modality domain.Modality) error {
	_, err := repo.DB.Exec(UpdateModalityQuery, modality.Name, modality.Amount_play, modality.Active, modality.ID)
	if err != nil {
		return err
	}
	return nil
}

const InactiveModalityQuery = `UPDATE modality SET active = false WHERE id = $1;`

func (repo *ModalityRepository) Inactive(id int) error {
	_, err := repo.DB.Exec(InactiveModalityQuery, id)
	if err != nil {
		return err
	}
	return nil
}

const ActiveModalityQuery = `UPDATE modality SET active = true WHERE id = $1;`

func (repo *ModalityRepository) Active(id int) error {
	_, err := repo.DB.Exec(ActiveModalityQuery, id)
	if err != nil {
		return err
	}
	return nil
}

const GetModalityByNameQuery = `SELECT * FROM modality WHERE name = $1;`

func (repo *ModalityRepository) GetByName(name string) (domain.Modality, error) {
	row := repo.DB.QueryRow(GetModalityByNameQuery, name)
	var modality domain.Modality
	if err := row.Scan(&modality.ID, &modality.Name, &modality.Active); err != nil {
		return modality, err
	}
	return modality, nil
}
