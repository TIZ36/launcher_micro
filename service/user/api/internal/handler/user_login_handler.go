package handler

import (
	"launcher_micro/service/user/api/internal/logic"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"launcher_micro/service/user/api/internal/svc"
	"launcher_micro/service/user/api/internal/types"
)

func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLoginLogic(r.Context(), svcCtx)
		l.UserLoginRedirect(&req, w, r)
	}
}
