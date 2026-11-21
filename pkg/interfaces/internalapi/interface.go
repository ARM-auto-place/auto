//  Ручка получения автомобиля пользователя(часть данных это internal ручка),

//  Ручка вывода списка понравившихся автомобилей(internal ручка)

// Package internal_service
// @tg version=0.0.1
// @tg backend=auto-api
// @tg title=`Auto API`
// @tg servers=
//
//go:generate tg transport --services . --out ../../../internal/transport/jsonRPC/internalapi --outSwagger ../../../swaggers/internalapi/swagger.yaml
//go:generate tg client --services . --outPath ../../client/transport -go

package internalapi

import (
	"context"

	publicModels "github.com/ARM-auto-place/auto/pkg/models"
	"github.com/google/uuid"
)

// InternalPayAPI
// @tg jsonRPC-server log metrics
// @tg 200=github.com/ARM-auto-place/auto/swaggers/internalapi/models:Resp200
// @tg 402=github.com/ARM-auto-place/auto/swaggers/internalapi/models:Err402
// @tg 403=github.com/ARM-auto-place/auto/swaggers/internalapi/models:Err403
// @tg 405=github.com/ARM-auto-place/auto/swaggers/internalapi/models:Err405
// @tg 500=github.com/ARM-auto-place/auto/swaggers/internalapi/models:Err500
type InternalAutoAPI interface {
	// GetListMyCars ...
	// @tg uuidPackage=github.com/google/uuid
	// @tg userID.format=uuid
	// @tg carID.format=uuid
	// @tg summary=`Получить список обьявлений пользователя`
	// @tg desc=`Получить список обьявлений пользователя`
	// @tg log-skip=serviceID
	// @tg 200=github.com/ARM-auto-place/auto/swaggers/internalapi/models:GetListMyCarsResp200
	GetListMyCars(
		ctx context.Context,
		userID uuid.UUID,
	) (cars []publicModels.Cars, err error)
	// GetListMyFavoriteCars ...
	// @tg uuidPackage=github.com/google/uuid
	// @tg userID.format=uuid
	// @tg carID.format=uuid
	// @tg summary=`Получить список обьявлений пользователя`
	// @tg desc=`Получить список обьявлений пользователя`
	// @tg log-skip=serviceID
	// @tg 200=github.com/ARM-auto-place/auto/swaggers/internalapi/models:GetListMyFavoriteCarsResp200
	GetListMyFavoriteCars(
		ctx context.Context,
		userID uuid.UUID,
	) (cars []publicModels.Cars, err error)
}
