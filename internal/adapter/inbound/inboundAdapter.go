package rest

import (
	"go-projects/envi-monitor/internal/adapter/inbound/rest"

	"go.uber.org/dig"
)

type Inbound struct {
	dig.In

	Inbound rest.Inbound
}
