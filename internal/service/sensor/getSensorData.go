package sensor

import (
	"context"

	sensorEntity "go-projects/envi-monitor/internal/service/entity/sensor"
)

func (s service) GetSensorData(ctx context.Context, req sensorEntity.GetSensorDataRequest) (sensorEntity.GetSensorDataResponse, error) {
	var response sensorEntity.GetSensorDataResponse

	sensorData, err := s.Outbound.Cache.Sensor.GetSensorData(ctx)
	if err != nil {
		return response, err
	}

	response = sensorEntity.GetSensorDataResponse{
		Temperature:   sensorData.Temperature,
		Humidity:      sensorData.Humidity,
		AirQualityPPM: sensorData.AirQualityPPM,
		UpdatedAt:     sensorData.UpdatedAt,
	}

	return response, nil
}

type IGetSensorData interface {
	GetSensorData(ctx context.Context, req sensorEntity.GetSensorDataRequest) (sensorEntity.GetSensorDataResponse, error)
}
