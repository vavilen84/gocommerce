package helpers

import (
	"errors"
	"fmt"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/logger"
	"math/rand"
	"os"
	"os/exec"
	"reflect"
	"strings"
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

func RunCmd(name string, arg ...string) {
	err := os.Chdir(os.Getenv(constants.AppRootEnvVar))
	if err != nil {
		logger.LogError(err)
	}
	cmd := exec.Command(
		name,
		arg...,
	)
	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			eStr := suppressErr(string(ee.Stderr))
			if eStr != "" {
				e := errors.New(string(ee.Stderr))
				logger.LogError(e)
			}
		}
	}
	oStr := suppressErr(string(out))
	if oStr != "" {
		logger.LogCmdOut(oStr)
	}
}

func suppressErr(i string) string {
	if strings.Contains(i, "Using a password on the command line interface can be insecure") {
		return ""
	}
	if strings.Contains(i, "Usage: mysql [OPTIONS] [database]") {
		return ""
	}
	return i
}
