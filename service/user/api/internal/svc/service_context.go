package svc

import (
	"context"
	"github.com/allegro/bigcache/v3"
	resty "github.com/go-resty/resty/v2"
	"github.com/zeromicro/go-zero/rest"
	"launcher_micro/service/user/api/internal/config"
	"launcher_micro/service/user/api/internal/middleware"
	"time"
)

type ServiceContext struct {
	Config    config.Config
	RoleCheck rest.Middleware
	// http client: resty
	HttpClient *resty.Client
	// in-memory cache: bigcache
	LocalCache *bigcache.BigCache
}

func NewServiceContext(c config.Config) *ServiceContext {
	httpClient := resty.New()
	cache, e := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))

	if e != nil {
		return nil
	}

	return &ServiceContext{
		Config:     c,
		RoleCheck:  middleware.NewRoleCheckMiddleware(c, cache).Handle,
		HttpClient: httpClient,
		LocalCache: cache,
	}
}
