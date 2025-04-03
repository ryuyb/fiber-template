package repository

import (
	"context"
	"live-pilot/api/presenter"
	"live-pilot/pkg/ent"
)

type UserRepository struct {
	repo *Repository
}

func NewUserRepository(repo *Repository) *UserRepository {
	return &UserRepository{repo: repo}
}

func (ur *UserRepository) GetUser(id uint32) (*ent.User, error) {
	user, err := ur.repo.db.User.Get(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) CreateUser(u presenter.SaveUserReq) (*ent.User, error) {
	return ur.repo.db.User.Create().
		SetUsername(u.Username).
		SetPassword(u.Password).
		Save(context.Background())
}

func (ur *UserRepository) UpdateUser(u presenter.SaveUserReq) (*ent.User, error) {
	return ur.repo.db.User.UpdateOneID(*u.Id).
		SetUsername(u.Username).
		SetPassword(u.Password).
		Save(context.Background())
}
