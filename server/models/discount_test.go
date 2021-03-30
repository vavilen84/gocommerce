package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/validation"
	"testing"
)

func TestDiscount_ValidateOnCreate(t *testing.T) {
	m := Discount{}
	err := validation.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
	assert.NotNil(t, err)
	assert.NotEmpty(t, err[constants.DiscountAmountField])
	assert.NotEmpty(t, err[constants.DiscountPercentageField])
	assert.NotEmpty(t, err[constants.DiscountTitleField])

	m = Discount{
		Title:  "product",
		Amount: 1,
	}
	err = validation.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
	assert.NotNil(t, err)

	m = Discount{
		Title:      "product",
		Percentage: 1,
	}
	err = validation.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
	assert.NotNil(t, err)
}
