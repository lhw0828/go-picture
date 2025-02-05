package middleware

import (
	"context"
	"fmt"
	"net/http"

	"picture/common/errorx"
	"picture/common/response"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthMiddleware struct {
	secretKey string
}

func NewAuthMiddleware(secretKey string) *AuthMiddleware {
	return &AuthMiddleware{
		secretKey: secretKey,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			httpx.OkJson(w, response.Error(errorx.UnauthorizedErr, "请提供认证信息"))
			return
		}

		if !strings.HasPrefix(authorization, "Bearer ") {
			httpx.OkJson(w, response.Error(errorx.UnauthorizedErr, "认证格式错误"))
			return
		}

		tokenString := strings.TrimPrefix(authorization, "Bearer ")

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(m.secretKey), nil
		})

		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				switch {
				case ve.Errors&jwt.ValidationErrorExpired != 0:
					httpx.OkJson(w, response.Error(errorx.UnauthorizedErr, "token已过期"))
				case ve.Errors&jwt.ValidationErrorSignatureInvalid != 0:
					httpx.OkJson(w, response.Error(errorx.UnauthorizedErr, "token签名无效"))
				default:
					httpx.OkJson(w, response.Error(errorx.UnauthorizedErr, "token验证失败"))
				}
			} else {
				httpx.OkJson(w, response.Error(errorx.UnauthorizedErr, "token验证失败"))
			}
			return
		}

		if !token.Valid {
			httpx.WriteJson(w, http.StatusOK, response.Error(errorx.UnauthorizedErr, "无效的token"))
			return
		}

		logx.Infof("Token claims: %+v", claims)

		// 将 claims 存入上下文
		ctx := context.WithValue(r.Context(), ClaimsKey, claims)

		// 获取并存储用户ID
		userId, ok := claims["userId"].(float64)
		if !ok {
			logx.Errorf("Failed to get userId from claims")
			httpx.WriteJson(w, http.StatusOK, response.Error(errorx.UnauthorizedErr, "token信息不完整"))
			return
		}
		logx.Infof("UserId from claims: %v", userId)

		logx.Infof("Raw token claims: %+v", claims)

		// 获取并存储用户角色
		userRole, ok := claims["userRole"].(string)
		if !ok {
			logx.Errorf("Failed to get userRole from claims, claims: %+v", claims)
			httpx.WriteJson(w, http.StatusOK, response.Error(errorx.UnauthorizedErr, "token信息不完整"))
			return
		}

		ctx = context.WithValue(r.Context(), UserRoleKey, userRole)
		logx.Infof("Set userRole in context: %s", userRole)

		next(w, r.WithContext(ctx))
	}
}
