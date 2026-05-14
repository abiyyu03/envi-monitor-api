package user

import "go-projects/envi-monitor/pkg"

type Cache interface {
	IGetAllUser
	ISetAllUser
}

type userCache struct {
	Package pkg.Package
}

func New(pkg pkg.Package) Cache {
	return &userCache{
		Package: pkg,
	}
}
