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

type ModalityUseCase struct {
	ModalityRepository repositories.ModalityRepositoryInterface
	db                 *sql.DB
	logger             *zap.Logger
}

func NewModalityUseCase(modalityRepository repositories.ModalityRepositoryInterface, db *sql.DB, logger *zap.Logger) *ModalityUseCase {
	return &ModalityUseCase{
		ModalityRepository: modalityRepository,
		db:                 db,
		logger:             logger,
	}
}

func (uc ModalityUseCase) GetAll(ctx context.Context, c *gin.Context) {
	modalities, err := uc.ModalityRepository.GetAll()
	if err != nil {
		uc.logger.Error("Error fetching modalities", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if modalities == nil {
		c.JSON(200, gin.H{"message": "No data found"})
		return
	}
	c.JSON(200, gin.H{"data": modalities})
}

func (uc ModalityUseCase) GetAllByInactive(ctx context.Context, c *gin.Context) {
	modalities, err := uc.ModalityRepository.GetAllByInactive()
	if err != nil {
		uc.logger.Error("Error fetching inactive modalities", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if modalities == nil {
		c.JSON(200, gin.H{"message": "No data found"})
		return
	}
	c.JSON(200, gin.H{"data": modalities})
}

func (uc ModalityUseCase) GetByID(ctx context.Context, c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	modality, err := uc.ModalityRepository.GetByID(id)
	if err != nil {
		uc.logger.Error("Error fetching modality by ID", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if modality == (domain.Modality{}) {
		c.JSON(404, gin.H{"message": "No data found"})
		return
	}
	c.JSON(200, gin.H{"data": modality})
}

func (uc ModalityUseCase) Create(ctx context.Context, c *gin.Context) {
	var modality domain.Modality
	if err := c.ShouldBindJSON(&modality); err != nil {
		uc.logger.Error("Invalid request body", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	createRequest := domain.CreateModalityRequest{
		Name:        modality.Name,
		Amount_play: modality.Amount_play,
		Active:      modality.Active,
	}

	idModality, err := uc.ModalityRepository.Create(createRequest)
	if err != nil {
		uc.logger.Error("Error creating modality", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(201, gin.H{"message": "Modality created successfully", "id": idModality})
}

func (uc ModalityUseCase) Update(ctx context.Context, c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var modality domain.Modality
	if err := c.ShouldBindJSON(&modality); err != nil {
		uc.logger.Error("Invalid request body", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	modality.ID = id

	err = uc.ModalityRepository.Update(modality)
	if err != nil {
		uc.logger.Error("Error updating modality", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "Modality updated successfully"})
}

func (uc ModalityUseCase) Inactive(ctx context.Context, c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	err = uc.ModalityRepository.Inactive(id)
	if err != nil {
		uc.logger.Error("Error inactivating modality", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "Modality inactivated successfully"})
}

func (uc ModalityUseCase) Active(ctx context.Context, c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	err = uc.ModalityRepository.Active(id)
	if err != nil {
		uc.logger.Error("Error activating modality", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "Modality activated successfully"})
}
