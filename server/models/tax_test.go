package models

//
//import (
//	"github.com/stretchr/testify/assert"
//	"github.com/vavilen84/gocommerce/constants"
//	"github.com/vavilen84/gocommerce/validator"
//	"testing"
//)
//
//func TestTax_ValidateOnCreate(t *testing.T) {
//	m := Tax{}
//	err := validator.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.NotNil(t, err)
//	assert.NotEmpty(t, err[constants.TaxTitleField])
//	assert.NotEmpty(t, err[constants.TaxAmountField])
//	assert.NotEmpty(t, err[constants.TaxPercentageField])
//	assert.NotEmpty(t, err[constants.TaxTypeField])
//
//	m = Tax{
//		Title:  "product",
//		Amount: 1,
//		Type:   constants.TaxCartType,
//	}
//	err = validator.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.NotNil(t, err)
//
//	m = Tax{
//		Title:      "product",
//		Percentage: 1,
//		Type:       constants.TaxCategoryType,
//	}
//	err = validator.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.NotNil(t, err)
//}
