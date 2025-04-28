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

type NationUseCase struct {
	NationRepository repositories.NationRepositoryInterface
	db               *sql.DB
	logger           *zap.Logger
}

func NewNationUseCase(nationRepository repositories.NationRepositoryInterface, db *sql.DB, logger *zap.Logger) *NationUseCase {
	return &NationUseCase{
		NationRepository: nationRepository,
		db:               db,
		logger:           logger,
	}
}

func (uc NationUseCase) GetAll(ctx context.Context, c *gin.Context) {
	nations, err := uc.NationRepository.GetAll()
	if err != nil {
		uc.logger.Error("Error fetching nations", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if nations == nil {
		c.JSON(200, gin.H{"message": "No data found"})
		return
	}
	c.JSON(200, gin.H{"data": nations})
}

func (uc NationUseCase) GetByID(ctx context.Context, c *gin.Context) {
	id := c.Param("id")
	nationID, err := strconv.Atoi(id)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	nation, err := uc.NationRepository.GetByID(nationID)
	if err != nil {
		uc.logger.Error("Error fetching nation by ID", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if nation == (domain.Nation{}) {
		c.JSON(404, gin.H{"message": "Nation not found"})
		return
	}
	c.JSON(200, gin.H{"data": nation})
}

func (uc NationUseCase) Create(ctx context.Context, c *gin.Context) {
	var nation domain.CreateNationRequest
	if err := c.ShouldBindJSON(&nation); err != nil {
		uc.logger.Error("Invalid request body", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := uc.NationRepository.Create(nation)
	if err != nil {
		uc.logger.Error("Error creating nation", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(201, gin.H{"message": "Nation created successfully", "id": id})
}

func (uc NationUseCase) Update(ctx context.Context, c *gin.Context) {
	id := c.Param("id")
	nationID, err := strconv.Atoi(id)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var nation domain.Nation
	if err := c.ShouldBindJSON(&nation); err != nil {
		uc.logger.Error("Invalid request body", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	err = uc.NationRepository.Update(nationID, nation)
	if err != nil {
		uc.logger.Error("Error updating nation", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "Nation updated successfully"})
}
