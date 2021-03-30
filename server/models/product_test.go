package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/gocommerce/constants"
	"github.com/vavilen84/gocommerce/validation"
	"testing"
)

func TestProduct_ValidateOnCreate(t *testing.T) {
	m := Product{}
	err := validation.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
	assert.NotNil(t, err)
	assert.NotEmpty(t, err[constants.ProductPriceField])
	assert.NotEmpty(t, err[constants.ProductSKUField])
	assert.NotEmpty(t, err[constants.ProductTitleField])

	m = Product{
		Title: "product",
		SKU:   "sku-123_123",
		Price: 100,
	}
	err = validation.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
	assert.NotNil(t, err)
}
