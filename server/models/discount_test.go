package models

//
//import (
//	"github.com/stretchr/testify/assert"
//	"github.com/vavilen84/gocommerce/constants"
//	"github.com/vavilen84/gocommerce/validator"
//	"testing"
//)
//
//func TestDiscount_ValidateOnCreate(t *testing.T) {
//	m := Discount{}
//	err := validator.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.NotNil(t, err)
//	assert.NotEmpty(t, err[constants.DiscountAmountField])
//	assert.NotEmpty(t, err[constants.DiscountPercentageField])
//	assert.NotEmpty(t, err[constants.DiscountTitleField])
//	assert.NotEmpty(t, err[constants.DiscountTypeField])
//
//	m = Discount{
//		Title:  "product",
//		Amount: 1,
//		Type:   constants.DiscountCartType,
//	}
//	err = validator.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.NotNil(t, err)
//
//	m = Discount{
//		Title:      "product",
//		Percentage: 1,
//		Type:       constants.DiscountCategoryType,
//	}
//	err = validator.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.NotNil(t, err)
//}
