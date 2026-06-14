package entity

import sensorEntity "go-projects/envi-monitor/internal/service/entity/sensor"

type SensorPayload struct {
	Temperature   float64 `json:"temperature"`
	AirQualityPPM float64 `json:"air_quality_ppm"`
	Humidity      float64 `json:"humidity"`
}

func (p SensorPayload) ToServiceEntity() sensorEntity.StoreSensorDataRequest {
	return sensorEntity.StoreSensorDataRequest{
		Temperature:   p.Temperature,
		AirQualityPPM: p.AirQualityPPM,
		Humidity:      p.Humidity,
	}
}
