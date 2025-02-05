package middleware

import (
	"net/http"
	"picture/common/response"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Recovery() func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logx.WithContext(r.Context()).Errorf("panic recovered: %v", err)
					httpx.WriteJson(w, http.StatusInternalServerError, response.Error(500, "系统内部错误"))
				}
			}()
			next(w, r)
		}
	}
}
