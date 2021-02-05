package router

import (
	"encoding/json"
	"net/http"

	"github.com/alex-airbnb/golang_spike_project/model"
)

// CreateArticleHandler Handle the a POST request and create new article
func CreateArticleHandler(writer http.ResponseWriter, request *http.Request) {
	var article model.Article

	// adapter := adapter.Postgres{DB: database.PostgresDB}
	err := json.NewDecoder(request.Body).Decode(&article)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	// createdArticle, err := module.CreateArticle(&article, &adapter)

	// if err != nil {
	// 	http.Error(writer, err.Error(), http.StatusBadRequest)
	// }

	// writer.Header().Set("Content-Type", "application/json")

	// writer.WriteHeader(http.StatusCreated)

	// json.NewEncoder(writer).Encode(createdArticle)
}
