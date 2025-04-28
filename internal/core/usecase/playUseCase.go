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

type PlayUseCase struct {
	PlayRepository repositories.PlayRepositoryInterface
	db             *sql.DB
	logger         *zap.Logger
}

func NewPlayUseCase(playRepository repositories.PlayRepositoryInterface, db *sql.DB, logger *zap.Logger) *PlayUseCase {
	return &PlayUseCase{
		PlayRepository: playRepository,
		db:             db,
		logger:         logger,
	}
}

func (uc PlayUseCase) GetAll(ctx context.Context, c *gin.Context) {
	plays, err := uc.PlayRepository.GetAll()
	if err != nil {
		uc.logger.Error("Error fetching plays", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if plays == nil {
		c.JSON(200, gin.H{"message": "No data found"})
		return
	}
	c.JSON(200, gin.H{"data": plays})
}

func (uc PlayUseCase) GetAllByInactive(ctx context.Context, c *gin.Context) {
	plays, err := uc.PlayRepository.GetAllByInactive()
	if err != nil {
		uc.logger.Error("Error fetching inactive plays", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if plays == nil {
		c.JSON(200, gin.H{"message": "No data found"})
		return
	}
	c.JSON(200, gin.H{"data": plays})
}

func (uc PlayUseCase) GetByID(ctx context.Context, c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	play, err := uc.PlayRepository.GetByID(uuid)
	if err != nil {
		uc.logger.Error("Error fetching play by ID", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if play == (domain.Play{}) {
		c.JSON(404, gin.H{"message": "Play not found"})
		return
	}
	c.JSON(200, gin.H{"data": play})
}

func (uc PlayUseCase) Create(ctx context.Context, c *gin.Context) {
	play := domain.CreatePlayRequest{}
	if err := c.ShouldBindJSON(&play); err != nil {
		uc.logger.Error("Invalid request body", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	playName, err := uc.PlayRepository.GetByName(play.Name)
	if err != nil {
		uc.logger.Error("Error fetching play by name", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if playName != (domain.Play{}) {
		c.JSON(400, gin.H{"message": "Play already exists"})
		return
	}
	id, err := uc.PlayRepository.Create(play)
	if err != nil {
		uc.logger.Error("Error creating play", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(201, gin.H{"message": "Play created successfully", "id": id})
}

func (uc PlayUseCase) Update(ctx context.Context, c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		uc.logger.Error("Invalid UUID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid UUID format"})
		return
	}
	play := domain.Play{}
	if err := c.ShouldBindJSON(&play); err != nil {
		uc.logger.Error("Invalid request body", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	playID, err := uc.fetchPlayByID(ctx, uuid)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if playID == (domain.Play{}) {
		c.JSON(404, gin.H{"message": "Play not found"})
		return
	}
	err = uc.PlayRepository.Update(uuid, play)
	if err != nil {
		uc.logger.Error("Error updating play", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "Play updated successfully"})
}

func (uc PlayUseCase) Delete(ctx context.Context, c *gin.Context) {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		uc.logger.Error("Invalid UUID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid UUID format"})
		return
	}
	playID, err := uc.fetchPlayByID(ctx, uuid)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	err = uc.PlayRepository.Delete(playID.ID)
	if err != nil {
		uc.logger.Error("Error deleting play", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "Play deleted successfully"})
}

func (uc PlayUseCase) GetByName(ctx context.Context, c *gin.Context) {
	name := c.Param("name")
	play, err := uc.PlayRepository.GetByName(name)
	if err != nil {
		uc.logger.Error("Error fetching play by name", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if play == (domain.Play{}) {
		c.JSON(404, gin.H{"message": "Play not found"})
		return
	}
	c.JSON(200, gin.H{"data": play})
}

func (uc PlayUseCase) fetchPlayByID(_ context.Context, id uuid.UUID) (domain.Play, error) {
	play, err := uc.PlayRepository.GetByID(id)
	if err != nil {
		uc.logger.Error("Error fetching play by ID", zap.Error(err))
		return domain.Play{}, err
	}
	if play == (domain.Play{}) {
		uc.logger.Error("Play not found")
		return play, err
	}
	return play, nil
}
