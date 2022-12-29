package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"launcher_micro/service/env/model"
	"net/http"

	"launcher_micro/service/env/api/internal/svc"
	"launcher_micro/service/env/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnvLogic {
	return &EnvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnvLogic) Env(req *types.NewParkEnvReq) (resp *types.NewParkEnvResp, err error) {
	// todo: add your logic here and delete this line

	var e error
	GameId := req.GameId

	_, e = l.svcCtx.GameModel.FindOne(l.ctx, GameId)

	if e != nil {
		return nil, status.Error(http.StatusBadRequest, e.Error())
	}

	NewEnv := model.ParkEnv{
		ParkEnvId:      req.ParkEnvId,
		GameId:         req.GameId,
		Name:           req.Name,
		UpdateStrategy: req.UpdateStrategy,
	}

	_, e = l.svcCtx.ParkEnvModel.Insert(l.ctx, &NewEnv)

	return &types.NewParkEnvResp{
		Code:   http.StatusOK,
		Result: "success",
	}, nil
}
