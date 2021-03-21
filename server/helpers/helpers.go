package helpers

import (
	"errors"
	"fmt"
	"github.com/vavilen84/gocommerce/types"
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

func MergeErrors(i types.ValidationErrors) error {
	errMsg := ""
	for _, v := range i {
		errMsg += v.Error() + "; "
	}
	return errors.New(errMsg)
}

func LogFatal(msg string) {
	err := errors.New(msg)
	LogError(err)
	os.Exit(1)
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

// need pass ptr or interface instead of struct - otherwise func panics
func StructToMap(input interface{}) map[types.Field]interface{} {
	r := make(map[types.Field]interface{})
	s := reflect.ValueOf(input).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		r[types.Field(typeOfT.Field(i).Name)] = f.Interface()
	}
	return r
}
