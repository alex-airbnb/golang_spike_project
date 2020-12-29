package adapter

import (
	"gorm.io/gorm"
)

// Postgres Adapter to use Postgres that implements the ExternalStorage port.
type Postgres struct {
	DB *gorm.DB
}

// Create Create a new item into the table of the model.
func (adapter *Postgres) Create(item interface{}) error {
	err := adapter.DB.Create(item).Error

	return err
}
