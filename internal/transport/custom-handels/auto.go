package custom_handlers

import (
	interfaces "github.com/ARM-auto-place/auto/pkg/interfaces/externalapi"
	publicModels "github.com/ARM-auto-place/auto/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const authAPIServiceName = "auto-api"

func GetCars(ctx *fiber.Ctx, svc interfaces.AutoApi, userID uuid.UUID) error {
	return HandelResponse(ctx, "get", "/v1/car/info", "GetCars", authAPIServiceName, map[string]interface{}{
		"userID": userID,
	}, func() (interface{}, error) {
		return svc.GetCars(ctx.UserContext(), userID)
	})
}
func GetCarInfo(ctx *fiber.Ctx, svc interfaces.AutoApi, userID uuid.UUID, carID uuid.UUID) error {
	return HandelResponse(ctx, "get", "/v1/car/info", "GetCarInfo", authAPIServiceName, map[string]interface{}{
		"userID": userID,
		"carID":  carID,
	}, func() (interface{}, error) {
		return svc.GetCarInfo(ctx.UserContext(), userID, carID)
	})
}
func CreateAuto(ctx *fiber.Ctx, svc interfaces.AutoApi, userID uuid.UUID, params publicModels.CarInfo) error {
	return HandelResponse(ctx, "get", "/v1/car/info", "GetCarInfo", authAPIServiceName, map[string]interface{}{
		"userID": userID,
		"params": params,
	}, func() (interface{}, error) {
		return svc.CreateAuto(ctx.UserContext(), userID, params)
	})
}
