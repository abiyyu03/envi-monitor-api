package user

import (
	"go-projects/envi-monitor/internal/service"

	"go.uber.org/dig"
)

type Handler struct {
	dig.In

	Service service.Service
}
