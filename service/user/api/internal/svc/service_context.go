package svc

import (
	// Import resty into your code and refer it as `resty`.
	resty "github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/rest"
	"launcher_micro/service/user/api/internal/config"
	"launcher_micro/service/user/api/internal/middleware"
)

type ServiceContext struct {
	Config     config.Config
	RoleCheck  rest.Middleware
	HttpClient *resty.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	httpClient := resty.New()
	return &ServiceContext{
		Config:     c,
		RoleCheck:  middleware.NewRoleCheckMiddleware().Handle,
		HttpClient: httpClient,
	}
}
