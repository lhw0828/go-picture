package middleware

import (
	"context"
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
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authorization, "Bearer ") {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authorization, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.secretKey), nil
		})

		if err != nil {
			logx.Errorf("Parse token error: %v", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		next(w, r.WithContext(ctx))
	}
}
