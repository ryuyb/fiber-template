package repository

import (
	"context"
	"live-poilot/pkg/ent"
)

type UserRepository struct {
	repo *Repository
}

func NewUserRepository(repo *Repository) *UserRepository {
	return &UserRepository{repo: repo}
}

func (ur *UserRepository) GetUser(id uint32) *ent.User {
	user, err := ur.repo.db.User.Get(context.Background(), id)
	if err != nil {
		return nil
	}
	return user
}
