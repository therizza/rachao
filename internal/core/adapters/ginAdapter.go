package adapters

import (
	"rachao/internal/core/usecase"

	"github.com/gin-gonic/gin"
)

type GinAdapter struct {
	HealthzUseCase  *usecase.HealthzUseCase
	PlayUseCase     *usecase.PlayUseCase
	CardUseCase     *usecase.CardUseCase
	CardPlayUseCase *usecase.CardPlayUseCase
	NationUseCase   *usecase.NationUseCase
	PhotoUseCase    *usecase.PhotoUseCase
	PositionUseCase *usecase.PositionUseCase
	Attributes      *usecase.AttributesUseCase
}

func NewGinAdapter(
	healthzUseCase *usecase.HealthzUseCase,
	playUseCase *usecase.PlayUseCase,
	cardUseCase *usecase.CardUseCase,
	cardPlayUseCase *usecase.CardPlayUseCase,
	nationUseCase *usecase.NationUseCase,
	photoUseCase *usecase.PhotoUseCase,
	positionUseCase *usecase.PositionUseCase,
	attributes *usecase.AttributesUseCase,
) *GinAdapter {
	return &GinAdapter{
		HealthzUseCase:  healthzUseCase,
		PlayUseCase:     playUseCase,
		CardUseCase:     cardUseCase,
		CardPlayUseCase: cardPlayUseCase,
		NationUseCase:   nationUseCase,
		PhotoUseCase:    photoUseCase,
		PositionUseCase: positionUseCase,
		Attributes:      attributes,
	}
}

func (ga *GinAdapter) SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/healthz", ga.HealthzUseCase.Healthz())

	r.GET("/play", func(c *gin.Context) {
		ga.PlayUseCase.GetAll(c.Request.Context(), c)
	})
	r.GET("/play/inactive", func(c *gin.Context) {
		ga.PlayUseCase.GetAllByInactive(c.Request.Context(), c)
	})
	r.GET("/play/:id", func(c *gin.Context) {
		ga.PlayUseCase.GetByID(c.Request.Context(), c)
	})
	r.POST("/play", func(c *gin.Context) {
		ga.PlayUseCase.Create(c.Request.Context(), c)
	})
	r.PUT("/play/:id", func(c *gin.Context) {
		ga.PlayUseCase.Update(c.Request.Context(), c)
	})
	r.DELETE("/play/:id", func(c *gin.Context) {
		ga.PlayUseCase.Delete(c.Request.Context(), c)
	})
	r.GET("/play/name/:name", func(c *gin.Context) {
		ga.PlayUseCase.GetByName(c.Request.Context(), c)
	})

	r.GET("/card/:id", func(c *gin.Context) {
		ga.CardUseCase.GetByID(c.Request.Context(), c)
	})
	r.POST("/card/:id", func(c *gin.Context) {
		ga.CardUseCase.Create(c.Request.Context(), c)
	})
	r.PUT("/card/:id", func(c *gin.Context) {
		ga.CardUseCase.Update(c.Request.Context(), c)
	})

	r.GET("/cardplay", func(c *gin.Context) {
		ga.CardPlayUseCase.GetAll(c.Request.Context(), c)
	})
	r.GET("/cardplay/inactive", func(c *gin.Context) {
		ga.CardPlayUseCase.GetAllByInactive(c.Request.Context(), c)
	})
	r.GET("/cardplay/:id", func(c *gin.Context) {
		ga.CardPlayUseCase.GetByID(c.Request.Context(), c)
	})

	r.GET("/nation", func(c *gin.Context) {
		ga.NationUseCase.GetAll(c.Request.Context(), c)
	})
	r.GET("/nation/:id", func(c *gin.Context) {
		ga.NationUseCase.GetByID(c.Request.Context(), c)
	})
	r.POST("/nation", func(c *gin.Context) {
		ga.NationUseCase.Create(c.Request.Context(), c)
	})
	r.PUT("/nation/:id", func(c *gin.Context) {
		ga.NationUseCase.Update(c.Request.Context(), c)
	})

	r.GET("/photo/:id", func(c *gin.Context) {
		ga.PhotoUseCase.GetByIDPlay(c.Request.Context(), c)
	})
	r.POST("/photo/:id", func(c *gin.Context) {
		ga.PhotoUseCase.Create(c.Request.Context(), c)
	})
	r.PUT("/photo/:id", func(c *gin.Context) {
		ga.PhotoUseCase.Update(c.Request.Context(), c)
	})
	r.DELETE("/photo/:id", func(c *gin.Context) {
		ga.PhotoUseCase.Delete(c.Request.Context(), c)
	})

	r.GET("/position", func(c *gin.Context) {
		ga.PositionUseCase.GetAll(c.Request.Context(), c)
	})
	r.GET("/position/:id", func(c *gin.Context) {
		ga.PositionUseCase.GetByID(c.Request.Context(), c)
	})
	r.POST("/position", func(c *gin.Context) {
		ga.PositionUseCase.Create(c.Request.Context(), c)
	})
	r.PUT("/position/:id", func(c *gin.Context) {
		ga.PositionUseCase.Update(c.Request.Context(), c)
	})
	r.DELETE("/position/:id", func(c *gin.Context) {
		ga.PositionUseCase.Delete(c.Request.Context(), c)
	})

	r.POST("/attributes", func(c *gin.Context) {
		ga.Attributes.Create(c.Request.Context(), c)
	})
	r.GET("/attributes", func(c *gin.Context) {
		ga.Attributes.GetAll(c.Request.Context(), c)
	})
	r.GET("/attributes/:id", func(c *gin.Context) {
		ga.Attributes.GetByIDAttributes(c.Request.Context(), c)
	})

	return r
}
