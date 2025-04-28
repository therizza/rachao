package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"rachao/infra/repositories"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type PhotoUseCase struct {
	PhotoRepository repositories.PhotoRepositoryInterface
	db              *sql.DB
	logger          *zap.Logger
}

func NewPhotoUseCase(photoRepository repositories.PhotoRepositoryInterface, db *sql.DB, logger *zap.Logger) *PhotoUseCase {
	return &PhotoUseCase{
		PhotoRepository: photoRepository,
		db:              db,
		logger:          logger,
	}
}

func (uc PhotoUseCase) GetByIDPlay(ctx context.Context, c *gin.Context) {
	idPlay := c.Param("id")
	uuid, err := uuid.Parse(idPlay)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	photos, err := uc.PhotoRepository.GetByIDPlay(uuid)
	if err != nil {
		uc.logger.Error("Error fetching photos", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	if len(photos) == 0 {
		c.JSON(404, gin.H{"error": "No photo found"})
		return
	}

	c.Header("Content-Type", "image/jpeg")
	c.Writer.WriteHeader(200)
	_, err = c.Writer.Write(photos[0].Photo)
	if err != nil {
		uc.logger.Error("Error writing photo to response", zap.Error(err))
	}
}

func (uc PhotoUseCase) Create(ctx context.Context, c *gin.Context) {
	idPlay := c.Param("id")
	uuid, err := uuid.Parse(idPlay)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	err = uc.validatePhotoExists(ctx, uuid)
	if err != nil {
		uc.logger.Error("Photo already exists", zap.Error(err))
		c.JSON(400, gin.H{"error": "Photo already exists"})
		return
	}
	photo, err := c.FormFile("photo")
	if err != nil {
		uc.logger.Error("Error getting photo from form", zap.Error(err))
		c.JSON(400, gin.H{"error": "Error getting photo from form"})
		return
	}
	file, err := photo.Open()
	if err != nil {
		uc.logger.Error("Error opening photo file", zap.Error(err))
		c.JSON(400, gin.H{"error": "Error opening photo file"})
		return
	}
	defer file.Close()
	photoBytes := make([]byte, photo.Size)
	_, err = file.Read(photoBytes)
	if err != nil {
		uc.logger.Error("Error reading photo file", zap.Error(err))
		c.JSON(400, gin.H{"error": "Error reading photo file"})
		return
	}
	idPhoto, err := uc.PhotoRepository.Create(uuid, photoBytes)
	if err != nil {
		uc.logger.Error("Error creating photo", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"data": idPhoto})
}

func (uc PhotoUseCase) Update(ctx context.Context, c *gin.Context) {
	idPlay := c.Param("id")
	uuid, err := uuid.Parse(idPlay)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	err = uc.validatePhotoExists(ctx, uuid)
	if err != nil {
		uc.logger.Error("Photo already exists", zap.Error(err))
		c.JSON(400, gin.H{"error": "Photo already exists"})
		return
	}
	photo, err := c.FormFile("photo")
	if err != nil {
		uc.logger.Error("Error getting photo from form", zap.Error(err))
		c.JSON(400, gin.H{"error": "Error getting photo from form"})
		return
	}
	file, err := photo.Open()
	if err != nil {
		uc.logger.Error("Error opening photo file", zap.Error(err))
		c.JSON(400, gin.H{"error": "Error opening photo file"})
		return
	}
	defer file.Close()
	photoBytes := make([]byte, photo.Size)
	_, err = file.Read(photoBytes)
	if err != nil {
		uc.logger.Error("Error reading photo file", zap.Error(err))
		c.JSON(400, gin.H{"error": "Error reading photo file"})
		return
	}
	err = uc.PhotoRepository.Update(uuid, photoBytes)
	if err != nil {
		uc.logger.Error("Error updating photo", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "Photo updated successfully"})
}

func (uc PhotoUseCase) Delete(ctx context.Context, c *gin.Context) {
	idPlay := c.Param("id")
	uuid, err := uuid.Parse(idPlay)
	if err != nil {
		uc.logger.Error("Invalid ID format", zap.Error(err))
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	err = uc.validatePhotoExists(ctx, uuid)
	if err != nil {
		uc.logger.Error("Photo already exists", zap.Error(err))
		c.JSON(400, gin.H{"error": "Photo already exists"})
		return
	}
	err = uc.PhotoRepository.Delete(uuid)
	if err != nil {
		uc.logger.Error("Error deleting photo", zap.Error(err))
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "Photo deleted successfully"})
}

func (uc PhotoUseCase) validatePhotoExists(_ context.Context, uuid uuid.UUID) error {
	play, err := uc.PhotoRepository.GetByIDPlay(uuid)
	if err != nil {
		return err
	}
	if len(play) != 0 {
		return fmt.Errorf("photo does not exist")
	}
	return nil
}
