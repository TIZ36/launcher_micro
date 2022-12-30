package logic

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc/status"
	"launcher_micro/common/jwt_util"
	"launcher_micro/service/user/api/internal/svc"
	"launcher_micro/service/user/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOauthCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type AuthToken struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ExpiresAt   int64  `json:"expires_at"`
	IdToken     string `json:"id_token"`
}

func NewUserOauthCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOauthCallbackLogic {
	return &UserOauthCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserOauthCallbackLogic) UserOauthCallback(req *types.UserOauthCallBack) (resp *types.UserOauthCallBackResp, err error) {

	// todo: add your logic here and delete this line

	var authToken AuthToken
	authCode := req.Code
	state := req.State

	OriginState := l.svcCtx.Config.SSO.State

	if OriginState != state {
		return nil, status.Error(http.StatusBadRequest, "state code miss match")
	}

	logx.Info("auth_code is ", authCode)
	logx.Info("auth_code is ", state)

	IdTokenReqUrl := l.svcCtx.Config.SSO.TokenEndpoint

	formData := map[string]string{
		"grant_type":    l.svcCtx.Config.SSO.GrantType,
		"code":          authCode,
		"client_id":     l.svcCtx.Config.SSO.ClientId,
		"client_secret": l.svcCtx.Config.SSO.ClientSecret,
		"redirect_uri":  l.svcCtx.Config.SSO.RedirectUrl,
	}

	res, e := l.svcCtx.HttpClient.R().
		SetHeader("Content_Type", "application/x-www-form-urlencoded").
		SetFormData(formData).
		Post(IdTokenReqUrl)

	if e != nil {
		return nil, status.Error(http.StatusBadRequest, e.Error())
	}

	e = json.Unmarshal([]byte(res.String()), &authToken)

	if e != nil {
		return nil, status.Error(http.StatusInternalServerError, e.Error())
	}

	jwt_util.ParseJwtToken(authToken.IdToken)
	// do token req
	return &types.UserOauthCallBackResp{
		IdToken:   authToken.IdToken,
		ExpiresIn: authToken.ExpiresIn,
		ExpiresAt: authToken.ExpiresAt,
	}, nil
}
