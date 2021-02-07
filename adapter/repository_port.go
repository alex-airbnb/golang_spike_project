package adapter

// Repository Port for databases.
type RepositoryPort interface {
	Create(model interface{}) error
}
