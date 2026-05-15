package mqtt

import (
	"go-projects/envi-monitor/internal/adapter/inbound/mqtt/sensor"

	"go.uber.org/dig"
)

type Inbound struct {
	dig.In

	Sensor *sensor.Subscriber
}

func Register(container *dig.Container) error {
	if err := container.Provide(sensor.New); err != nil {
		return err
	}

	return nil
}
