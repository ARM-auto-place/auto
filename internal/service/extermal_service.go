package service

import (
	"context"

	"github.com/ARM-auto-place/auto/internal/storage/postgres"
	"github.com/ARM-auto-place/auto/pkg/interfaces/externalapi"
	publicModels "github.com/ARM-auto-place/auto/pkg/models"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type authAPIService struct {
	logger  zerolog.Logger
	storage postgres.AutoStorage
}

func NewPayAPIService(
	logger zerolog.Logger,
	storage postgres.AutoStorage,
) externalapi.AutoApi {
	return &authAPIService{
		logger:  logger,
		storage: storage,
	}
}
func (a *authAPIService) GetCars(ctx context.Context, userID uuid.UUID) (cars []publicModels.Cars, err error) {
	a.logger.Info().Msg("call get cars")
	cars, err = a.storage.GetCars(ctx)
	if err != nil {
		return nil, err
	}
	return cars, nil
}
func (a *authAPIService) GetCarInfo(ctx context.Context, userID uuid.UUID, carID uuid.UUID) (car publicModels.CarInfo, err error) {
	a.logger.Info().Msg("call get car info")
	car, err = a.storage.GetCarInfo(ctx, carID)
	if err != nil {
		return publicModels.CarInfo{}, err
	}
	// TODO: надо добавить поход в userClient чтобы получить данные о владельце
	// IsUserCar   ,

	return car, nil
}
func (a *authAPIService) CreateAuto(ctx context.Context, userID uuid.UUID, params publicModels.CarInfo) (carID uuid.UUID, err error) {
	a.logger.Info().Msg("call create auto card")
	carID, err = a.storage.CreateAuto(ctx, userID, params)
	if err != nil {
		return uuid.Nil, err
	}
	return carID, nil
}
