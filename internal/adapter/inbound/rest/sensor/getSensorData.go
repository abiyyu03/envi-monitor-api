package sensor

import (
	"context"
	"go-projects/envi-monitor/internal/adapter/inbound/rest/sensor/entity"

	sensorEntity "go-projects/envi-monitor/internal/service/entity/sensor"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetSensorData(fctx *fiber.Ctx) error {
	ctx := context.Background()
	var (
		response entity.BaseResponse
	)

	result, err := h.Service.Sensor.GetSensorData(ctx, sensorEntity.GetSensorDataRequest{})
	if err != nil {
		return err
	}

	return fctx.Status(fiber.StatusOK).JSON(response.ToResponse(
		"Sensor data retrieved successfully",
		fiber.StatusOK,
		result,
		nil,
	))
}
