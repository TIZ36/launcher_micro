package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/status"
	"launcher_micro/service/user/api/internal/svc"
	"launcher_micro/service/user/api/internal/types"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq, w http.ResponseWriter, r *http.Request) (resp *types.UserLoginResp, err error) {

	var userLoginCache types.IdaasUserLoginCache

	phoneNum := req.PhoneNum

	userLoginCacheBytes, e := l.svcCtx.LocalCache.Get(phoneNum)

	if e != nil {
		UserLoginRedirect(l, w, r)
		return
	}

	e = json.Unmarshal(userLoginCacheBytes, &userLoginCache)

	if e != nil {
		return nil, status.Error(http.StatusInternalServerError, "parse user login cache error")
	}

	// 还没有过期，可以直接登陆
	if userLoginCache.ExpiresAt > int64(time.Now().Second()) {
		logx.Info("userLoginCache is valid yet: ", userLoginCache)
		return &types.UserLoginResp{
			Idtoken:   userLoginCache.IdToken,
			ExpiresIn: userLoginCache.ExpiresIn,
			ExpiresAt: userLoginCache.ExpiresAt,
		}, nil
	}

	UserLoginRedirect(l, w, r)

	return nil, nil
}

func UserLoginRedirect(l *UserLoginLogic, w http.ResponseWriter, r *http.Request) {

	authorizationUrl := fmt.Sprintf(
		"%s?client_id=%s&redirect_uri=%s&response_type=%s&scope=%s&state=%s",
		l.svcCtx.Config.SSO.AuthorizationEndPoint,
		l.svcCtx.Config.SSO.ClientId,
		l.svcCtx.Config.SSO.RedirectUrl,
		l.svcCtx.Config.SSO.ResponseType,
		l.svcCtx.Config.SSO.Scope,
		l.svcCtx.Config.SSO.State,
	)

	logx.Infof("authorization url is", authorizationUrl)

	http.Redirect(w, r, authorizationUrl, http.StatusFound)
}
