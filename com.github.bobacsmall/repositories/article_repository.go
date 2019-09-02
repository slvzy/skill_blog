package repositories

import (
	"github.com/jinzhu/gorm"
	"skill_blog/com.github.bobacsmall/datamodels"
	"skill_blog/com.github.bobacsmall/datasource"
)

type ArticleRepository interface {
	ExistArticleById(id uint) (bool, error)
	GetArticleTotal(article datamodels.Article) (int, error)
	GetArticle(id uint) (*datamodels.Article, error)
	EditArticle(id uint, article datamodels.Article) error
	AddArticle(article datamodels.Article) error
	DeleteArticleById(id uint) error
	CleanAllArticle() error
	GetArticles(pageNum int, pageSize int, article datamodels.Article) ([]*datamodels.Article, error)
}

func NewArticleRepository() ArticleRepository {
	return &articleRepository{db: datasource.DB}

}

type articleRepository struct {
	db *gorm.DB
}

func (a articleRepository) ExistArticleById(id uint) (bool, error) {
	var article datamodels.Article

	err := a.db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}

func (a articleRepository) GetArticleTotal(article datamodels.Article) (int, error) {
	var count int
	if err := a.db.Model(&article).Where(&article).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (a articleRepository) GetArticle(id uint) (*datamodels.Article, error) {
	var article datamodels.Article
	err := a.db.Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	err = a.db.Model(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}

func (a articleRepository) EditArticle(id uint, article datamodels.Article) error {
	return a.db.Model(&datamodels.Article{}).Where("id = ? AND deleted_on = ? ", id, 0).Update(&article).Error
}

func (a articleRepository) AddArticle(article datamodels.Article) error {
	return a.db.Create(&article).Error
}

func (a articleRepository) DeleteArticleById(id uint) error {
	return a.db.Where("id=?", id).Delete(&datamodels.Article{}).Error
}

func (a articleRepository) CleanAllArticle() error {
	return a.db.Unscoped().Delete(&datamodels.Article{}).Error
}

func (a articleRepository) GetArticles(pageNum int, pageSize int, article datamodels.Article) ([]*datamodels.Article, error) {
	var (
		articles []*datamodels.Article
		err      error
	)
	if pageSize > 0 && pageNum > 0 {
		err = a.db.Preload("Tag").Where(&article).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	} else {
		err = a.db.Preload("Tag").Where(&article).Find(&articles).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articles, nil
}
