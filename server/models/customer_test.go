package models

//
//import (
//	"github.com/stretchr/testify/assert"
//	"github.com/vavilen84/gocommerce/constants"
//	"github.com/vavilen84/gocommerce/validator"
//	"testing"
//)
//
//func TestCustomer_ValidateOnCreate(t *testing.T) {
//	m := Customer{}
//	err := validator.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.NotNil(t, err)
//	assert.NotEmpty(t, err[constants.CustomerEmailField])
//	assert.NotEmpty(t, err[constants.CustomerFirstNameField])
//	assert.NotEmpty(t, err[constants.CustomerLastNameField])
//
//	m = Customer{
//		FirstName: "John",
//		LastName:  "Dou",
//		Email:     "user@example.com",
//	}
//	err = validator.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.NotNil(t, err)
//}
