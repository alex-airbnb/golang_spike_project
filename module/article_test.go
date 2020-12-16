package module

import (
	"testing"

	"github.com/franela/goblin"

	"github.com/alex-airbnb/golang_spike_project/adapter"
	"github.com/alex-airbnb/golang_spike_project/entity"
)

func Test(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Article Module", func() {
		g.Describe("AddArticle", func() {
			g.Describe("when receives name and content", func() {
				g.It("it should create a new post", func() {
					article := entity.Article{
						Name:    "ArticleName",
						Content: "ArticleContent",
					}
					externalStorageAdapter := adapter.ExternalStorageFake{}

					currentArticles := CreateArticle(article, externalStorageAdapter)

					g.Assert(len(currentArticles)).Equal(1)
				})
			})
		})
	})
}
