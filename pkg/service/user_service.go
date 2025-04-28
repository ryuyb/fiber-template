package service

import (
	"fmt"
	"github.com/jinzhu/copier"
	"live-pilot/api/presenter"
	"live-pilot/pkg/ent"
	"live-pilot/pkg/repository"
	"live-pilot/pkg/utils"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) GetUser(id uint32) (*presenter.UserResp, error) {
	user, err := us.repo.GetUser(id)
	if err != nil {
		return nil, err
	}
	resp := &presenter.UserResp{}
	err = copier.Copy(resp, user)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (us *UserService) SaveUser(u presenter.SaveUserReq) (*ent.User, error) {
	if u.Password != "" {
		hashPassword, err := utils.HashPassword(u.Password)
		if err != nil {
			return nil, fmt.Errorf("error hashing password: %v", err)
		}
		u.Password = hashPassword
	}
	if u.Id == nil {
		return us.repo.CreateUser(u)
	}
	return us.repo.UpdateUser(u)
}

func (us *UserService) GetUserByUsername(username string) (*ent.User, error) {
	return us.repo.GetUserByUsername(username)
}

func (us *UserService) CheckPassword(req presenter.UserLoginReq) (*ent.User, error) {
	user, err := us.GetUserByUsername(req.Username)
	if err != nil || !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, err
	}
	return user, nil
}
