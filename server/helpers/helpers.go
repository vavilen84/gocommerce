package helpers

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

var (
	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func Dump(i interface{}) {
	fmt.Printf("%+v\r\n", i)
	fmt.Printf("%T\r\n", i)
}

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// need pass ptr or interface instead of struct - otherwise func panics
func StructToMap(input interface{}) map[string]interface{} {
	r := make(map[string]interface{})
	s := reflect.ValueOf(input).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		r[typeOfT.Field(i).Name] = f.Interface()
	}
	return r
}
