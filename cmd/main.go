package main

import (
	"database/sql"
	"rachao/infra/repositories"
	"rachao/internal/core/adapters"
	"rachao/internal/core/constantes"
	"rachao/internal/core/usecase"

	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	envPath := "../.env"
	err := godotenv.Load(envPath)
	if err != nil {
		panic("Error loading .env file. Ensure the file exists and the path is correct.")
	}

	port := os.Getenv(constantes.Port)
	dbSource := os.Getenv(constantes.DbSource)
	if dbSource == "" {
		panic("DbSource environment variable is not set. Check your .env file.")
	}

	conn, err := sql.Open("postgres", dbSource)
	if err != nil {
		panic("Error connecting to the database: " + err.Error())
	}
	defer conn.Close()
	err = conn.Ping()
	if err != nil {
		panic("Database connection test failed: " + err.Error())
	}

	repoPlay := repositories.PlayRepository{DB: conn}
	repoCard := repositories.CardRepository{DB: conn}
	repoCardPlay := repositories.CardPlayRepository{DB: conn}
	repoNation := repositories.NationRepository{DB: conn}
	repoPhoto := repositories.PhotoRepository{DB: conn}
	repoPosition := repositories.PositionRepository{DB: conn}
	repoAttribute := repositories.AttributesRepository{DB: conn}
	logger, _ := zap.NewProduction()

	healthzUseCase := &usecase.HealthzUseCase{}
	playUseCase := usecase.NewPlayUseCase(&repoPlay, conn, logger)
	cardUseCase := usecase.NewCardUseCase(&repoCard, conn, logger)
	cardPlayUseCase := usecase.NewCardPlayUseCase(&repoCardPlay, conn, logger)
	nationUseCase := usecase.NewNationUseCase(&repoNation, conn, logger)
	photoUseCase := usecase.NewPhotoUseCase(&repoPhoto, conn, logger)
	positionUseCase := usecase.NewPositionUseCase(&repoPosition, conn, logger)
	attributesUseCase := usecase.NewAttributesUseCase(&repoAttribute, conn, logger)

	GinAdapater := adapters.NewGinAdapter(
		healthzUseCase,
		playUseCase,
		cardUseCase,
		cardPlayUseCase,
		nationUseCase,
		photoUseCase,
		positionUseCase,
		attributesUseCase,
	)

	r := GinAdapater.SetupRouter()

	r.Run(port)

}
