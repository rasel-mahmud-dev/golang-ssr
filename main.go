package main

import (
	"awesomeProject/models"
	"fmt"
)

func main() {

	var article = models.InitArticleRepository()
	//var users = models.InitUserRepository()

	//article.AddArticle(models.Article{
	//	Title:         "Test projeject",
	//	Content:       "23sdff",
	//	AuthorID:      1,
	//	Slug:          "f34",
	//	CoverImageURL: "",
	//	Status:        "",
	//	WordCount:     0,
	//	ReadTime:      0,
	//})

	ar := article.GetArticles()

	fmt.Println(ar)

}
