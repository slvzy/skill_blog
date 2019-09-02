package repositories

import (
	"fmt"
	"skill_blog/com.github.bobacsmall/datamodels"
	"skill_blog/com.github.bobacsmall/datasource"
	"skill_blog/com.github.bobacsmall/pkg/setting"
	"testing"
)

func TestArticleRepository_AddArticle(t *testing.T) {
	setting.Setup()
	datasource.Setup()
	repository := NewArticleRepository()
	article := repository.AddArticle(datamodels.Article{
		TagID:   6,
		Title:   "JVM深入理解",
		Desc:    "test",
		Content: "hahahahahahhaahhahah",
	})

	fmt.Println(article)
}

func TestArticleRepository_GetArticleTotal(t *testing.T) {
	setting.Setup()
	datasource.Setup()
	repository := NewArticleRepository()
	i, e := repository.GetArticleTotal(datamodels.Article{})
	fmt.Println("=====================>",i,e)
}

func TestArticleRepository_ExistArticleById(t *testing.T) {
	setting.Setup()
	datasource.Setup()
	repository := NewArticleRepository()
	b,e := repository.ExistArticleById(1)
	fmt.Println("----->",b,e)
}

func TestArticleRepository_GetArticle(t *testing.T) {
	setting.Setup()
	datasource.Setup()
	repository := NewArticleRepository()
	article, e := repository.GetArticle(1)
	fmt.Println("------>",article,e)
}

func TestArticleRepository_EditArticle(t *testing.T) {
	setting.Setup()
	datasource.Setup()
	repository := NewArticleRepository()
	article := repository.EditArticle(1, datamodels.Article{
		Desc:    "如果深入理解JVM",
		Content: "请参考xxxx",
	})
	fmt.Println(article)
}

func TestArticleRepository_DeleteArticleById(t *testing.T) {
	setting.Setup()
	datasource.Setup()
	repository := NewArticleRepository()
	id := repository.DeleteArticleById(1)
	fmt.Println(id)
}

func TestArticleRepository_GetArticles(t *testing.T) {
	setting.Setup()
	datasource.Setup()
	repository := NewArticleRepository()
	articles, e := repository.GetArticles(0, 0, datamodels.Article{})
	fmt.Println(articles[0],e)
}