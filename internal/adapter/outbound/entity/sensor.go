package entity

import "time"

type Sensor struct {
	Temperature   float64   `json:"temperature"`
	Humidity      float64   `json:"humidity"`
	AirQualityPPM float64   `json:"air_quality_ppm"`
	UpdatedAt     time.Time `json:"updated_at"`
}
