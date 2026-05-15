package di

import (
	"go-projects/envi-monitor/internal/adapter/inbound"
	"go-projects/envi-monitor/internal/adapter/outbound"
	"go-projects/envi-monitor/internal/service"
	"go-projects/envi-monitor/pkg"

	"go.uber.org/dig"
)

func Container() (*dig.Container, error) {
	container := dig.New()

	if err := pkg.Register(container); err != nil {
		return nil, err
	}

	if err := outbound.Register(container); err != nil {
		return nil, err
	}

	if err := service.Register(container); err != nil {
		return nil, err
	}

	if err := inbound.Register(container); err != nil {
		return nil, err
	}

	return container, nil
}
