// @tg version=0.0.1
// @tg backend=auto-api
// @tg title=`Auto API`
// @tg servers=
//
//go:generate tg transport --services . --out ../../../internal/transport/jsonRPC/externalapi --outSwagger ../../../swaggers/externalapi/swagger.yaml
package externalapi

import (
	"context"

	publicModels "github.com/ARM-auto-place/auto/pkg/models"
	"github.com/google/uuid"
)

// AutoApi
// @tg http-server metrics log
// @tg http-prefix=/api/auto
// @tg 200=github.com/ARM-auto-place/auto/swaggers/externalapi/models:Resp200
// @tg 400=github.com/ARM-auto-place/auto/swaggers/externalapi/models:Err400
// @tg 401=github.com/ARM-auto-place/auto/swaggers/externalapi/models:Err401
// @tg 402=github.com/ARM-auto-place/auto/swaggers/externalapi/models:Err402
// @tg 403=github.com/ARM-auto-place/auto/swaggers/externalapi/models:Err403
// @tg 404=github.com/ARM-auto-place/auto/swaggers/externalapi/models:Err404
// @tg 405=github.com/ARM-auto-place/auto/swaggers/externalapi/models:Err405
// @tg 409=github.com/ARM-auto-place/auto/swaggers/externalapi/models:Err409
// @tg 422=github.com/ARM-auto-place/auto/swaggers/externalapi/models:Err422
// @tg 429=github.com/ARM-auto-place/auto/swaggers/externalapi/models:Err429
// @tg 500=github.com/ARM-auto-place/auto/swaggers/externalapi/models:Err500
type AutoApi interface {
	// GetCars ...
	// @tg http-method=GET
	// @tg http-path=/v1/cars
	// @tg http-headers=userID|X-User-Id
	// @tg uuidPackage=github.com/google/uuid
	// @tg http-response=github.com/ARM-auto-place/auto/internal/transport/custom-handels:GetCars
	// @tg summary=`Получение списка автомобилей`
	// @tg desc=`Получение списка автомобилей`
	GetCars(ctx context.Context, userID uuid.UUID) (cars []publicModels.Cars, err error)
	// GetCarInfo ...
	// @tg http-method=GET
	// @tg http-path=/v1/car/info
	// @tg http-headers=userID|X-User-Id
	// @tg http-args=carID|carID
	// @tg uuidPackage=github.com/google/uuid
	// @tg http-response=github.com/ARM-auto-place/auto/internal/transport/custom-handels:GetCarInfo
	// @tg summary=`Получение полной информации об автомобиле`
	// @tg desc=`Получение полной информации об автомобиле`
	GetCarInfo(ctx context.Context, userID uuid.UUID, carID uuid.UUID) (cars publicModels.CarInfo, err error)

	// CreateAuto ...
	// @tg http-method=POST
	// @tg http-path=/v1/car
	// @tg http-headers=userID|X-User-Id
	// @tg http-args=carID|carID
	// @tg uuidPackage=github.com/google/uuid
	// @tg http-response=github.com/ARM-auto-place/auto/internal/transport/custom-handels:CreateAuto
	// @tg summary=`создание обьявления`
	// @tg desc=`создание обьявления`
	CreateAuto(ctx context.Context, userID uuid.UUID, params publicModels.CarInfo) (carID uuid.UUID, err error)
}
