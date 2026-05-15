package sensor

import "go-projects/envi-monitor/pkg"

type Repository interface {
	IStoreSensorData
}

type sensor struct {
	pkg pkg.Package
}

func New(pkg pkg.Package) Repository {
	return &sensor{pkg: pkg}
}
