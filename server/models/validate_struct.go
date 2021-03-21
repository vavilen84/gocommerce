package models

import (
	"fmt"
	"github.com/vavilen84/gocommerce/constants"
	"gopkg.in/go-playground/validator.v9"
	"reflect"
)

type StructError struct {
	Tag     string
	Field   string
	Message string
	Value   string
	Param   string
	Name    string
}

type StructErrors []StructError

func (s StructErrors) Error() (result string) {
	for _, e := range s {
		result += e.Message + ";\n"
	}
	return
}

func (s *StructError) setErrorMessage() {
	switch s.Tag {
	case constants.RequiredTag:
		s.Message = fmt.Sprintf(constants.RequiredErrorMsg, s.Name, s.Field)
	case constants.MinTag:
		s.Message = fmt.Sprintf(constants.MinValueErrorMsg, s.Name, s.Field, s.Param)
	case constants.MaxTag:
		s.Message = fmt.Sprintf(constants.MaxValueErrorMsg, s.Name, s.Field, s.Param)
	case constants.EmailTag:
		s.Message = fmt.Sprintf(constants.EmailErrorMsg, s.Name)
	}
}

func getType(s interface{}) string {
	if t := reflect.TypeOf(s); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func ValidateStruct(s interface{}) error {
	v := validator.New()
	err := v.Struct(s)
	if err != nil {
		result := make(StructErrors, 0)
		var structError StructError
		for _, e := range err.(validator.ValidationErrors) {
			structError = StructError{
				Name:  getType(s),
				Tag:   e.Tag(),
				Field: e.Field(),
				Value: fmt.Sprintf("%v", e.Value()),
				Param: e.Param(),
			}
			structError.setErrorMessage()
			result = append(result, structError)
		}
		return result
	}
	return nil
}

func validate(s interface{}) []error {
	errs := make([]error, 0)
	err := ValidateStruct(s)
	if err != nil {
		errs = append(errs, err)
	}
	return errs
}
