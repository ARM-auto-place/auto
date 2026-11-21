package service

import (
	"context"

	"github.com/ARM-auto-place/auto/internal/storage/postgres"
	"github.com/ARM-auto-place/auto/pkg/interfaces/internalapi"
	publicModels "github.com/ARM-auto-place/auto/pkg/models"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type InternalAuthAPIService struct {
	logger  zerolog.Logger
	storage postgres.AutoStorage
}

func NewInternalPayAPIService(
	logger zerolog.Logger,
	storage postgres.AutoStorage,
) internalapi.InternalAutoAPI {
	return &InternalAuthAPIService{
		logger:  logger,
		storage: storage,
	}
}
func (i *InternalAuthAPIService) GetListMyCars(
	ctx context.Context,
	userID uuid.UUID,
) (cars []publicModels.Cars, err error) {
	i.logger.Info().Msg("call get list my cars")
	cars, err = i.storage.GetListMyCars(ctx, userID)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func (i *InternalAuthAPIService) GetListMyFavoriteCars(
	ctx context.Context,
	userID uuid.UUID,
) (cars []publicModels.Cars, err error) {
	i.logger.Info().Msg("call get list my favorite cars")
	cars, err = i.storage.GetListMyFavoriteCars(ctx, userID)
	if err != nil {
		return nil, err
	}
	return cars, nil
}
