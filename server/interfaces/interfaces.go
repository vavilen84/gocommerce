package interfaces

type Model interface {
	GetTableName() string
	GetId() uint32
	SetId(id uint32) Model
	GetValidator() interface{}
	GetValidationRules() interface{}
}
