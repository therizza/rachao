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

type CardUseCase struct {
	CardRepository repositories.CardRepositoryInterface
	db             *sql.DB
	logger         *zap.Logger
}

func NewCardUseCase(cardRepository repositories.CardRepositoryInterface, db *sql.DB, logger *zap.Logger) *CardUseCase {
	return &CardUseCase{
		CardRepository: cardRepository,
		db:             db,
		logger:         logger,
	}
}

func (uc CardUseCase) GetByID(ctx context.Context, c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	card, err := uc.CardRepository.GetByID(uuid)
	if err != nil {
		uc.logger.Error("Error fetching card", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"data": card})
}

func (uc CardUseCase) Create(ctx context.Context, c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	var card domain.CardRequest
	if err := c.ShouldBindJSON(&card); err != nil {
		uc.logger.Error("Invalid request body", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	cardID, err := uc.CardRepository.Create(uuid, card)
	if err != nil {
		uc.logger.Error("Error creating card", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(201, gin.H{"id": cardID})
}
func (uc CardUseCase) Update(ctx context.Context, c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var card domain.CardRequest
	if err := c.ShouldBindJSON(&card); err != nil {
		uc.logger.Error("Invalid request body", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	if card == (domain.CardRequest{}) {
		uc.logger.Error("Request body is empty")
		c.JSON(400, gin.H{"error": "Request body cannot be null"})
		return
	}
	err = uc.CardRepository.Update(uuid, card)
	if err != nil {
		uc.logger.Error("Error updating card", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "Card updated successfully"})
}
