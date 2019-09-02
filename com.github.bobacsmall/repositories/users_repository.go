package repositories

import (
	"github.com/jinzhu/gorm"
	"skill_blog/com.github.bobacsmall/datamodels"
	"skill_blog/com.github.bobacsmall/datasource"
)

type UsersRepository interface {
	SaveUser(user datamodels.User) error
	CheckLogin(username, password string) (*datamodels.User, error)
	EditUser(id uint, user datamodels.User) error
}

func NewUsersRepository() UsersRepository {
	return &usersRepository{db: datasource.DB}
}

type usersRepository struct {
	db *gorm.DB
}

func (u usersRepository) SaveUser(user datamodels.User) error {
	return u.db.Create(&user).Error
}

func (u usersRepository) CheckLogin(username, password string) (*datamodels.User, error) {
	var user datamodels.User
	err := u.db.Model(&datamodels.User{}).Where("username=? and password=?", username, password).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, err
}

func (u usersRepository) EditUser(id uint, user datamodels.User) error {
	return u.db.Model(&datamodels.Article{}).Where("id = ?", id).Update(&user).Error
}
