package repositories

import (
	"database/sql"
	"rachao/internal/core/domain"

	"github.com/google/uuid"
)

type CardRepository struct {
	DB *sql.DB
}

const GetByIDQuery = `SELECT * FROM card WHERE id_play = $1;`

func (repo *CardRepository) GetByIDPlay(id uuid.UUID) (domain.Card, error) {
	row := repo.DB.QueryRow(GetByIDQuery, id)
	var card domain.Card
	if err := row.Scan(&card.ID, &card.IDPlay, &card.PAC, &card.SHO, &card.PAS, &card.DRI, &card.DEF, &card.PHY); err != nil {
		return domain.Card{}, err
	}
	return card, nil
}

const GetByID = `SELECT * FROM card WHERE id = $1;`

func (repo *CardRepository) GetByID(id uuid.UUID) (domain.Card, error) {
	row := repo.DB.QueryRow(GetByID, id)
	var card domain.Card
	if err := row.Scan(&card.ID, &card.IDPlay, &card.PAC, &card.SHO, &card.PAS, &card.DRI, &card.DEF, &card.PHY); err != nil {
		return domain.Card{}, err
	}
	return card, nil
}

const CreateQuery = `INSERT INTO card (id_play, pac, sho, pas, dri, def, phy) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`

func (repo *CardRepository) Create(ID uuid.UUID, card domain.CardRequest) (uuid.UUID, error) {
	var id uuid.UUID

	var idPlayExists bool
	err := repo.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM card WHERE id_play = $1)", ID).Scan(&idPlayExists)
	if err != nil {
		return uuid.Nil, err
	}
	if idPlayExists {
		return uuid.Nil, sql.ErrNoRows
	}

	err = repo.DB.QueryRow(CreateQuery, ID, card.PAC, card.SHO, card.PAS, card.DRI, card.DEF, card.PHY).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

const UpdateQuery = `UPDATE card SET pac = $1, sho = $2, pas = $3, dri = $4, def = $5, phy = $6 WHERE id_play = $7;`

func (repo *CardRepository) Update(id uuid.UUID, card domain.CardRequest) (idCard uuid.UUID, erro error) {
	result, err := repo.DB.Exec(UpdateQuery, card.PAC, card.SHO, card.PAS, card.DRI, card.DEF, card.PHY, id)
	if err != nil {
		return uuid.Nil, err
	}
	cardResult, err := repo.GetByIDPlay(id)
	if err != nil {
		return uuid.Nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return uuid.Nil, err
	}
	if rowsAffected == 0 {
		return uuid.Nil, sql.ErrNoRows
	}
	return cardResult.ID, nil
}
