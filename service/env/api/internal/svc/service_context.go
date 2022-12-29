package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"launcher_micro/service/env/api/internal/config"
	"launcher_micro/service/env/api/internal/middleware"
	"launcher_micro/service/env/model"
	"launcher_micro/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	Example rest.Middleware

	UserRpc userclient.User

	model.GameModel
	model.ParkEnvModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		Example:      middleware.NewExampleMiddleware().Handle,
		UserRpc:      userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		GameModel:    model.NewGameModel(conn, c.CacheRedis),
		ParkEnvModel: model.NewParkEnvModel(conn, c.CacheRedis),
	}
}
