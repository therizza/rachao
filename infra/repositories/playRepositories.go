package repositories

import (
	"database/sql"
	"rachao/internal/core/domain"

	"github.com/google/uuid"
)

type PlayRepository struct {
	DB *sql.DB
}

const GetPlayAllQuery = `SELECT * FROM play WHERE active = true;`

func (repo *PlayRepository) GetAll() ([]domain.Play, error) {
	rows, err := repo.DB.Query(GetPlayAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plays []domain.Play
	for rows.Next() {
		var play domain.Play
		if err := rows.Scan(&play.ID, &play.Name, &play.IDPosition, &play.IDNation, &play.Field, &play.Active); err != nil {
			return nil, err
		}
		plays = append(plays, play)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return plays, nil
}

const GetPlayAllByInactiveQuery = `SELECT * FROM play WHERE active = false;`

func (repo *PlayRepository) GetAllByInactive() ([]domain.Play, error) {
	rows, err := repo.DB.Query(GetPlayAllByInactiveQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plays []domain.Play
	for rows.Next() {
		var play domain.Play
		if err := rows.Scan(&play.ID, &play.Name, &play.IDPosition, &play.IDNation, &play.Field, &play.Active); err != nil {
			return nil, err
		}
		plays = append(plays, play)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return plays, nil
}

const GetPlayByIDQuery = `SELECT * FROM play WHERE id = $1;`

func (repo *PlayRepository) GetByID(id uuid.UUID) (domain.Play, error) {
	row := repo.DB.QueryRow(GetPlayByIDQuery, id)
	var play domain.Play
	if err := row.Scan(&play.ID, &play.Name, &play.IDPosition, &play.IDNation, &play.Field, &play.Active); err != nil {
		if err == sql.ErrNoRows {
			return play, nil
		}
		return play, err
	}
	return play, nil
}

const CreatePlayQuery = `INSERT INTO play (name, id_position, id_nation, field, active) VALUES ($1, $2, $3, $4, $5) RETURNING id;`

func (repo *PlayRepository) Create(play domain.CreatePlayRequest) (uuid.UUID, error) {
	var id uuid.UUID
	err := repo.DB.QueryRow(CreatePlayQuery, play.Name, play.IDPosition, play.IDNation, play.Field, play.Active).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

const UpdatePlayQuery = `UPDATE play SET name = $1, id_position = $2, id_nation = $3, field = $4, active = $5 WHERE id = $6;`

func (repo *PlayRepository) Update(id uuid.UUID, play domain.Play) error {
	_, err := repo.DB.Exec(UpdatePlayQuery, play.Name, play.IDPosition, play.IDNation, play.Field, play.Active, id)
	if err != nil {
		return err
	}
	return nil
}

const DeletePlayQuery = `UPDATE play SET active = false WHERE id = $1;`

func (repo *PlayRepository) Delete(id uuid.UUID) error {
	_, err := repo.DB.Exec(DeletePlayQuery, id)
	if err != nil {
		return err
	}
	return nil
}

const GetPlayByNameQuery = `SELECT * FROM play WHERE name = $1;`

func (repo *PlayRepository) GetByName(name string) (domain.Play, error) {
	row := repo.DB.QueryRow(GetPlayByNameQuery, name)
	var play domain.Play
	if err := row.Scan(&play.ID, &play.Name, &play.IDPosition, &play.IDNation, &play.Field, &play.Active); err != nil {
		if err == sql.ErrNoRows {
			return play, nil
		}
		return play, err
	}
	return play, nil
}
