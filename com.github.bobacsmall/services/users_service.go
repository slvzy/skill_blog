package services

import (
	"skill_blog/com.github.bobacsmall/datamodels"
	"skill_blog/com.github.bobacsmall/repositories"
)

type UserService interface {
	Register(user datamodels.User) error
	Login(user datamodels.User) (*datamodels.User, error)
}

func NewUserService() UserService {
	return &userService{rep: repositories.NewUsersRepository()}
}

type userService struct {
	rep repositories.UsersRepository
}

func (u userService) Register(user datamodels.User) error {
	return u.rep.SaveUser(user)
}

func (u userService) Login(user datamodels.User) (*datamodels.User, error) {
	uu, err := u.rep.CheckLogin(user.Username, user.Password)
	return uu, err
}
