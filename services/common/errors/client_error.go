package errors

import "fmt"

const (

	//DefaultClientErrorStatus 默认客户端错误状态码
	DefaultClientErrorStatus = 400

	//DefaultClientErrorCode 默认客户端错误信息
	DefaultClientErrorCode = "SDK.ClientError"
)

//ClientError 客户端错误定义
type ClientError struct {
	errorCode   string
	message     string
	originError error
}

//NewClientError 实例化 ClientError
func NewClientError(errorCode, message string, originErr error) Error {
	return &ClientError{
		errorCode:   errorCode,
		message:     message,
		originError: originErr,
	}
}

func (err *ClientError) Error() string {
	clientErrMsg := fmt.Sprintf("[%s] %s", err.ErrorCode(), err.message)
	if err.originError != nil {
		return clientErrMsg + "\ncaused by:\n" + err.originError.Error()
	}
	return clientErrMsg
}

//OriginError HTTP 原始错误信息
func (err *ClientError) OriginError() error {
	return err.originError
}

//HttpStatus HTTP 状态码
func (*ClientError) HttpStatus() int {
	return DefaultClientErrorStatus
}

//ErrorCode HTTP 错误码
func (err *ClientError) ErrorCode() string {
	if err.errorCode == "" {
		return DefaultClientErrorCode
	}

	return err.errorCode
}

//Message HTTP 消息
func (err *ClientError) Message() string {
	return err.message
}

//String HTTP 错误消息
func (err *ClientError) String() string {
	return err.Error()
}
