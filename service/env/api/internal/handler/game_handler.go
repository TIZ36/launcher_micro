package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"launcher_micro/service/env/api/internal/logic"
	"launcher_micro/service/env/api/internal/svc"
	"launcher_micro/service/env/api/internal/types"
)

func GameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NewGameReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGameLogic(r.Context(), svcCtx)
		resp, err := l.Game(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
