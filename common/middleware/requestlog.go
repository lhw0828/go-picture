package middleware

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

func RequestLog(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next(w, r)

		logx.WithContext(r.Context()).Infow("HTTP Request",
			logx.Field("method", r.Method),
			logx.Field("path", r.URL.Path),
			logx.Field("duration", time.Since(start)),
			logx.Field("ip", r.RemoteAddr),
		)
	}
}
