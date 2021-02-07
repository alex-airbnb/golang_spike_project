package entity

import (
	"testing"

	"github.com/franela/goblin"
)

func TestArticleEntity(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Article Entity", func() {
		g.Describe("CreateArticle", func() {
			g.It("It should create a new article with the specified name and content", func() {
				articleName := "Article Name"
				articleContent := "Article Content"
				expectedArticle := Article{Name: articleName, Content: articleContent}

				currentArticle, err := CreateArticle(articleContent, articleName)

				g.Assert(currentArticle).Equal(expectedArticle)
				g.Assert(err).Equal(nil)
			})
		})
	})
}
