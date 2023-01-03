package logic

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/status"
	"net/http"

	"launcher_micro/service/user/api/internal/svc"
	"launcher_micro/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	var userLoginCache types.IdaasUserLoginCache
	// 这里一定是有的
	userCacheBytes, e := l.svcCtx.LocalCache.Get(req.PhoneNum)

	e = json.Unmarshal(userCacheBytes, &userLoginCache)

	if e != nil {
		return nil, status.Error(http.StatusInternalServerError, e.Error())
	}

	return &types.UserInfoResp{
		Code:   http.StatusOK,
		Name:   userLoginCache.Name,
		Number: userLoginCache.PhoneNum,
	}, nil
}
