package contract

type ContractService[T any] interface {
	Create(entity *T) (*T, error)
	FindById(id string) (*T, error)
	FindAll() ([]*T, error)
	Update(entity *T) (*T, error)
	DeleteById(id string) error
}