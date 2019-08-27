package zlog

import (
	"fmt"
	"log"
	"os"
)

var (
	_log *log.Logger
)

//Zlog 日志
type Zlog struct {
}

func init() {
	_log = log.New(os.Stdout, "[INFO]", log.LstdFlags|log.Lshortfile)
}

//Info 信息
func Info(v ...interface{}) {
	_log.SetPrefix("[INFO]")
	_ = _log.Output(2, fmt.Sprint(v...))
}

//Infof 信息
func Infof(format string, v ...interface{}) {
	_log.SetPrefix("[INFO]")
	_ = _log.Output(2, fmt.Sprintf(format, v...))
}

//Infoln 信息
func Infoln(v ...interface{}) {
	_log.SetPrefix("[INFO]")
	_ = _log.Output(2, fmt.Sprintln(v...))
}

//Warn 提示
func Warn(v ...interface{}) {
	_log.SetPrefix("[WARN]")
	_ = _log.Output(2, fmt.Sprint(v...))
}

//Warnf 提示
func Warnf(format string, v ...interface{}) {
	_log.SetPrefix("[WARN]")
	_ = _log.Output(2, fmt.Sprintf(format, v...))
}

//Warnln 提示
func Warnln(v ...interface{}) {
	_log.SetPrefix("[WARN]")
	_ = _log.Output(2, fmt.Sprintln(v...))
}

//Error 错误
func Error(v ...interface{}) {
	_log.SetPrefix("[ERRO]")
	_ = _log.Output(2, fmt.Sprint(v...))
}

//Errorf 错误
func Errorf(format string, v ...interface{}) {
	_log.SetPrefix("[ERRO]")
	_ = _log.Output(2, fmt.Sprintf(format, v...))
}

//Errorln 错误
func Errorln(v ...interface{}) {
	_log.SetPrefix("[ERRO]")
	_ = _log.Output(2, fmt.Sprintln(v...))
}

//Debug 调试
func Debug(v ...interface{}) {
	_log.SetPrefix("[DEBG]")
	_ = _log.Output(2, fmt.Sprint(v...))
}

//Debugf 调试
func Debugf(format string, v ...interface{}) {
	_log.SetPrefix("[DEBG]")
	_ = _log.Output(2, fmt.Sprintf(format, v...))
}

//Debugln 调试
func Debugln(v ...interface{}) {
	_log.SetPrefix("[DEBG]")
	_ = _log.Output(2, fmt.Sprintln(v...))
}

//Fatal 致命信息
func Fatal(v ...interface{}) {
	_log.SetPrefix("[FTAL]")
	_ = _log.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

//Fatalf 致命信息
func Fatalf(format string, v ...interface{}) {
	_log.SetPrefix("[FTAL]")
	_ = _log.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

//Fataln 致命信息
func Fataln(v ...interface{}) {
	_log.SetPrefix("[FTAL]")
	_ = _log.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

//Painc Painc
func Painc(v ...interface{}) {
	_log.SetPrefix("[PANC]")
	s := fmt.Sprint(v...)
	_ = _log.Output(2, s)
	panic(s)
}

//Paincf Painc
func Paincf(format string, v ...interface{}) {
	_log.SetPrefix("[PANC]")
	s := fmt.Sprintf(format, v...)
	_ = _log.Output(2, s)
	panic(s)
}

//Paincln Painc
func Paincln(v ...interface{}) {
	_log.SetPrefix("[PANC]")
	s := fmt.Sprintln(v...)
	_ = _log.Output(2, s)
	panic(s)
}
