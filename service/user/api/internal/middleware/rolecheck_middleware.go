package middleware

import (
	"github.com/allegro/bigcache/v3"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"launcher_micro/service/user/api/internal/config"
	"launcher_micro/service/user/api/internal/types"
	"net/http"
)

type RoleCheckMiddleware struct {
	GlobalConfig config.Config
	LocalCache   *bigcache.BigCache
}

func NewRoleCheckMiddleware(config config.Config, cache *bigcache.BigCache) *RoleCheckMiddleware {
	return &RoleCheckMiddleware{
		GlobalConfig: config,
		LocalCache:   cache,
	}
}

func (m *RoleCheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		_, e := m.LocalCache.Get(req.PhoneNum)
		if e != nil {
			httpx.Error(w, status.Error(http.StatusBadRequest, "user not login yet"))
			return
		}

		next(w, r)
	}
}
