package logger

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"os"
	"runtime"
)

var l *logs.BeeLogger

func InitLogger() {
	l = logs.NewLogger(10000)
	err := l.SetLogger("file", `{"filename":"app.log"}`)
	if err != nil {
		fmt.Println(err)
	}
}

func LogModelFieldNotValidError(model string, v ...interface{}) {
	logs.Error("Validation error. Model %v. Field: ", model, v)
}

func LogError(err error) {
	_, fn, line, _ := runtime.Caller(1)
	logs.Error("Error. %v; Line: %d; Message: %v", fn, line, err.Error())
}

func LogCmdOut(i string) {
	logs.Info("CMD Out. Message: ", i)
}

func LogOrmerError(model string, err error) {
	_, fn, line, _ := runtime.Caller(1)
	logs.Error("Ormer error. Model %v; %v; Line: %d; Message: ", model, fn, line, err.Error())
}

func LogFatal(i interface{}) {
	var err error

	switch i.(type) {
	case error:
	case fmt.Stringer:
		err = errors.New(i.(fmt.Stringer).String())
	case string:
		err = errors.New(i.(string))
	default:
		msg := fmt.Sprintf("Log Fatal: %v")
		err = errors.New(msg)
	}

	LogError(err)
	os.Exit(1)
}
