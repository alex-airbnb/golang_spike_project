package usecase

import "encoding/json"

// Article Use Case interface.
type Article interface {
	CreateArticle(interface{}) (interface{}, error)
}

// ArticleREST implement the Article interface for a REST use case.
type ArticleREST struct{}

// CreateArticleRESTRequest Request Model for CreateArticle ArticleREST's method.
type CreateArticleRESTRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// CreateArticleRESTResponse Reponse Model for CreateArticle ArticleREST's method.
type CreateArticleRESTResponse struct {
	Content string `json:"content"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
}

// CreateArticle Create a new article in the Repository using the REST use case.
func (a *ArticleREST) CreateArticle(req []byte) (CreateArticleRESTResponse, error) {
	var res CreateArticleRESTResponse

	if err := json.Unmarshal(req, &res); err != nil {
		return CreateArticleRESTResponse{}, err
	}

	return res, nil
}
