package entity

import "time"

type Sensor struct {
	Temperature float64   `json:"temperature"`
	Humidity    float64   `json:"humidity"`
	UpdatedAt   time.Time `json:"updated_at"`
}
