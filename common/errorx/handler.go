package errorx

import (
	"errors"
	"net/http"
)

// 统一错误处理器
func ErrorHandler(err error) (int, interface{}) {
	var e *CodeError
	switch {
	case errors.As(err, &e):
		return http.StatusOK, map[string]interface{}{
			"code":    e.Code,
			"message": e.Message,
		}
	default:
		return http.StatusOK, map[string]interface{}{
			"code":    SystemErr,
			"message": err.Error(),
		}
	}
}
