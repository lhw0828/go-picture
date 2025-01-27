package errorx

import (
	"errors"
	"net/http"
	"strings"

	"picture/common/response"

	"github.com/zeromicro/go-zero/core/logx"
)

// 统一错误处理器
func ErrorHandler(err error) (int, interface{}) {
	var e *CodeError
	switch {
	case errors.As(err, &e):
		return http.StatusOK, response.Error(e.Code, e.Message)
	case strings.Contains(err.Error(), "type mismatch"):
		return http.StatusOK, response.Error(ParamError, "参数类型错误")
	case strings.Contains(err.Error(), "invalid parameter"):
		return http.StatusOK, response.Error(ParamError, "无效的参数")
	default:
		logx.Errorf("系统错误: %+v", err)
		return http.StatusOK, response.Error(SystemErr, "系统错误")
	}
}
