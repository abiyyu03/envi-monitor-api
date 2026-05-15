package cache

import (
	"go-projects/envi-monitor/internal/adapter/outbound/cache/sensor"
	"go-projects/envi-monitor/internal/adapter/outbound/cache/user"

	"go.uber.org/dig"
)

type Cache struct {
	dig.In

	User   user.Cache
	Sensor sensor.Cache
}

func Register(container *dig.Container) error {
	if err := container.Provide(user.New); err != nil {
		return err
	}

	if err := container.Provide(sensor.New); err != nil {
		return err
	}

	return nil
}
