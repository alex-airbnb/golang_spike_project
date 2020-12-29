package module

import (
	"testing"

	"github.com/alex-airbnb/golang_spike_project/model"
	"github.com/franela/goblin"
	"github.com/stretchr/testify/mock"
)

type postgresAdapterMock struct {
	mock.Mock
}

func (mock *postgresAdapterMock) Create(item interface{}) error {
	args := mock.Called(item)

	return args.Error(0)
}

func Test(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Article Module", func() {
		g.Describe("CreateArticle", func() {
			g.It("should create a new article with the specified name and content", func() {
				expectedArticle := &model.Article{
					Name:    "Article Name",
					Content: "Article Content",
				}
				adapter := &postgresAdapterMock{}

				adapter.On("Create", expectedArticle).Return(nil)

				currentArticle, currentErr := CreateArticle(expectedArticle, adapter)

				g.Assert(currentArticle).Equal(expectedArticle)
				g.Assert(currentErr).Equal(nil)
			})
		})
	})
}
