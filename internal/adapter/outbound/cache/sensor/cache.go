package sensor

import "go-projects/envi-monitor/pkg"

type Cache interface {
	ISetSensorData
	IGetSensorData
}

type sensorCache struct {
	pkg pkg.Package
}

func New(pkg pkg.Package) Cache {
	return &sensorCache{pkg: pkg}
}
