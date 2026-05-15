package inbound

import (
	"go-projects/envi-monitor/internal/adapter/inbound/mqtt"
	"go-projects/envi-monitor/internal/adapter/inbound/rest"

	"go.uber.org/dig"
)

type Inbound struct {
	dig.In

	Rest rest.Inbound
	Mqtt mqtt.Inbound
}

func Register(container *dig.Container) error {
	return mqtt.Register(container)
}
