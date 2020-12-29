package port

// ExternalStorage Port to handle the Databases.
type ExternalStorage interface {
	Create(model interface{}) error
}
