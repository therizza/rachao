package repositories

import (
	"database/sql"
	"rachao/internal/core/domain"

	"github.com/google/uuid"
)

type AttributesRepository struct {
	DB *sql.DB
}

const GetByIDPositionQuery = `SELECT * FROM position WHERE id_position = $1;`

func (repo *AttributesRepository) GetByIDPosition(idPosition uuid.UUID) (domain.Attributes, error) {
	row := repo.DB.QueryRow(GetByIDPositionQuery, idPosition)
	var attributes domain.Attributes
	if err := row.Scan(&attributes.ID, &attributes.IDPosition, &attributes.PAC, &attributes.SHO, &attributes.PAS, &attributes.DRI, &attributes.DEF, &attributes.PHY); err != nil {
		if err == sql.ErrNoRows {
			return attributes, nil
		}
		return attributes, err
	}
	return attributes, nil
}

const CreateAttributesQuery = `INSERT INTO attributes (id_position, pac, sho, pas, dri, def, phy) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`

func (repo *AttributesRepository) Create(attributes domain.AttributesRequest) (int, error) {
	var id int
	err := repo.DB.QueryRow(CreateAttributesQuery, attributes.IDPosition, attributes.PAC, attributes.SHO, attributes.PAS, attributes.DRI, attributes.DEF, attributes.PHY).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

const GetAllAttributesQuery = `SELECT * FROM attributes ORDER BY id DESC;`

func (repo *AttributesRepository) GetAll() ([]domain.Attributes, error) {
	rows, err := repo.DB.Query(GetAllAttributesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attributes []domain.Attributes
	for rows.Next() {
		var attribute domain.Attributes
		if err := rows.Scan(&attribute.ID, &attribute.IDPosition, &attribute.PAC, &attribute.SHO, &attribute.PAS, &attribute.DRI, &attribute.DEF, &attribute.PHY); err != nil {
			return nil, err
		}
		attributes = append(attributes, attribute)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return attributes, nil
}

const UpdateAttributesQuery = `UPDATE attributes SET id_position = $1, pac = $2, sho = $3, pas = $4, dri = $5, def = $6, phy = $7 WHERE id = $8;`

func (repo *AttributesRepository) Update(attributes domain.AttributesRequest, id uuid.UUID) error {
	_, err := repo.DB.Exec(UpdateAttributesQuery, attributes.IDPosition, attributes.PAC, attributes.SHO, attributes.PAS, attributes.DRI, attributes.DEF, attributes.PHY, id)
	if err != nil {
		return err
	}
	return nil
}

const DeleteAttributesQuery = `DELETE FROM attributes WHERE id = $1;`

func (repo *AttributesRepository) Delete(id uuid.UUID) error {
	_, err := repo.DB.Exec(DeleteAttributesQuery, id)
	if err != nil {
		return err
	}
	return nil
}
