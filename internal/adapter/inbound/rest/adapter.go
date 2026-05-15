package rest

import (
	"go-projects/envi-monitor/internal/adapter/inbound/rest/sensor"

	"go.uber.org/dig"
)

type Inbound struct {
	dig.In

	Sensor sensor.Handler
}
