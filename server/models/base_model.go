package models

type BaseModel struct {
	scenario  Scenario
	CreatedAt int64 `json:"created_at" column:"created_at"`
	UpdatedAt int64 `json:"updated_at" column:"updated_at"`
}

func (m *BaseModel) SetScenario(scenario string) {
	m.scenario = Scenario(scenario)
}

func (m *BaseModel) GetScenario(scenario string) {
	m.scenario = Scenario(scenario)
}
