package usecase

import (
	"context"
	"database/sql"
	"rachao/infra/repositories"
	"rachao/internal/core/domain"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AttributesUseCase struct {
	AttributesRepository repositories.AttributeRepositoryInterface
	db                   *sql.DB
	logger               *zap.Logger
}

func NewAttributesUseCase(attributesRepository repositories.AttributeRepositoryInterface, db *sql.DB, logger *zap.Logger) *AttributesUseCase {
	return &AttributesUseCase{
		AttributesRepository: attributesRepository,
		db:                   db,
		logger:               logger,
	}
}

func (uc AttributesUseCase) Create(ctx context.Context, c *gin.Context) {
	attributes := domain.AttributesRequest{}
	if err := c.ShouldBindJSON(&attributes); err != nil {
		uc.logger.Error("Error binding JSON", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}
	attributesID, err := uc.AttributesRepository.Create(attributes)
	if err != nil {
		uc.logger.Error("Error creating attributes", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(201, gin.H{"data": attributesID})

}

// func (uc AttributesUseCase) fetchAttributesByID(_ context.Context, id uuid.UUID) (domain.Attributes, error) {
// 	attributes, err := uc.AttributesRepository.GetByIDPosition(id)
// 	if err != nil {
// 		uc.logger.Error("Error fetching attributes by ID", zap.Error(err))
// 		return domain.Attributes{}, err
// 	}
// 	if attributes == (domain.Attributes{}) {
// 		return domain.Attributes{}, nil
// 	}
// 	return attributes, nil
// }
