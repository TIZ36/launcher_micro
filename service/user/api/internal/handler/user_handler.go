package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"launcher_micro/service/user/api/internal/logic"
	"launcher_micro/service/user/api/internal/svc"
	"launcher_micro/service/user/api/internal/types"
)

func UserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.User(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
