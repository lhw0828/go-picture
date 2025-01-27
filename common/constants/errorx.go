package constants

type CodeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *CodeError) Error() string {
	return e.Message
}

func NewCodeError(code int, message string) error {
	return &CodeError{
		Code:    code,
		Message: message,
	}
}

func NewSystemError(message string) error {
	return NewCodeError(SystemErr, message)
}

func NewParamError(message string) error {
	return NewCodeError(ParamError, message)
}
