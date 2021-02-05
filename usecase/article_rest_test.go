package usecase

import (
	"encoding/json"
	"testing"

	"github.com/franela/goblin"
)

func TestArticleRESTUseCase(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Article REST Use Case", func() {
		g.Describe("CreateArticle", func() {
			g.Describe("When the request is valid", func() {
				g.It("It should return a response", func() {
					articleREST := &ArticleREST{}
					articleContent := "Article Content"
					articleName := "Article Name"
					expectedResponse := CreateArticleRESTResponse{
						Content: articleContent,
						Name:    articleName,
					}
					request := CreateArticleRESTRequest{
						Content: articleContent,
						Name:    articleName,
					}
					jsonRequest, _ := json.Marshal(request)

					currentResponse, err := articleREST.CreateArticle(jsonRequest)

					g.Assert(currentResponse).Equal(expectedResponse)
					g.Assert(err).Equal(nil)
				})
			})
		})
	})
}
