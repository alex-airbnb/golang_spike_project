package entity

import (
	"time"

	"gorm.io/gorm"
)

// Article DB Entity to manage the Article data and also work as model for Gorm.
type Article struct {
	Content   string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ID        uint           `gorm:"primarykey"`
	Name      string
	UpdatedAt time.Time
}

// CreateArticle Creates a new article.
func CreateArticle(content, name string) (Article, error) {
	return Article{Content: content, Name: name}, nil
}
