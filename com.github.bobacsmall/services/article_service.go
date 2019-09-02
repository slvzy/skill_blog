package services

import (
	"errors"
	"skill_blog/com.github.bobacsmall/datamodels"
	"skill_blog/com.github.bobacsmall/repositories"
)

type ArticleService interface {
	ExistArticleById(id uint) (bool, error)
	GetArticleTotal(article datamodels.Article) (int, error)
	GetArticle(id uint) (*datamodels.Article, error)
	EditArticle(id uint, article datamodels.Article) error
	AddArticle(article datamodels.Article) error
	DeleteArticleById(id uint) error
	CleanAllArticle() error
	GetArticles(pageNum int, pageSize int, article datamodels.Article) ([]*datamodels.Article, error)
}

func NewArticleService() ArticleService {
	return &articleService{
		repo:    repositories.NewArticleRepository(),
		tagRepo: repositories.NewTagRepository(),
	}
}

type articleService struct {
	repo    repositories.ArticleRepository
	tagRepo repositories.TagRepository
}

func (a articleService) ExistArticleById(id uint) (bool, error) {
	// add other logic
	return a.repo.ExistArticleById(id)
}

func (a articleService) GetArticleTotal(article datamodels.Article) (int, error) {
	// add other logic
	return a.repo.GetArticleTotal(article)
}

func (a articleService) GetArticle(id uint) (*datamodels.Article, error) {
	return a.repo.GetArticle(id)
}

func (a articleService) EditArticle(id uint, article datamodels.Article) error {
	return a.repo.EditArticle(id, article)
}

func (a articleService) AddArticle(article datamodels.Article) error {
	bool, _ := a.tagRepo.ExistTagById(article.TagID)
	if !bool {
		return errors.New("tag type error")
	}

	return a.repo.AddArticle(article)
}

func (a articleService) DeleteArticleById(id uint) error {
	return a.repo.DeleteArticleById(id)
}

func (a articleService) CleanAllArticle() error {
	return a.repo.CleanAllArticle()
}

func (a articleService) GetArticles(pageNum int, pageSize int, article datamodels.Article) ([]*datamodels.Article, error) {
	// add other logic
	return a.repo.GetArticles(pageNum, pageSize, article)
}
