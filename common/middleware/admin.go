package middleware

import (
	"net/http"
	"picture/common/response"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Admin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从上下文中获取JWT claims
		claims, ok := r.Context().Value(jwt.MapClaims{}).(jwt.MapClaims)
		if !ok {
			httpx.WriteJson(w, http.StatusUnauthorized, response.Error(401, "未登录"))
			return
		}

		// 获取用户角色
		userRole, ok := claims["role"].(string)
		if !ok || userRole != "admin" {
			httpx.WriteJson(w, http.StatusForbidden, response.Error(403, "无权限"))
			return
		}

		next(w, r)
	}
}
