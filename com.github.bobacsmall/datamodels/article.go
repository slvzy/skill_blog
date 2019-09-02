package datamodels

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	TagID         uint    `json:"tag_id" gorm:"index" from:"tagId"`
	Tag           Tag    `json:"tag"`
	Title         string `json:"title" from:"title" validate:"required"`
	Desc          string `json:"desc"  from:"desc"`
	Content       string `json:"content" from:"content"`
	CoverImageUrl string `json:"cover_image_url" from:"coverImageUrl"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state" from:"state"`
}
