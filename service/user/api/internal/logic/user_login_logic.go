package logic

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"launcher_micro/service/user/api/internal/svc"
	"launcher_micro/service/user/api/internal/types"

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

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	// todo: add your logic here and delete this line

	return &types.UserLoginResp{
		Idtoken: "33",
		Expire:  strconv.FormatInt(300, 10),
	}, nil
}

func (l *UserLoginLogic) UserLoginRedirect(req *types.UserLoginReq, w http.ResponseWriter, r *http.Request) {

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
