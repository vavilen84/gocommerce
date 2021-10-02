package logger

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"log"
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

func LogModelNotValidError(model string) {
	logs.Error("Model validation . Model %v is not valid")
}

func LogError(err error) {
	_, fn, line, _ := runtime.Caller(1)
	log.Printf("[error] %s:%d %v\n", fn, line, err)
	logs.Error("Error. Message: ", err.Error())
}

func LogCmdOut(i string) {
	logs.Info("CMD Out. Message: ", i)
}

func LogFormattedError(err error) {
	_, fn, line, _ := runtime.Caller(1)
	log.Printf("[error] %s:%d %v\n", fn, line, err)
}

func LogOrmerError(model string, err error) {
	logs.Error(nil, "Ormer error. Model %v. Message: ", err.Error())
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
