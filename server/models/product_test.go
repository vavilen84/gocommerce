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
}
