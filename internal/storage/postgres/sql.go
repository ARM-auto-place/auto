package postgres

import (
	"context"
	_ "embed"
	"errors"
	"fmt"

	publicModels "github.com/ARM-auto-place/auto/pkg/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed sql/selectCars.sql
var sqlSelectCars string

//go:embed sql/selectMyCars.sql
var sqlSelectMyCars string

//go:embed sql/sqlselectMyFavoriteCars.sql
var sqlselectMyFavoriteCars string

//go:embed sql/selectCarInfo.sql
var sqlSelectCarInfo string

//go:embed sql/insertAutoCard.sql
var sqlInsertAutoCard string

func selectCars(ctx context.Context, conn *pgxpool.Pool) ([]publicModels.Cars, error) {
	internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	rows, err := conn.Query(internalCtx, sqlSelectCars)
	if err != nil {
		return nil, fmt.Errorf("postgresql: %w", err)
	}
	defer rows.Close()

	var res []publicModels.Cars
	for rows.Next() {
		var p publicModels.Cars
		err = rows.Scan(
			&p.CarID,
			&p.Mark,
			&p.Model,
			&p.Kilometre,
			&p.Price,
			&p.Images,
		)
		if err != nil {
			return nil, fmt.Errorf("postgresql: %w", err)
		}
		res = append(res, p)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("while processing results row: %w", rows.Err())
	}
	return res, nil
}
func selectMyFavoriteCars(ctx context.Context, conn *pgxpool.Pool, userID uuid.UUID) ([]publicModels.Cars, error) {
	internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	rows, err := conn.Query(internalCtx, sqlselectMyFavoriteCars, userID)
	if err != nil {
		return nil, fmt.Errorf("postgresql: %w", err)
	}
	defer rows.Close()

	var res []publicModels.Cars
	for rows.Next() {
		var p publicModels.Cars
		err = rows.Scan(
			&p.CarID,
			&p.Mark,
			&p.Model,
			&p.Kilometre,
			&p.Price,
			&p.Images,
		)
		if err != nil {
			return nil, fmt.Errorf("postgresql: %w", err)
		}
		res = append(res, p)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("while processing results row: %w", rows.Err())
	}
	return res, nil
}
func selectMyCars(ctx context.Context, conn *pgxpool.Pool, userID uuid.UUID) ([]publicModels.Cars, error) {
	internalCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	rows, err := conn.Query(internalCtx, sqlSelectMyCars, userID)
	if err != nil {
		return nil, fmt.Errorf("postgresql: %w", err)
	}
	defer rows.Close()

	var res []publicModels.Cars
	for rows.Next() {
		var p publicModels.Cars
		err = rows.Scan(
			&p.CarID,
			&p.Mark,
			&p.Model,
			&p.Kilometre,
			&p.Price,
			&p.Images,
		)
		if err != nil {
			return nil, fmt.Errorf("postgresql: %w", err)
		}
		res = append(res, p)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("while processing results row: %w", rows.Err())
	}
	return res, nil
}
func selectCarInfo(ctx context.Context, conn *pgxpool.Pool,  carID uuid.UUID) (publicModels.CarInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	var car publicModels.CarInfo
	err := conn.QueryRow(ctx, sqlSelectCarInfo, carID).Scan(
		&car.CarID,
		&car.SellerID,
		&car.Mark,
		&car.Model,
		&car.BuildYear,
		&car.Power,
		&car.TypePower,
		&car.Owners,
		&car.Condition,
		&car.SteeringWheel,
		&car.Color,
		&car.Descriptions,
		&car.CreatedAt,
		&car.UpdatedAt,
		&car.Release,
		&car.Images,
		&car.VIN,
		&car.Kilometre,
		&car.Price,
		&car.IsUserCar,
		&car.IsFavorite,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return publicModels.CarInfo{}, fmt.Errorf("not found")
		}
		return publicModels.CarInfo{}, fmt.Errorf("GetUserByEmail: postgresql: %w", err)
	}
	return car, nil
}
func insertAutoCard(ctx context.Context, conn *pgxpool.Pool, userID uuid.UUID, params publicModels.CarInfo) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("could not start transaction: %w", err)
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
		var carID uuid.UUID
	err = tx.QueryRow(
		ctx,
		sqlInsertAutoCard,
		userID,
		params.Mark,
		params.Model,
		params.BuildYear,
		params.Power,
		params.TypePower,
		params.Owners,
		params.Condition,
		params.SteeringWheel,
		params.Color,
		params.Descriptions,
		params.CreatedAt,
		params.UpdatedAt,
		params.Release,
		params.Images,
		params.VIN,
		params.Kilometre,
		params.Price,
	).Scan(&carID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("insert car: postgresql: %w", err)
	}
	if err = tx.Commit(ctx); err != nil {
		return uuid.Nil, fmt.Errorf("commit car insert: %w", err)
	}
	return carID, nil
}
