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

type GameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GameLogic {
	return &GameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GameLogic) Game(req *types.NewGameReq) (resp *types.NewGameReply, err error) {
	// todo: add your logic here and delete this line

	NewGame := model.Game{
		GameId:         req.GameId,
		GameName:       req.GameName,
		GameNameAbbrev: req.GameNameAbbrev,
	}

	_, er := l.svcCtx.GameModel.Insert(l.ctx, &NewGame)

	if er != nil {
		return nil, status.Error(http.StatusInternalServerError, er.Error())
	}

	return &types.NewGameReply{
		Code:   http.StatusOK,
		Result: "success",
	}, nil
}
