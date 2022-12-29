package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"launcher_micro/service/env/api/internal/config"
	"launcher_micro/service/env/api/internal/middleware"
	"launcher_micro/service/env/model"
)

type ServiceContext struct {
	Config config.Config

	Example rest.Middleware

	model.GameModel
	model.ParkEnvModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		Example:      middleware.NewExampleMiddleware().Handle,
		GameModel:    model.NewGameModel(conn, c.CacheRedis),
		ParkEnvModel: model.NewParkEnvModel(conn, c.CacheRedis),
	}
}
