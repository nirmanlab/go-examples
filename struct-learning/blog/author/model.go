package author

import (
	"fmt"
	"structlearning/blog/article"
)

type Author struct {
	NameAuthor      string
	numberOfArticle int64
	article         article.Article
}

// NewAuthor creates new author
func NewAuthor(name string, numOfArticles int64, article article.Article) Author {
	return Author{
		NameAuthor:      name,
		numberOfArticle: numOfArticles,
		article:         article,
	}
}

func (author Author) ToString() {
	fmt.Printf("Name of Author :- %v \n", author.NameAuthor)
	fmt.Printf("Number of article :- %v \n", author.numberOfArticle)
	author.article.ToString()
}
