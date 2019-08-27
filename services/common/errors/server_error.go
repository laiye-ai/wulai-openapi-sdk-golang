package errors

import (
	"fmt"
)

//ServerError 服务端错误定义
type ServerError struct {
	error
	httpStatus  int
	errorCode   string
	message     string
	originError error
}

//NewServerError 实例化 ServerError
func NewServerError(httpStatus int, errorCode, message string, originErr error) Error {
	return &ServerError{
		httpStatus:  httpStatus,
		errorCode:   errorCode,
		message:     message,
		originError: originErr,
	}
}

func (err *ServerError) Error() string {
	clientErrMsg := fmt.Sprintf("[%s] %s", err.ErrorCode(), err.message)
	if err.originError != nil {
		return clientErrMsg + "\ncaused by:\n" + err.originError.Error()
	}
	return clientErrMsg
}

//OriginError HTTP 原始错误信息
func (err *ServerError) OriginError() error {
	return err.originError
}

//HttpStatus HTTP 状态码
func (err *ServerError) HttpStatus() int {
	return err.httpStatus
}

//ErrorCode HTTP 错误码
func (err *ServerError) ErrorCode() string {
	return err.errorCode
}

//Message HTTP 消息
func (err *ServerError) Message() string {
	return err.message
}

//String HTTP 错误消息
func (err *ServerError) String() string {
	return err.Error()
}
