package repositories

import (
	"github.com/jinzhu/gorm"
	"skill_blog/com.github.bobacsmall/datamodels"
	"skill_blog/com.github.bobacsmall/datasource"
)

type TagRepository interface {
	ExistTagByName(name string) (bool, error)
	ExistTagById(id uint) (bool, error)
	GetTagTotal(tag datamodels.Tag) (int, error)
	GetTag(id uint) (*datamodels.Tag, error)
	EditTag(id uint, tag datamodels.Tag) error
	AddTag(tag datamodels.Tag) error
	DeleteTagById(id uint) error
	CleanAllTag() error
	GetTags(pageNum int, pageSize int, tag datamodels.Tag) ([]*datamodels.Tag, error)
}

func NewTagRepository() TagRepository {
	return &tagRepository{db: datasource.DB}
}

type tagRepository struct {
	db *gorm.DB
}

func (t tagRepository) ExistTagByName(name string) (bool, error) {
	var tag datamodels.Tag
	err := t.db.Select("id").Where("name = ?", name).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

func (t tagRepository) ExistTagById(id uint) (bool, error) {
	var tag datamodels.Tag

	err := t.db.Select("id").Where("id = ?", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

func (t tagRepository) GetTagTotal(tag datamodels.Tag) (int, error) {
	var count int
	if err := t.db.Model(&tag).Where(tag).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t tagRepository) GetTag(id uint) (*datamodels.Tag, error) {
	var tag datamodels.Tag
	err := t.db.Where("id=?", id).First(&tag).Error
	return &tag, err
}

func (t tagRepository) EditTag(id uint, tag datamodels.Tag) error {
	return t.db.Model(&tag).Where("id = ?", id).Update(&tag).Error
}

func (t tagRepository) AddTag(tag datamodels.Tag) error {
	return t.db.Create(&tag).Error
}

func (t tagRepository) DeleteTagById(id uint) error {
	return t.db.Where("id=?", id).Delete(&datamodels.Tag{}).Error
}

func (t tagRepository) CleanAllTag() error {
	return t.db.Unscoped().Delete(&datamodels.Tag{}).Error
}

func (t tagRepository) GetTags(pageNum int, pageSize int, tag datamodels.Tag) ([]*datamodels.Tag, error) {
	var (
		tags []*datamodels.Tag
		err  error
	)

	if pageSize > 0 && pageNum > 0 {
		err = t.db.Where(&tag).Find(&tags).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = t.db.Where(&tag).Find(&tags).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}
