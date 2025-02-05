package middleware

import (
	"net/http"
	"picture/common/response"
	"strings"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			httpx.WriteJson(w, http.StatusUnauthorized, response.Error(401, "未登录"))
			return
		}

		parts := strings.SplitN(authorization, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			httpx.WriteJson(w, http.StatusUnauthorized, response.Error(401, "非法的认证头"))
			return
		}

		next(w, r)
	}
}
