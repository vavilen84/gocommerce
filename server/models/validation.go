package models

import (
	"fmt"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/helpers"
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

type Scenario string

type ValidationMap map[Scenario]ValidationRules

type Rules string

type Field string

type ValidationRules map[Field]Rules

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

func validateByScenario(m interface{}, validationMap ValidationMap) []error {
	scenarioField, ok := reflect.TypeOf(m).Elem().FieldByName("Scenario")
	if !ok {
		helpers.LogFatal(fmt.Sprintf("No Scenario field in struct: %+v", m))
	}
	scenario := reflect.ValueOf(scenarioField)
	if _, ok := validationMap[Scenario(scenario.String())]; !ok {
		helpers.LogFatal(fmt.Sprintf("No such scenario: %s", scenario))
	}
	errs := make([]error, 0)
	validate := validator.New()
	for fieldName, validation := range validationMap[Scenario(scenario.String())] {
		field, ok := reflect.TypeOf(m).Elem().FieldByName(string(fieldName))
		if !ok {
			helpers.LogFatal(fmt.Sprintf("Field not found: %s", fieldName))
		}
		result := validate.Var(field, string(validation))
		if result != nil {
			errs = append(errs, result)
		}
	}
	return errs
}
