package main

import (
	"rachao/config"
	"rachao/infra/messaging"
	"rachao/infra/repositories"
	"rachao/internal/core/adapters"
	"rachao/internal/core/usecase"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	db := config.InitDatabase(cfg.DbSource)
	defer db.Close()

	rabbitMQChannel := config.InitRabbitMQ(cfg.Messaging, cfg.MessagingChannel)
	defer rabbitMQChannel.Close()

	repoPlay := repositories.PlayRepository{DB: db}
	repoCard := repositories.CardRepository{DB: db}
	repoCardPlay := repositories.CardPlayRepository{DB: db}
	repoNation := repositories.NationRepository{DB: db}
	repoPhoto := repositories.PhotoRepository{DB: db}
	repoPosition := repositories.PositionRepository{DB: db}
	repoAttribute := repositories.AttributesRepository{DB: db}
	repoOverall := repositories.OverallRepository{DB: db}
	repoModality := repositories.ModalityRepository{DB: db}
	rabbitmq := messaging.RabbitMQ{Channel: rabbitMQChannel, Exchange: cfg.MessagingChannel}

	healthzUseCase := &usecase.HealthzUseCase{}

	overallUseCase := usecase.NewOverallUseCase(&rabbitmq, &repoOverall, db, logger)
	playUseCase := usecase.NewPlayUseCase(&repoPlay, db, logger)
	cardUseCase := usecase.NewCardUseCase(&repoCard, &repoCardPlay, &repoAttribute, &rabbitmq, db, logger)
	cardPlayUseCase := usecase.NewCardPlayUseCase(&repoCardPlay, db, logger)
	nationUseCase := usecase.NewNationUseCase(&repoNation, db, logger)
	photoUseCase := usecase.NewPhotoUseCase(&repoPhoto, db, logger)
	positionUseCase := usecase.NewPositionUseCase(&repoPosition, db, logger)
	attributesUseCase := usecase.NewAttributesUseCase(&repoAttribute, &repoPosition, db, logger)
	modalitiesUseCase := usecase.NewModalityUseCase(&repoModality, db, logger)

	go overallUseCase.Start()

	GinAdapater := adapters.NewGinAdapter(
		healthzUseCase,
		playUseCase,
		cardUseCase,
		cardPlayUseCase,
		nationUseCase,
		photoUseCase,
		positionUseCase,
		attributesUseCase,
		modalitiesUseCase,
	)

	r := GinAdapater.SetupRouter()
	r.Run(cfg.Port)

}
