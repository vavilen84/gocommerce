package validation

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type Errors map[string][]FieldError

func (s Errors) Error() (result string) {
	for _, errs := range s {
		for _, e := range errs {
			result += e.Message + ";\n"
		}
	}
	return
}

func (s Errors) setStructValidationErrors(err error, model interface{}) {
	for _, e := range err.(validator.ValidationErrors) {
		field := e.Field()
		if _, ok := s[field]; !ok {
			s[field] = make([]FieldError, 0)
		}
		validationError := FieldError{
			Name:  getType(model),
			Tag:   e.Tag(),
			Field: field,
			Value: fmt.Sprintf("%v", e.Value()),
			Param: e.Param(),
		}
		validationError.setErrorMessage()
		s[field] = append(s[field], validationError)
	}
}
