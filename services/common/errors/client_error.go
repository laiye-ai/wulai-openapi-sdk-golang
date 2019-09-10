package errors

import "fmt"

const (
	DefaultClientErrorStatus = 400
	DefaultClientErrorCode   = "SDK.ClientError"

	NetWorkErrorCode    = "SDK.NetWorkError"
	NetWorkErrorMessage = "Network error %s,try using err.Message() to get detail message"

	UnsupportedMethodErrorCode    = "SDK.UnsupportedMethod"
	UnsupportedMethodErrorMessage = "The version (%s) is not supported,retry using (%s)"

	UnsupportedTypeErrorCode    = "SDK.UnsupportedType"
	UnsupportedTypeErrorMessage = "Specified type (%s) is not supported,retry using (%s)"

	UnknownRequestTypeErrorCode    = "SDK.UnknownRequestType"
	UnknownRequestTypeErrorMessage = "Unknown Request Type: %s"

	MissingParamErrorCode = "SDK.MissingParam"
	InvalidParamErrorCode = "SDK.InvalidParam"

	InvalidFormatErrorCode    = "SDK.InvalidFormat"
	InvalidFormatErrorMessage = "Format invalid %s, try using err.Message() to get detail message"

	JsonMarshalErrorCode    = "SDK.JsonMarshalError"
	JsonMarshalErrorMessage = "Failed to marshal response, try using err.Message() to get detail message"

	JsonUnmarshalErrorCode    = "SDK.JsonUnmarshalError"
	JsonUnmarshalErrorMessage = "Failed to unmarshal response,try using err.Message() to get detail message"

	TimeoutErrorCode    = "SDK.TimeoutError"
	TimeoutErrorMessage = "The request timed out %s times(%s for retry), perhaps we should have the threshold raised a little?"
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
