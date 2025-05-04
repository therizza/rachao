package config

import (
	"database/sql"
	"log"
	"os"
	"rachao/internal/core/constantes"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	Port             string
	DbSource         string
	Messaging        string
	MessagingChannel string
}

func Load() *Config {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	return &Config{
		Port:             os.Getenv(constantes.Port),
		DbSource:         os.Getenv(constantes.DbSource),
		Messaging:        os.Getenv(constantes.Messaging),
		MessagingChannel: os.Getenv(constantes.MessagingChannel),
	}
}

func InitDatabase(dbSource string) *sql.DB {
	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		panic("Error connecting to the database: " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic("Database connection test failed: " + err.Error())
	}

	return db
}

func InitRabbitMQ(messagingEnv string, messagingChannel string) *amqp.Channel {
	conn, err := amqp.Dial(messagingEnv)
	if err != nil {
		panic("Error connecting to RabbitMQ: " + err.Error())
	}

	channel, err := conn.Channel()
	if err != nil {
		panic("Error creating RabbitMQ channel: " + err.Error())
	}

	err = channel.ExchangeDeclare(
		messagingChannel,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
	}
	return channel
}
