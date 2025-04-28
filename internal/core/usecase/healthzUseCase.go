package usecase

import (
	"net/http"
	"rachao/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type HealthzUseCase struct{}

func (h *HealthzUseCase) Healthz() gin.HandlerFunc {
	return func(c *gin.Context) {
		healthzEntity := &domain.Healthz{
			Message: "system is up",
			Version: 1.0,
			Status:  "ok",
		}

		c.JSON(http.StatusOK, healthzEntity)

	}
}
