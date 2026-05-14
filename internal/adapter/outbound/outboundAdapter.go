package outbound

import (
	"go-projects/envi-monitor/internal/adapter/outbound/cache"
	"go-projects/envi-monitor/internal/adapter/outbound/repository"

	"go.uber.org/dig"
)

type Outbound struct {
	dig.In

	Cache      cache.Cache
	Repository repository.Repository
}

func Register(container *dig.Container) error {
	if err := repository.Register(container); err != nil {
		return err
	}
	if err := cache.Register(container); err != nil {
		return err
	}

	return nil
}
