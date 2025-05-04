package usecase

import (
	"database/sql"
	"encoding/json"
	"rachao/infra/messaging"
	"rachao/infra/repositories"
	"rachao/internal/core/domain"

	"go.uber.org/zap"
)

type OverallUseCase struct {
	Messaging         messaging.MessagePublisherInterface
	OverallRepository repositories.OverallRepositoryInterface
	db                *sql.DB
	Logger            *zap.Logger
}

func NewOverallUseCase(
	Messaging messaging.MessagePublisherInterface,
	overallRepository repositories.OverallRepositoryInterface,
	db *sql.DB,
	logger *zap.Logger,

) *OverallUseCase {
	return &OverallUseCase{
		Messaging:         Messaging,
		OverallRepository: overallRepository,
		db:                db,
		Logger:            logger,
	}
}

func (uc *OverallUseCase) Start() {

	handler := func(message string) {

		uc.Logger.Info("Received message", zap.String("message", message))

		err := uc.overallCreateUpdate(message)
		if err != nil {
			uc.Logger.Error("Error processing message", zap.Error(err))
			return
		}

		uc.Logger.Info("Message processed successfully")

	}

	err := uc.Messaging.Consumer(handler, "overall")
	if err != nil {
		uc.Logger.Error("Error starting consumer", zap.Error(err))
		return
	}

}

func (uc *OverallUseCase) overallCreateUpdate(message string) error {
	var overallRequest domain.OverallRequest
	err := json.Unmarshal([]byte(message), &overallRequest)
	if err != nil {
		uc.Logger.Error("Error unmarshalling message", zap.Error(err))
		return err
	}

	exist, err := uc.OverallRepository.Exists(overallRequest.IDPlay)
	if err != nil && err == sql.ErrNoRows {
		uc.Logger.Error("Error checking overall existence", zap.Error(err))
	}
	if exist {
		updateErr := uc.OverallRepository.Update(overallRequest, overallRequest.IDPlay)
		if updateErr != nil {
			uc.Logger.Error("Error updating overall", zap.Error(updateErr))
			return updateErr
		}
	}

	if !exist {
		_, Err := uc.OverallRepository.Create(overallRequest)
		if Err != nil {
			uc.Logger.Error("Error creating overall", zap.Error(Err))
			return Err
		}
	}

	uc.Logger.Info("Overall created/updated successfully", zap.String("id", overallRequest.IDPlay.String()))
	return nil
}
