package service

import (
	"LangBot/Internal/models"
	"LangBot/Internal/repository"
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUserService(user *models.User) error {
	err := u.repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) GetUserService(id int64) (*models.User, error) {
	user, err := u.repo.GetByTgId(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
