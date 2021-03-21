package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/vavilen84/gocommerce/constants"
	"testing"
)

func TestMigration_ValidateByScenario(t *testing.T) {
	m := Migration{}
	m.Scenario = constants.ScenarioCreate
	m.ValidateByScenario()
	assert.NotEmpty(t, m.ValidationErrors)
	fmt.Printf("%+v", m.ValidationErrors[constants.MigrationUpdatedAtField])
	assert.NotEmpty(t, m.ValidationErrors[constants.MigrationUpdatedAtField])
	assert.NotEmpty(t, m.ValidationErrors[constants.MigrationCreatedAtField])
	assert.NotEmpty(t, m.ValidationErrors[constants.MigrationVersionField])
	assert.NotEmpty(t, m.ValidationErrors[constants.MigrationFilenameField])
}
