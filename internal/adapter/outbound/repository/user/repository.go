package user

import "go-projects/envi-monitor/pkg"

type Repository interface {
	ICreateUser
	IGetAll
}

type user struct {
	Package pkg.Package
}

func New(pkg pkg.Package) Repository {
	return &user{
		Package: pkg,
	}
}
