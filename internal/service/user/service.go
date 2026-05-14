package user

import (
	"go-projects/envi-monitor/internal/adapter/outbound"
	"go-projects/envi-monitor/pkg"
)

type UserService interface {
	ICreateUser
	IGetAll
}

type service struct {
	outbound.Outbound
	cache      *pkg.Redis
	repository *pkg.SQL
}

func New(
	outbound outbound.Outbound,
	pkg pkg.Package,
) UserService {
	return &service{
		Outbound:   outbound,
		cache:      pkg.Cache,
		repository: pkg.DB,
	}
}
