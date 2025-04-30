package usecase

import (
	"context"
	"database/sql"
	"rachao/infra/repositories"
	"rachao/internal/core/domain"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AttributesUseCase struct {
	AttributesRepository repositories.AttributeRepositoryInterface
	PositionRepository   repositories.PositionRepositoryInterface
	db                   *sql.DB
	logger               *zap.Logger
}

func NewAttributesUseCase(attributesRepository repositories.AttributeRepositoryInterface, positionRepository repositories.PositionRepositoryInterface, db *sql.DB, logger *zap.Logger) *AttributesUseCase {
	return &AttributesUseCase{
		AttributesRepository: attributesRepository,
		PositionRepository:   positionRepository,
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

	err := uc.validateExists(ctx, attributes.IDPosition)
	if err != nil {
		uc.logger.Error("Error fetching attributes by ID", zap.Error(err))
		c.JSON(400, gin.H{"error": "Bad Request"})
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

func (uc AttributesUseCase) GetAll(ctx context.Context, c *gin.Context) {
	attributes, err := uc.AttributesRepository.GetAll()
	if err != nil {
		uc.logger.Error("Error fetching attributes", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if attributes == nil {
		c.JSON(200, gin.H{"message": "No data found"})
		return
	}
	c.JSON(200, gin.H{"data": attributes})
}

func (uc AttributesUseCase) GetByIDPosition(ctx context.Context, c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	attributes, err := uc.AttributesRepository.GetByIDPosition(idInt)
	if err != nil {
		uc.logger.Error("Error fetching attributes by ID", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if attributes == (domain.Attributes{}) {
		c.JSON(404, gin.H{"message": "Attributes not found"})
		return
	}
	c.JSON(200, gin.H{"data": attributes})
}

func (uc AttributesUseCase) GetByIDAttributes(ctx context.Context, c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	attributes, err := uc.AttributesRepository.GetByIDAttributes(idInt)
	if err != nil {
		uc.logger.Error("Error fetching attributes by ID", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if attributes == (domain.Attributes{}) {
		c.JSON(404, gin.H{"message": "Attributes not found"})
		return
	}
	c.JSON(200, gin.H{"data": attributes})
}

func (uc AttributesUseCase) Update(ctx context.Context, c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	attributes := domain.AttributesRequest{}
	if err := c.ShouldBindJSON(&attributes); err != nil {
		uc.logger.Error("Error binding JSON", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	err = uc.validateAttributesExists(ctx, idInt)
	if err != nil {
		uc.logger.Error("Error fetching attributes by ID", zap.Error(err))
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}

	err = uc.AttributesRepository.Update(attributes, idInt)
	if err != nil {
		uc.logger.Error("Error updating attributes", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "Attributes updated successfully"})
}

func (uc AttributesUseCase) Delete(ctx context.Context, c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	err = uc.validateAttributesExists(ctx, idInt)
	if err != nil {
		uc.logger.Error("Error fetching attributes by ID", zap.Error(err))
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}
	err = uc.AttributesRepository.Delete(idInt)
	if err != nil {
		uc.logger.Error("Error deleting attributes", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "Attributes deleted successfully"})
}

func (uc AttributesUseCase) validateAttributesExists(_ context.Context, id int) error {
	var attributes domain.Attributes

	attributes, err := uc.AttributesRepository.GetByIDAttributes(id)
	if err != nil {
		uc.logger.Error("Error fetching attributes by ID", zap.Error(err))
		return err
	}
	if attributes == (domain.Attributes{}) {
		return sql.ErrNoRows
	}

	return nil
}

func (uc AttributesUseCase) validateExists(_ context.Context, id int) error {
	var attributes domain.Attributes

	possition, err := uc.PositionRepository.GetByID(id)
	if err != nil {
		uc.logger.Error("Error fetching position by ID", zap.Error(err))
		return err
	}
	if possition == (domain.Position{}) {
		return sql.ErrNoRows
	}

	attributes, err = uc.AttributesRepository.GetByIDPosition(id)
	if err != nil {
		uc.logger.Error("Error fetching attributes by ID", zap.Error(err))
		return err
	}
	if attributes != (domain.Attributes{}) {
		return sql.ErrNoRows
	}

	return nil
}
