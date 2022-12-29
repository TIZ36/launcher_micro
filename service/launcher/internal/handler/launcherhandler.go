package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"launcher_micro/service/launcher/internal/logic"
	"launcher_micro/service/launcher/internal/svc"
	"launcher_micro/service/launcher/internal/types"
)

func LauncherHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLauncherLogic(r.Context(), svcCtx)
		resp, err := l.Launcher(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
