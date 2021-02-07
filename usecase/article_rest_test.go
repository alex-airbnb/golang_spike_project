package usecase

import (
	"testing"

	"github.com/alex-airbnb/golang_spike_project/entity"
	"github.com/franela/goblin"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Create(m interface{}) error {
	args := r.Called(m)

	return args.Error(0)
}

func TestArticleRESTUseCase(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Article REST Use Case", func() {
		g.Describe("CreateArticle", func() {
			articleName := "Article Name"
			articleContent := "Article Content"

			g.Describe("When the request is valid", func() {
				g.It("It should save the article in a reporitory and return a response", func() {
					repositoryMock := &repositoryMock{}
					articleREST := &ArticleREST{Repository: repositoryMock}
					expectedResponse := CreateArticleRESTResponse{
						Content: articleContent,
						Name:    articleName,
					}
					request := []byte(`{
						"content": "Article Content",
						"name": "Article Name"
					}`)
					model := entity.Article{
						Content: articleContent,
						Name:    articleName,
					}
					repositoryMock.On("Create", model).Return(nil)

					currentResponse, err := articleREST.CreateArticle(request)

					g.Assert(repositoryMock.AssertCalled(t, "Create", model)).Equal(true)
					g.Assert(currentResponse).Equal(expectedResponse)
					g.Assert(err).Equal(nil)
				})
			})

			g.Describe("When the request is invalid", func() {
				g.It("It should return ErrInvalidCreateArticleRESTRequest", func() {
					repositoryMock := &repositoryMock{}
					articleREST := &ArticleREST{Repository: repositoryMock}
					request := []byte(`{
						"contnt": "Article Content",
						"name": "Article Name"
					}`)
					// model := entity.Article{
					// 	Content: articleContent,
					// 	Name:    articleName,
					// }
					// repositoryMock.On("Create").Return(nil)

					currentResponse, err := articleREST.CreateArticle(request)

					// g.Assert(repositoryMock.AssertNotCalled(t, "Create")).Equal(true)
					g.Assert(currentResponse).Equal(CreateArticleRESTResponse{})
					g.Assert(err).Equal(ErrInvalidCreateArticleRESTRequest)
				})
			})
		})
	})
}
