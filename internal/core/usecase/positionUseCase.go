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

type PositionUseCase struct {
	PositionRepository repositories.PositionRepositoryInterface
	db                 *sql.DB
	logger             *zap.Logger
}

func NewPositionUseCase(positionRepository repositories.PositionRepositoryInterface, db *sql.DB, logger *zap.Logger) *PositionUseCase {
	return &PositionUseCase{
		PositionRepository: positionRepository,
		db:                 db,
		logger:             logger,
	}
}

func (uc PositionUseCase) GetAll(ctx context.Context, c *gin.Context) {
	positions, err := uc.PositionRepository.GetAll()
	if err != nil {
		uc.logger.Error("Error fetching positions", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if positions == nil {
		c.JSON(200, gin.H{"message": "No data found"})
		return
	}
	c.JSON(200, gin.H{"data": positions})
}

func (uc PositionUseCase) GetByID(ctx context.Context, c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	position, err := uc.PositionRepository.GetByID(idInt)
	if err != nil {
		uc.logger.Error("Error fetching position by ID", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if position == (domain.Position{}) {
		c.JSON(404, gin.H{"message": "Position not found"})
		return
	}
	c.JSON(200, gin.H{"data": position})
}

func (uc PositionUseCase) Create(ctx context.Context, c *gin.Context) {
	positionRequest := domain.CreatePositionRequest{}
	if err := c.ShouldBindJSON(&positionRequest); err != nil {
		uc.logger.Error("Invalid request body", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	id, err := uc.PositionRepository.Create(positionRequest)
	if err != nil {
		uc.logger.Error("Error creating position", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(201, gin.H{"message": "Position created successfully", "id": id})
}

func (uc PositionUseCase) Update(ctx context.Context, c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	err = uc.validatePositionExists(ctx, idInt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"message": "Position not found"})
			return
		}
	}
	positionRequest := domain.Position{}
	if err := c.ShouldBindJSON(&positionRequest); err != nil {
		uc.logger.Error("Invalid request body", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	positionRequest.ID = idInt
	err = uc.PositionRepository.Update(idInt, positionRequest)
	if err != nil {
		uc.logger.Error("Error updating position", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "Position updated successfully"})
}

func (uc PositionUseCase) Delete(ctx context.Context, c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	err = uc.validatePositionExists(ctx, idInt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(404, gin.H{"message": "Position not found"})
			return
		}
	}
	err = uc.PositionRepository.Delete(idInt)
	if err != nil {
		uc.logger.Error("Error deleting position", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "Position deleted successfully"})
}

func (uc PositionUseCase) validatePositionExists(_ context.Context, id int) error {
	position, err := uc.PositionRepository.GetByID(id)
	if err != nil {
		return err
	}
	if position == (domain.Position{}) {
		return sql.ErrNoRows
	}
	return nil
}
