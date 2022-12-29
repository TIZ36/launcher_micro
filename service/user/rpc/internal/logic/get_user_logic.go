package logic

import (
	"context"

	"launcher_micro/service/user/rpc/internal/svc"
	"launcher_micro/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdReq) (*user.UserInfoReply, error) {
	// todo: add your logic here and delete this line

	one, e := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)

	if e != nil {
		logx.Error("find err")
		return nil, e
	}

	return &user.UserInfoReply{
		Id:     one.Id,
		Name:   one.Name,
		Gender: one.Gender,
		Number: "123",
	}, nil
}
