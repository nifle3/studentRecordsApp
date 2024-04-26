package service

import (
	"context"
	"fmt"

	"studentRecordsApp/internal/service/entites"
)

type UserDb interface {
	AddUser(user entities.User, ctx context.Context) error
	UpdateUser(user entities.User, ctx context.Context) error
	GetUserByEmailAndRole(email, role string, ctx context.Context) (entities.User, error)
	GetUser(id string, ctx context.Context) (entities.User, error)
	GetUsers(ctx context.Context) ([]entities.User, error)
	DeleteUser(id string, ctx context.Context) error
}

type User struct {
	db *UserDb
}

func NewUser(db UserDb) User {
	return User{
		db: &db,
	}
}

func (u User) Add(user entities.User, ctx context.Context) error {
	if !user.CheckIsNotEmpty() {
		return fmt.Errorf("400")
	}

	if !user.CheckRole() {
		return fmt.Errorf("400")
	}

	err := user.HashPassword()
	if err != nil {
		return err
	}

	return (*u.db).AddUser(user, ctx)
}

func (u User) Update(user entities.User, ctx context.Context) error {
	if !user.CheckIsNotEmpty() {
		return fmt.Errorf("400")
	}

	return (*u.db).UpdateUser(user, ctx)
}

func (u User) Login(password, login, role string, ctx context.Context) (entities.User, error) {
	result, err := (*u.db).GetUserByEmailAndRole(login, role, ctx)
	if err != nil {
		return entities.User{}, err
	}

	return result, result.CheckHash(password)
}

func (u User) Get(userId string, ctx context.Context) (entities.User, error) {
	return (*u.db).GetUser(userId, ctx)
}

func (u User) GetAll(ctx context.Context) ([]entities.User, error) {
	return (*u.db).GetUsers(ctx)
}

func (u User) Delete(userId string, ctx context.Context) error {
	return (*u.db).DeleteUser(userId, ctx)
}
