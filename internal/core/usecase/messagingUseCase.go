package usecase

import (
	"encoding/json"
	"rachao/infra/messaging"
	"rachao/internal/core/domain"

	"go.uber.org/zap"
)

type MessagingUseCase struct {
	messagePublisherInterface messaging.MessagePublisherInterface
	logger                    *zap.Logger
}

func NewMessagingaUseCase(messagePublisherInterface messaging.MessagePublisherInterface, logger *zap.Logger) *MessagingUseCase {
	return &MessagingUseCase{
		messagePublisherInterface: messagePublisherInterface,
		logger:                    logger,
	}
}

func (uc *MessagingUseCase) Publish(overallRequest domain.OverallBodyRequest) error {
	serializedRequest, err := json.Marshal(overallRequest)
	if err != nil {
		uc.logger.Error("Failed to serialize overallRequest", zap.Error(err))
		return err
	}

	card := "card." + overallRequest.Card.IDPlay.String()
	err = uc.messagePublisherInterface.Publish("rachao", card, []byte(serializedRequest))
	if err != nil {
		uc.logger.Error("Failed to publish message", zap.Error(err))
		return err
	}
	uc.logger.Info("Message published successfully", zap.String("card", card), zap.String("message", string(serializedRequest)))
	return nil
}
