package model

import "gorm.io/gorm"

// Article DB model
type Article struct {
	gorm.Model
	Name    string
	Content string
}
