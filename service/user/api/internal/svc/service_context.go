package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"launcher_micro/service/user/api/internal/config"
	"launcher_micro/service/user/api/internal/middleware"
)

type ServiceContext struct {
	Config config.Config

	RoleCheck rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		RoleCheck: middleware.NewRoleCheckMiddleware().Handle,
	}
}
