package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
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
			logx.WithContext(r.Context()).Error("Missing authorization header")
			http.Error(w, "请提供认证信息", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authorization, "Bearer ") {
			logx.WithContext(r.Context()).Error("Invalid authorization format")
			http.Error(w, "认证格式错误", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authorization, "Bearer ")
		
		// 添加 token 解析前的日志
		logx.WithContext(r.Context()).Infof("Parsing token: %s", tokenString)
		logx.WithContext(r.Context()).Infof("Using secret key: %s", m.secretKey)

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				logx.WithContext(r.Context()).Errorf("Unexpected signing method: %v", token.Header["alg"])
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			
			// 确保密钥不为空
			if m.secretKey == "" {
				logx.WithContext(r.Context()).Error("Secret key is empty")
				return nil, fmt.Errorf("secret key is empty")
			}
			
			return []byte(m.secretKey), nil
		})

		if err != nil {
			logx.WithContext(r.Context()).Errorf("Token validation details - Error: %v, Token: %s, Claims: %+v", 
				err, tokenString, claims)
			
			if ve, ok := err.(*jwt.ValidationError); ok {
				logx.WithContext(r.Context()).Errorf("Validation error bits: %b", ve.Errors)
				switch {
				case ve.Errors&jwt.ValidationErrorExpired != 0:
					http.Error(w, "token已过期", http.StatusUnauthorized)
				case ve.Errors&jwt.ValidationErrorSignatureInvalid != 0:
					http.Error(w, "token签名无效", http.StatusUnauthorized)
				default:
					http.Error(w, fmt.Sprintf("token验证失败: %v", err), http.StatusUnauthorized)
				}
			} else {
				http.Error(w, fmt.Sprintf("token验证失败: %v", err), http.StatusUnauthorized)
			}
			return
		}

		if !token.Valid {
			logx.WithContext(r.Context()).Error("Invalid token")
			http.Error(w, "无效的token", http.StatusUnauthorized)
			return
		}

		// 添加用户信息到上下文
		userId, ok := claims["userId"].(float64)
		if !ok {
			logx.WithContext(r.Context()).Error("Missing or invalid userId in token claims")
			http.Error(w, "token信息不完整", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		ctx = context.WithValue(ctx, "userId", int64(userId))
		next(w, r.WithContext(ctx))
	}
}
