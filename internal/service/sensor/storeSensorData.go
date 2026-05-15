package sensor

import (
	"context"
	"time"

	outboundEntity "go-projects/envi-monitor/internal/adapter/outbound/entity"
	sensorEntity "go-projects/envi-monitor/internal/service/entity/sensor"
)

func (s service) StoreSensorData(ctx context.Context, req sensorEntity.StoreSensorDataRequest) error {
	err := s.Outbound.Cache.Sensor.SetSensorData(ctx, outboundEntity.Sensor{
		Temperature: req.Temperature,
		Humidity:    req.Humidity,
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}

type IStoreSensorData interface {
	StoreSensorData(ctx context.Context, req sensorEntity.StoreSensorDataRequest) error
}
