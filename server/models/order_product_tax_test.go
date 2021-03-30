package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/validation"
	"testing"
)

func TestOrderProductTax_ValidateOnCreate(t *testing.T) {
	m := OrderProductTax{}
	err := validation.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
	assert.NotNil(t, err)
	assert.NotEmpty(t, err[constants.OrderProductTaxOrderProductIdField])
	assert.NotEmpty(t, err[constants.OrderProductTaxTaxIdField])

	m = OrderProductTax{
		OrderProductId: 1,
		TaxId:          1,
	}
	err = validation.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
	assert.NotNil(t, err)
}
