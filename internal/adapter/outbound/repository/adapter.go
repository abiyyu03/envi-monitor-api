package repository

import (
	"go-projects/envi-monitor/internal/adapter/outbound/repository/sensor"
	"go-projects/envi-monitor/internal/adapter/outbound/repository/user"

	"go.uber.org/dig"
)

type Repository struct {
	dig.In

	User   user.Repository
	Sensor sensor.Repository
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
