package sensor

import (
	"go-projects/envi-monitor/internal/adapter/outbound"
	"go-projects/envi-monitor/pkg"
)

type SensorService interface {
	IGetSensorData
	IStoreSensorData
}

type service struct {
	outbound.Outbound
	cache      *pkg.Redis
	repository *pkg.SQL
}

func New(
	outbound outbound.Outbound,
	pkg pkg.Package,
) SensorService {
	return &service{
		Outbound:   outbound,
		cache:      pkg.Cache,
		repository: pkg.DB,
	}
}
