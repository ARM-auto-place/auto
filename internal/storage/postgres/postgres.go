package postgres

import (
	"context"
	"time"

	publicModels "github.com/ARM-auto-place/auto/pkg/models"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

var defaultTimeout = 5 * time.Second

type AutoStorage interface {
	GetCars(ctx context.Context) (cars []publicModels.Cars, err error)
	GetCarInfo(ctx context.Context, carID uuid.UUID) (cars publicModels.CarInfo, err error)
	CreateAuto(ctx context.Context, userID uuid.UUID, params publicModels.CarInfo) (carID uuid.UUID, err error)
	GetListMyFavoriteCars(ctx context.Context, userID uuid.UUID) (cars []publicModels.Cars, err error)
	GetListMyCars(ctx context.Context, userID uuid.UUID) (cars []publicModels.Cars, err error)
}
type storage struct {
	connectManager ConnectManager
	logger         zerolog.Logger
}

func NewAutoStorage(logger zerolog.Logger, connectManager ConnectManager) AutoStorage {
	return &storage{
		connectManager: connectManager,
		logger:         logger,
	}
}

func (s *storage) GetCars(ctx context.Context) (cars []publicModels.Cars, err error) {
	conn := s.connectManager.GetReplicaConn()
	return selectCars(ctx, conn)
}
func (s *storage) GetListMyFavoriteCars(ctx context.Context, userID uuid.UUID) (cars []publicModels.Cars, err error) {
	conn := s.connectManager.GetReplicaConn()
	return selectMyFavoriteCars(ctx, conn, userID)
}
func (s *storage) GetListMyCars(ctx context.Context, userID uuid.UUID) (cars []publicModels.Cars, err error) {
	conn := s.connectManager.GetReplicaConn()
	return selectMyCars(ctx, conn, userID)
}
func (s *storage) GetCarInfo(ctx context.Context, carID uuid.UUID) (car publicModels.CarInfo, err error) {
	conn := s.connectManager.GetReplicaConn()
	return selectCarInfo(ctx, conn, carID)
}
func (s *storage) CreateAuto(ctx context.Context, userID uuid.UUID, params publicModels.CarInfo) (carID uuid.UUID, err error) {
	conn := s.connectManager.GetMasterConn()
	return insertAutoCard(ctx, conn, userID, params)
}
