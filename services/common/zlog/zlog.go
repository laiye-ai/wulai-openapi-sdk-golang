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

func Info(v ...interface{}) {
	_log.SetPrefix("[INFO]")
	_ = _log.Output(2, fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
	_log.SetPrefix("[INFO]")
	_ = _log.Output(2, fmt.Sprintf(format, v...))
}

func Infoln(v ...interface{}) {
	_log.SetPrefix("[INFO]")
	_ = _log.Output(2, fmt.Sprintln(v...))
}

func Warn(v ...interface{}) {
	_log.SetPrefix("[WARN]")
	_ = _log.Output(2, fmt.Sprint(v...))
}

func Warnf(format string, v ...interface{}) {
	_log.SetPrefix("[WARN]")
	_ = _log.Output(2, fmt.Sprintf(format, v...))
}

func Warnln(v ...interface{}) {
	_log.SetPrefix("[WARN]")
	_ = _log.Output(2, fmt.Sprintln(v...))
}

func Error(v ...interface{}) {
	_log.SetPrefix("[ERRO]")
	_ = _log.Output(2, fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	_log.SetPrefix("[ERRO]")
	_ = _log.Output(2, fmt.Sprintf(format, v...))
}

func Errorln(v ...interface{}) {
	_log.SetPrefix("[ERRO]")
	_ = _log.Output(2, fmt.Sprintln(v...))
}

func Debug(v ...interface{}) {
	_log.SetPrefix("[DEBG]")
	_ = _log.Output(2, fmt.Sprint(v...))
}

func Debugf(format string, v ...interface{}) {
	_log.SetPrefix("[DEBG]")
	_ = _log.Output(2, fmt.Sprintf(format, v...))
}

func Debugln(v ...interface{}) {
	_log.SetPrefix("[DEBG]")
	_ = _log.Output(2, fmt.Sprintln(v...))
}

func Fatal(v ...interface{}) {
	_log.SetPrefix("[FTAL]")
	_ = _log.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	_log.SetPrefix("[FTAL]")
	_ = _log.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Fatalln(v ...interface{}) {
	_log.SetPrefix("[FTAL]")
	_ = _log.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

func Painc(v ...interface{}) {
	_log.SetPrefix("[PANC]")
	s := fmt.Sprint(v...)
	_ = _log.Output(2, s)
	panic(s)
}

func Paincf(format string, v ...interface{}) {
	_log.SetPrefix("[PANC]")
	s := fmt.Sprintf(format, v...)
	_ = _log.Output(2, s)
	panic(s)
}

func Paincln(v ...interface{}) {
	_log.SetPrefix("[PANC]")
	s := fmt.Sprintln(v...)
	_ = _log.Output(2, s)
	panic(s)
}
