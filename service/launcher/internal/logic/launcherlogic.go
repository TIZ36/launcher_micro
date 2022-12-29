package logic

import (
	"context"

	"launcher_micro/service/launcher/internal/svc"
	"launcher_micro/service/launcher/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LauncherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLauncherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LauncherLogic {
	return &LauncherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LauncherLogic) Launcher(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
