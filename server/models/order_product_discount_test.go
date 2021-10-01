package models

//
//import (
//	"github.com/stretchr/testify/assert"
//	"github.com/vavilen84/gocommerce/constants"
//	"github.com/vavilen84/gocommerce/validator"
//	"testing"
//)
//
//func TestOrderProductDiscount_ValidateOnCreate(t *testing.T) {
//	m := OrderProductDiscount{}
//	err := validator.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.NotNil(t, err)
//	assert.NotEmpty(t, err[constants.OrderProductDiscountOrderProductIdField])
//	assert.NotEmpty(t, err[constants.OrderProductDiscountDiscountIdField])
//
//	m = OrderProductDiscount{
//		OrderProductId: 1,
//		DiscountId:     1,
//	}
//	err = validator.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.NotNil(t, err)
//}
