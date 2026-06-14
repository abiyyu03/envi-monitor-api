package sensor

import "time"

type (
	GetSensorDataRequest struct{}

	GetSensorDataResponse struct {
		Temperature   float64   `json:"temperature"`
		Humidity      float64   `json:"humidity"`
		AirQualityPPM float64   `json:"air_quality_ppm"`
		UpdatedAt     time.Time `json:"updated_at"`
	}

	StoreSensorDataRequest struct {
		Temperature   float64
		Humidity      float64
		AirQualityPPM float64
	}
)
