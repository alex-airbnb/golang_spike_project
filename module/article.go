package module

import (
	"github.com/alex-airbnb/golang_spike_project/model"
	"github.com/alex-airbnb/golang_spike_project/port"
)

// CreateArticle Creates a new Article.
func CreateArticle(article *model.Article, adapter port.ExternalStorage) (*model.Article, error) {
	err := adapter.Create(article)

	return article, err
}
