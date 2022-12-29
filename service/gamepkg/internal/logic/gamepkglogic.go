package logic

import (
	"context"

	"launcher_micro/service/gamepkg/internal/svc"
	"launcher_micro/service/gamepkg/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GamepkgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGamepkgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GamepkgLogic {
	return &GamepkgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GamepkgLogic) Gamepkg(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
