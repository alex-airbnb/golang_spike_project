package entity

import (
	"time"

	"gorm.io/gorm"
)

// Article DB Entity to manage the Article data and also work as model for Gorm.
type Article struct {
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"createdAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	ID        uint           `json:"id" gorm:"primarykey"`
	Name      string         `json:"name"`
	UpdatedAt time.Time      `json:"updatedAt"`
}

// CreateArticle Creates a new article.
func CreateArticle(name, content string) (Article, error) {
	return Article{Name: name, Content: content}, nil
}
