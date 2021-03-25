package helpers

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"time"
)

var (
	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func LogFatal(msg string) {
	err := errors.New(msg)
	LogError(err)
	os.Exit(1)
}

func DebugInterface(i interface{}) (v reflect.Value, t reflect.Type, ts string) {
	v = reflect.ValueOf(i)
	t = reflect.TypeOf(i)
	ts = t.String()
	return
}

func DebugError(err error) (msg string) {
	msg = err.Error()
	return
}

func LogError(err error) {
	_, fn, line, _ := runtime.Caller(1)
	log.Printf("[error] %s:%d %v\n", fn, line, err)
}

func Dump(i interface{}) {
	fmt.Printf("%+v\n", i)
}

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
