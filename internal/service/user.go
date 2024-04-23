package service

import (
	"context"
	"studentRecordsApp/internal/service/entites"
)

type UserDb interface {
	AddUser(user entities.User, ctx context.Context) error
	UpdateUser(user entities.User, ctx context.Context) error
	LoginUser(password, login string, ctx context.Context) (bool, error)
	GetUser(userId string, ctx context.Context) (User, error)
	GetUsers(ctx context.Context) ([]User, error)
	DeleteUser(userId string, ctx context.Context) error
}

type User struct {
	db *UserDb
}
