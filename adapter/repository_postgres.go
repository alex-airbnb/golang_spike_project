package adapter

import (
	"gorm.io/gorm"
)

// RepositoryPostgres Adapter Implements the Repository Port for PostgreSQL.
type RepositoryPostgres struct {
	DBClient *gorm.DB
}

// Create Create a new item into the table of the model.
func (r *RepositoryPostgres) Create(m interface{}) error {
	return r.DBClient.Create(m).Error
}
