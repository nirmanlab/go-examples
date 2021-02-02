package article

import "fmt"

type Article struct {
	NameArticle  string `json:"name" bson:"bname" db:"name_article"`
	numberOfPage int64
}

// NewArticle will create new article
func NewArticle(name string, numOfPage int64) Article {
	return Article{
		NameArticle:  name,
		numberOfPage: numOfPage,
	}
}

// ToString prints information about article struct
func (article Article) ToString() {
	fmt.Printf("Name of Article :- %v \n", article.NameArticle)
	fmt.Printf("number of page in article :- %v \n", article.numberOfPage)
}
