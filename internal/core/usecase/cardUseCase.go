package usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"rachao/infra/messaging"
	"rachao/infra/repositories"
	"rachao/internal/core/domain"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type CardUseCase struct {
	CardRepository      repositories.CardRepositoryInterface
	CardPlayRepository  repositories.CardPlayRepositoryInterface
	AttributeRepository repositories.AttributeRepositoryInterface
	MessagingUseCase    messaging.MessagePublisherInterface
	db                  *sql.DB
	logger              *zap.Logger
}

func NewCardUseCase(
	cardRepository repositories.CardRepositoryInterface,
	cardPlayRepository repositories.CardPlayRepositoryInterface,
	attributeRepository repositories.AttributeRepositoryInterface,
	messagingUseCase messaging.MessagePublisherInterface,
	db *sql.DB,
	logger *zap.Logger,

) *CardUseCase {
	return &CardUseCase{
		CardRepository:      cardRepository,
		CardPlayRepository:  cardPlayRepository,
		AttributeRepository: attributeRepository,
		MessagingUseCase:    messagingUseCase,
		db:                  db,
		logger:              logger,
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
	card, err := uc.CardRepository.GetByIDPlay(uuid)
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

	err = uc.calculatorOverall(cardID)
	if err != nil {
		uc.logger.Error("Error calculating overall", zap.Error(err))
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
	IDCard, err := uc.CardRepository.Update(uuid, card)
	if err != nil {
		uc.logger.Error("Error updating card", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	err = uc.calculatorOverall(IDCard)
	if err != nil {
		uc.logger.Error("Error calculating overall", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{"message": "Card updated successfully"})
}

func (uc CardUseCase) calculatorOverall(id uuid.UUID) error {
	card, err := uc.CardRepository.GetByID(id)
	if err != nil {
		uc.logger.Error("Error fetching card", zap.Error(err))
		return err
	}

	cardPlay, err := uc.CardPlayRepository.GetByID(card.IDPlay)
	if err != nil {
		uc.logger.Error("Error fetching card play", zap.Error(err))
		return err
	}

	attributes, err := uc.AttributeRepository.GetByIDPosition(cardPlay.IDPosition)
	if err != nil {
		uc.logger.Error("Error fetching attributes", zap.Error(err))
		return err

	}

	overallRequest := domain.OverallBodyRequest{
		Card:       card,
		Attributes: attributes,
	}

	overallRequestBytes, err := json.Marshal(overallRequest)
	if err != nil {
		uc.logger.Error("Error serializing overall request", zap.Error(err))
		return err
	}

	cardMessaging := "card." + card.IDPlay.String()
	err = uc.MessagingUseCase.Publish("rachao", cardMessaging, overallRequestBytes)
	if err != nil {
		uc.logger.Error("Error publishing: ", zap.Error(err))
		return err

	}

	return nil
}
