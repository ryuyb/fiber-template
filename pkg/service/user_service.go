package service

import (
	"live-poilot/pkg/ent"
	"live-poilot/pkg/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) GetUser(id uint32) *ent.User {
	return us.repo.GetUser(id)
}
