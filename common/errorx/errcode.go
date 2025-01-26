package errorx

const (
	Success       = 0
	ParamError    = 40000
	AuthError     = 40100
	NotFoundError = 40400
	SystemError   = 50000
)

var message = map[int]string{
	Success:       "success",
	ParamError:    "参数错误",
	AuthError:     "未登录或无权限",
	NotFoundError: "资源不存在",
	SystemError:   "系统内部异常",
}

func MapErrMsg(errcode int) string {
	if msg, ok := message[errcode]; ok {
		return msg
	}
	return "服务器开小差啦,请稍后再试"
}
