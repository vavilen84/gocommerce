package models

//
//import (
//	"github.com/stretchr/testify/assert"
//	"github.com/vavilen84/gocommerce/constants"
//	"github.com/vavilen84/gocommerce/validator"
//	"testing"
//	"time"
//)
//
//func TestMigration_ValidateOnCreate(t *testing.T) {
//	m := Migration{}
//	err := validator.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.NotEmpty(t, err)
//	assert.NotEmpty(t, err[constants.MigrationCreatedAtField])
//	assert.NotEmpty(t, err[constants.MigrationVersionField])
//	assert.NotEmpty(t, err[constants.MigrationFilenameField])
//
//	m = Migration{
//		Version:   time.Now().Unix(),
//		Filename:  "inital_migration",
//		CreatedAt: time.Now().Unix(),
//	}
//	err = validator.ValidateByScenario(constants.ScenarioCreate, &m, m.getValidator(), m.getValidationRules())
//	assert.Empty(t, err)
//}
