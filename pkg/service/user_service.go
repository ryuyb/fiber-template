package service

import (
	"live-pilot/api/presenter"
	"live-pilot/pkg/ent"
	"live-pilot/pkg/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) GetUser(id uint32) (*ent.User, error) {
	return us.repo.GetUser(id)
}

func (us *UserService) SaveUser(u presenter.SaveUserReq) (*ent.User, error) {
	if u.Id == nil {
		return us.repo.CreateUser(u)
	}
	return us.repo.UpdateUser(u)
}
