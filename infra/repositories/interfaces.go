package repositories

import (
	"rachao/internal/core/domain"

	"github.com/google/uuid"
)

type PlayRepositoryInterface interface {
	GetAll() ([]domain.Play, error)
	GetAllByInactive() ([]domain.Play, error)
	GetByID(id uuid.UUID) (domain.Play, error)
	Create(play domain.CreatePlayRequest) (uuid.UUID, error)
	Update(id uuid.UUID, play domain.Play) error
	Delete(id uuid.UUID) error
	GetByName(name string) (domain.Play, error)
}

type CardRepositoryInterface interface {
	GetByIDPlay(id uuid.UUID) (domain.Card, error)
	GetByID(id uuid.UUID) (domain.Card, error)
	Create(id uuid.UUID, card domain.CardRequest) (uuid.UUID, error)
	Update(id uuid.UUID, card domain.CardRequest) (idCard uuid.UUID, erro error)
}

type CardPlayRepositoryInterface interface {
	GetAll() ([]domain.CardPlay, error)
	GetAllByInactive() ([]domain.CardPlay, error)
	GetByID(id uuid.UUID) (domain.CardPlay, error)
}

type NationRepositoryInterface interface {
	GetAll() ([]domain.Nation, error)
	GetByID(id int) (domain.Nation, error)
	Create(nation domain.CreateNationRequest) (int, error)
	Update(id int, nation domain.Nation) error
}

type PhotoRepositoryInterface interface {
	GetByIDPlay(idPlay uuid.UUID) ([]domain.Photo, error)
	Create(idPlay uuid.UUID, photo []byte) (uuid.UUID, error)
	Update(idPlay uuid.UUID, photo []byte) error
	Delete(idPlay uuid.UUID) error
}

type PositionRepositoryInterface interface {
	GetAll() ([]domain.Position, error)
	GetByID(id int) (domain.Position, error)
	Create(position domain.CreatePositionRequest) (int, error)
	Update(id int, position domain.Position) error
	Delete(id int) error
}

type AttributeRepositoryInterface interface {
	GetByIDPosition(idPosition int) (domain.Attributes, error)
	GetByIDAttributes(id int) (domain.Attributes, error)
	GetAll() ([]domain.Attributes, error)
	Create(attributes domain.AttributesRequest) (int, error)
	Update(attributes domain.AttributesRequest, id int) error
	Delete(id int) error
}

type OverallRepositoryInterface interface {
	Exists(idPlay uuid.UUID) (bool, error)
	GetByIDPlay(idUser uuid.UUID) (domain.Overall, error)
	Create(overall domain.OverallRequest) (uuid.UUID, error)
	Update(overall domain.OverallRequest, idUser uuid.UUID) error
	Delete(idUser uuid.UUID) error
}

type ModalityRepositoryInterface interface {
	GetAll() ([]domain.Modality, error)
	GetAllByInactive() ([]domain.Modality, error)
	GetByID(id int) (domain.Modality, error)
	Create(modality domain.CreateModalityRequest) (int, error)
	Update(modality domain.Modality) error
	Inactive(id int) error
	Active(id int) error
	GetByName(name string) (domain.Modality, error)
}
