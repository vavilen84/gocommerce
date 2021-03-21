package interfaces

type Model interface {
	GetTableName() string
	GetId() uint32
}
