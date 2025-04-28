package usecase

import (
	"context"
	"database/sql"
	"rachao/infra/repositories"
	"rachao/internal/core/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type CardPlayUseCase struct {
	CardPlayRepository repositories.CardPlayRepositoryInterface
	db                 *sql.DB
	logger             *zap.Logger
}

func NewCardPlayUseCase(cardPlayRepository repositories.CardPlayRepositoryInterface, db *sql.DB, logger *zap.Logger) *CardPlayUseCase {
	return &CardPlayUseCase{
		CardPlayRepository: cardPlayRepository,
		db:                 db,
		logger:             logger,
	}
}

func (uc CardPlayUseCase) GetAll(ctx context.Context, c *gin.Context) {
	cardPlays, err := uc.CardPlayRepository.GetAll()
	if err != nil {
		uc.logger.Error("Error fetching card plays", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if cardPlays == nil {
		c.JSON(200, gin.H{"message": "No data found"})
		return
	}
	c.JSON(200, gin.H{"data": cardPlays})
}

func (uc CardPlayUseCase) GetAllByInactive(ctx context.Context, c *gin.Context) {
	cardPlays, err := uc.CardPlayRepository.GetAllByInactive()
	if err != nil {
		uc.logger.Error("Error fetching inactive card plays", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if cardPlays == nil {
		c.JSON(200, gin.H{"message": "No data found"})
		return
	}
	c.JSON(200, gin.H{"data": cardPlays})
}

func (uc CardPlayUseCase) GetByID(ctx context.Context, c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	cardPlay, err := uc.CardPlayRepository.GetByID(uuid)
	if err != nil {
		uc.logger.Error("Error fetching card play by ID", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if cardPlay == (domain.CardPlay{}) {
		c.JSON(404, gin.H{"message": "Card play not found"})
		return
	}
	c.JSON(200, gin.H{"data": cardPlay})
}
