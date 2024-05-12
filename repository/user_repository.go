package repository

import (
	"context"
	"next-learn-go/model"

	"github.com/uptrace/bun"
)

type IUserRepository interface {
	GetUserByEmail(ctx context.Context, user *model.User, email string) error
	CreateUser(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, user *model.User, userId uint) error
}

type userRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(ctx context.Context, user *model.User, email string) error {
	if err := ur.db.NewSelect().
		Model(user).
		Where("email=?", email).
		Scan(ctx); err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	if _, err := ur.db.NewInsert().Model(user).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserById(ctx context.Context, user *model.User, userId uint) error {
	if err := ur.db.NewSelect().
		Model(user).
		Where("id=?", userId).
		Scan(ctx); err != nil {
		return err
	}
	return nil
}
