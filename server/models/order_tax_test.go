package models

//
//import (
//	"github.com/stretchr/testify/assert"
//	"github.com/vavilen84/gocommerce/constants"
//	"github.com/vavilen84/gocommerce/validator"
//	"testing"
//)
//
//func TestOrderTax_ValidateOnCreate(t *testing.T) {
//	m := OrderTax{}
//	err := validator.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.NotNil(t, err)
//	assert.NotEmpty(t, err[constants.OrderTaxOrderIdField])
//	assert.NotEmpty(t, err[constants.OrderTaxTaxIdField])
//
//	m = OrderTax{
//		OrderId: 1,
//		TaxId:   1,
//	}
//	err = validator.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.NotNil(t, err)
//}
