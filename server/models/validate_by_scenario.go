package models

import (
	"fmt"
	"github.com/vavilen84/gocommerce/helpers"
	"github.com/vavilen84/gocommerce/types"
	"gopkg.in/go-playground/validator.v9"
)

func validateByScenario(scenario types.Scenario, m interface{}, validationMap types.ValidationMap) types.ValidationErrors {
	errs := make(types.ValidationErrors)
	validate := validator.New()
	data := helpers.StructToMap(m)
	for fieldName, validation := range validationMap[scenario] {
		field, ok := data[fieldName]
		if !ok {
			helpers.LogFatal(fmt.Sprintf("Field not found: %s", fieldName))
		}
		err := validate.Var(field, string(validation))
		if err != nil {
			errs[fieldName] = err
		}
	}
	return errs
}
