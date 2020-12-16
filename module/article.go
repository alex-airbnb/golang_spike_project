package module

import (
	"github.com/alex-airbnb/golang_spike_project/entity"

	"github.com/alex-airbnb/golang_spike_project/port"
)

// CreateArticle Creates a new article
func CreateArticle(article entity.Article, adapter port.ExternalStorage) entity.Articles {
	return adapter.CreateArticle(article)
}
