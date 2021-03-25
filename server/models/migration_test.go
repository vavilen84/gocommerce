package models

import (
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/gocommerce/constants"
	"testing"
	"time"
)

func TestMigration_ValidateByScenario(t *testing.T) {
	m := Migration{}
	m.Scenario = constants.ScenarioCreate
	m.ValidateByScenario()

	assert.NotEmpty(t, m.Errors[constants.MigrationUpdatedAtField])
	assert.NotEmpty(t, m.Errors[constants.MigrationCreatedAtField])
	assert.NotEmpty(t, m.Errors[constants.MigrationVersionField])
	assert.NotEmpty(t, m.Errors[constants.MigrationFilenameField])

	m = Migration{
		Version:   time.Now().Unix(),
		Filename:  "inital_migration",
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	m.Scenario = constants.ScenarioCreate
	m.ValidateByScenario()
	assert.Empty(t, m.Errors)
}
