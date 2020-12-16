package port

import "github.com/alex-airbnb/golang_spike_project/entity"

// ExternalStorage Port to handle the Databases.
type ExternalStorage interface {
	CreateArticle(article entity.Article) entity.Articles
}
