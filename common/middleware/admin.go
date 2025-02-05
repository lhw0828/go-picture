package middleware

import (
	"net/http"
	"picture/common/errorx"
	"picture/common/response"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Admin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 从请求头获取 token
		authorization := r.Header.Get("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			httpx.WriteJson(w, http.StatusOK, response.Error(errorx.UnauthorizedErr, errorx.UnauthorizedErrMsg))
			return
		}

		token := strings.TrimPrefix(authorization, "Bearer ")
		// 解析 token
		parser := jwt.NewParser()
		claims := make(jwt.MapClaims)
		_, err := parser.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("ning4256"), nil // 使用实际的 secret
		})

		if err != nil {
			logx.Errorf("Parse token error: %v", err)
			httpx.WriteJson(w, http.StatusOK, response.Error(errorx.UnauthorizedErr, errorx.UnauthorizedErrMsg))
			return
		}

		// 从 claims 中获取用户角色
		userRole, ok := claims["userRole"].(string)
		logx.Infof("UserRole from claims: %s, ok: %v", userRole, ok)
		if !ok || userRole != "admin" {
			logx.Errorf("Invalid user role: %s", userRole)
			httpx.WriteJson(w, http.StatusOK, response.Error(errorx.ForbiddenErr, errorx.ForbiddenErrMsg))
			return
		}

		next(w, r)
	}
}
