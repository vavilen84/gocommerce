package validation

import (
	"fmt"
	"github.com/vavilen84/gocommerce/helpers"
	"github.com/vavilen84/gocommerce/interfaces"
	"gopkg.in/go-playground/validator.v9"
)

// should be passed ptr to modelm otherwise - func will panic
func ValidateByScenario(scenario Scenario, m interfaces.Model, validate *validator.Validate, validationMap ScenarioRules) Errors {
	errs := make(Errors)
	data := structToMap(m)
	for fieldName, validation := range validationMap[scenario] {
		field, ok := data[fieldName]
		if !ok {
			helpers.LogFatal(fmt.Sprintf("Field not found: %s", fieldName))
		}
		err := validate.Var(field, string(validation))
		if err != nil {
			if _, ok := errs[fieldName]; !ok {
				errs[fieldName] = make([]FieldError, 0)
			}
			for _, e := range err.(validator.ValidationErrors) {
				validationError := FieldError{
					Name:  getType(m),
					Tag:   e.Tag(),
					Field: fieldName,
					Value: fmt.Sprintf("%v", e.Value()),
					Param: e.Param(),
				}
				validationError.setErrorMessage()
				errs[fieldName] = append(errs[fieldName], validationError)
			}
		}
	}
	return errs
}
