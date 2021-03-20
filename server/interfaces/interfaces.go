package interfaces

import "github.com/vavilen84/gocommerce/models"

type Model interface {
	GetTableName() string
	GetId() uint32
	ValidateByScenario(scenario models.Scenario) []error
}
