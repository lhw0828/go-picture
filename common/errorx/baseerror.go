package errorx

import "fmt"

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("Code: %d, Msg: %s", e.Code, e.Msg)
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(ParamError, msg)
}
