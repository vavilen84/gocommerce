package validation

import (
	"reflect"
)

type Scenario string

type Rules string

type Field string

type ScenarioRules map[Scenario]FieldRules

type FieldRules map[Field]Rules

func getType(s interface{}) string {
	if t := reflect.TypeOf(s); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

// need pass ptr or interface instead of struct - otherwise func panics
func structToMap(input interface{}) map[Field]interface{} {
	r := make(map[Field]interface{})
	s := reflect.ValueOf(input).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		r[Field(typeOfT.Field(i).Name)] = f.Interface()
	}
	return r
}
