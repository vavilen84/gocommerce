package interfaces

type Model interface {
	GetTableName() string
	GetId() uint32
	SetId(id uint32)
	GetValidator() interface{}
	GetValidationRules() interface{}
}
