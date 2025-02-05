package errorx

import (
	"net/http"
)

// 统一错误处理器
func ErrorHandler(err error) (int, interface{}) {
	switch e := err.(type) {
	case *CodeError:
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
