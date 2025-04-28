package repositories

import (
	"database/sql"
	"rachao/internal/core/domain"

	"github.com/google/uuid"
)

type CardPlayRepository struct {
	DB *sql.DB
}

const GetCardPlayAllQuery = `SELECT * FROM play p INNER JOIN card c ON p.id = c.id_play WHERE p.active = true;`

func (repo *CardPlayRepository) GetAll() ([]domain.CardPlay, error) {
	rows, err := repo.DB.Query(GetCardPlayAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cardPlays []domain.CardPlay
	for rows.Next() {
		var cardPlay domain.CardPlay
		if err := rows.Scan(
			&cardPlay.Play.ID,
			&cardPlay.Play.Name,
			&cardPlay.Play.IDPosition,
			&cardPlay.Play.IDNation,
			&cardPlay.Play.Field,
			&cardPlay.Play.Active,
			&cardPlay.Card.ID,
			&cardPlay.Card.IDPlay,
			&cardPlay.Card.PAC,
			&cardPlay.Card.SHO,
			&cardPlay.Card.PAS,
			&cardPlay.Card.DRI,
			&cardPlay.Card.DEF,
			&cardPlay.Card.PHY,
		); err != nil {
			return nil, err
		}
		cardPlays = append(cardPlays, cardPlay)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cardPlays, nil
}

const GetCardPlayAllByInactiveQuery = `SELECT * FROM play p INNER JOIN card c ON p.id = c.id_play WHERE p.active = false;`

func (repo *CardPlayRepository) GetAllByInactive() ([]domain.CardPlay, error) {
	rows, err := repo.DB.Query(GetCardPlayAllByInactiveQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cardPlays []domain.CardPlay
	for rows.Next() {
		var cardPlay domain.CardPlay
		if err := rows.Scan(
			&cardPlay.Play.ID,
			&cardPlay.Play.Name,
			&cardPlay.Play.IDPosition,
			&cardPlay.Play.IDNation,
			&cardPlay.Play.Field,
			&cardPlay.Play.Active,
			&cardPlay.Card.ID,
			&cardPlay.Card.IDPlay,
			&cardPlay.Card.PAC,
			&cardPlay.Card.SHO,
			&cardPlay.Card.PAS,
			&cardPlay.Card.DRI,
			&cardPlay.Card.DEF,
			&cardPlay.Card.PHY,
		); err != nil {
			return nil, err
		}
		cardPlays = append(cardPlays, cardPlay)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cardPlays, nil
}

const GetCardPlayByIDQuery = `SELECT * FROM play p INNER JOIN card c ON p.id = c.id_play WHERE p.id = $1;`

func (repo *CardPlayRepository) GetByID(id uuid.UUID) (domain.CardPlay, error) {
	row := repo.DB.QueryRow(GetCardPlayByIDQuery, id)
	var cardPlay domain.CardPlay
	if err := row.Scan(
		&cardPlay.Play.ID,
		&cardPlay.Play.Name,
		&cardPlay.Play.IDPosition,
		&cardPlay.Play.IDNation,
		&cardPlay.Play.Field,
		&cardPlay.Play.Active,
		&cardPlay.Card.ID,
		&cardPlay.Card.IDPlay,
		&cardPlay.Card.PAC,
		&cardPlay.Card.SHO,
		&cardPlay.Card.PAS,
		&cardPlay.Card.DRI,
		&cardPlay.Card.DEF,
		&cardPlay.Card.PHY,
	); err != nil {
		if err == sql.ErrNoRows {
			return cardPlay, nil
		}
		return cardPlay, err
	}
	return cardPlay, nil
}
