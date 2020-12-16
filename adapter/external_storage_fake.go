package adapter

import "github.com/alex-airbnb/golang_spike_project/entity"

// ExternalStorageFake Adapter for test purposes that implements the ExternalStorage port
type ExternalStorageFake struct{}

// CreateArticle Creates a new article
func (adapter ExternalStorageFake) CreateArticle(article entity.Article) entity.Articles {
	return []entity.Article{article}
}
