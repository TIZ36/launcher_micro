package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"launcher_micro/service/gamepkg/internal/logic"
	"launcher_micro/service/gamepkg/internal/svc"
	"launcher_micro/service/gamepkg/internal/types"
)

func GamepkgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGamepkgLogic(r.Context(), svcCtx)
		resp, err := l.Gamepkg(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
