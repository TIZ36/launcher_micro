package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"launcher_micro/service/env/api/internal/logic"
	"launcher_micro/service/env/api/internal/svc"
	"launcher_micro/service/env/api/internal/types"
)

func EnvHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NewParkEnvReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewEnvLogic(r.Context(), svcCtx)
		resp, err := l.Env(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
