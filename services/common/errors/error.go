package errors

//Error 定义错误
type Error interface {
	error
	HttpStatus() int
	ErrorCode() string
	Message() string
	String() string
	OriginError() error
}
