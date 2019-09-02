package services

import (
	"skill_blog/com.github.bobacsmall/datamodels"
	"skill_blog/com.github.bobacsmall/repositories"
)

// 定义方法
type TagService interface {
	EditTag(id uint, tag datamodels.Tag) error
	AddTag(tag datamodels.Tag) error
	DeleteTagById(id uint) error
}

func NewTagService() TagService {
	return &tagService{repo: repositories.NewTagRepository()}
}

type tagService struct {
	repo repositories.TagRepository
}

func (t tagService) EditTag(id uint, tag datamodels.Tag) error {
	return t.repo.EditTag(id, tag)
}

func (t tagService) AddTag(tag datamodels.Tag) error {
	return t.repo.AddTag(tag)
}

func (t tagService) DeleteTagById(id uint) error {
	return t.repo.DeleteTagById(id)
}
