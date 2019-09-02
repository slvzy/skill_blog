package datamodels

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name       string `gorm:"not null VARCHAR(191)"`
	Username   string `gorm:"unique;VARCHAR(191)"`
	Password   string `gorm:"not null VARCHAR(191)"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}
