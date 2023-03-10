// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"launcher_micro/service/env/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Example},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/env/game/new",
					Handler: GameHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/env/new",
					Handler: EnvHandler(serverCtx),
				},
			}...,
		),
	)
}
