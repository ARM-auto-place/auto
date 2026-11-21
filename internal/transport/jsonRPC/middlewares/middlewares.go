package middlewares

import (
	"context"

	"github.com/ARM-auto-place/auto/internal/clients"
	customErrors "github.com/ARM-auto-place/auto/internal/errors"
	"github.com/ARM-auto-place/auto/pkg/interfaces/externalapi"
	publicModels "github.com/ARM-auto-place/auto/pkg/models"
	"github.com/google/uuid"
)

type internalMiddleware struct {
	authAPI     externalapi.AutoApi
	authService clients.AccessRoleClient
}

func NewInternalMiddleware(next externalapi.AutoApi, authService clients.AccessRoleClient) externalapi.AutoApi {
	return &internalMiddleware{
		authAPI:     next,
		authService: authService,
	}
}

func (i *internalMiddleware) GetCars(ctx context.Context, userID uuid.UUID) (cars []publicModels.Cars, err error) {
	role, err := i.authService.GetRole(ctx, userID)
	if err != nil || (role != "admin" && role != "owner") {
		return nil, customErrors.AccessDeniedError()
	}
	return i.authAPI.GetCars(ctx, userID)
}

func (i *internalMiddleware) GetCarInfo(ctx context.Context, userID uuid.UUID, carID uuid.UUID) (cars publicModels.CarInfo, err error) {
	role, err := i.authService.GetRole(ctx, userID)
	if err != nil || (role != "admin" && role != "owner") {
		return publicModels.CarInfo{}, customErrors.AccessDeniedError()
	}
	return i.authAPI.GetCarInfo(ctx, userID, carID)
}

func (i *internalMiddleware) CreateAuto(ctx context.Context, userID uuid.UUID, params publicModels.CarInfo) (carID uuid.UUID, err error) {
	role, err := i.authService.GetRole(ctx, userID)
	if err != nil || (role != "admin" && role != "owner") {
		return uuid.Nil, customErrors.AccessDeniedError()
	}
	return i.authAPI.CreateAuto(ctx, userID, params)
}
