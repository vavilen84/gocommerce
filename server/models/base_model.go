package models

type BaseModel struct {
	Scenario         Scenario
	ValidationErrors []error
	CreatedAt        int64 `json:"created_at" column:"created_at"`
	UpdatedAt        int64 `json:"updated_at" column:"updated_at"`
}
