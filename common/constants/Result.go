package constants

import "picture/common/errorx"

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"Message"`
	Data    any    `json:"data"`
}

func Success(data any) *Result {
	return &Result{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

func Fail(err *errorx.CodeError) *Result {
	return &Result{
		Code:    err.Code,
		Message: err.Message,
	}
}
