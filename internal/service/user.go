package service

import (
	"context"
	"fmt"
	"studentRecordsApp/internal/entites"
	"studentRecordsApp/pkg/password"
)

type UserDb interface {
	AddUser(user entities.entities, ctx context.Context) error
	UpdateUser(user entities.User, ctx context.Context) error
	GetUserByEmailAndRole(email, role string, ctx context.Context) (entities.User, error)
	GetUser(id string, ctx context.Context) (entities.User, error)
	GetUsers(ctx context.Context) ([]entities.User, error)
	DeleteUser(id string, ctx context.Context) error
}

type User struct {
	db UserDb
}

func NewUser(db UserDb) User {
	return User{
		db: db,
	}
}

func (u User) Add(user entities.User, ctx context.Context) error {
	if !user.IsNotEmpty() {
		return fmt.Errorf("400")
	}

	if !user.IsRoleCorrect() {
		return fmt.Errorf("400")
	}

	var err error
	user.Password, err = password.Hash(user.Password)
	if err != nil {
		return err
	}

	return (u.db).AddUser(user, ctx)
}

func (u User) Update(user entities.User, ctx context.Context) error {
	if !user.IsNotEmpty() {
		return fmt.Errorf("400")
	}

	return (u.db).UpdateUser(user, ctx)
}

func (u User) Login(pass, login, role string, ctx context.Context) (entities.User, error) {
	result, err := (u.db).GetUserByEmailAndRole(login, role, ctx)
	if err != nil {
		return entities.User{}, err
	}

	return result, password.CheckHash(pass, []byte(result.Password))
}

func (u User) Get(userId string, ctx context.Context) (entities.User, error) {
	return (u.db).GetUser(userId, ctx)
}

func (u User) GetAllWorker(ctx context.Context) ([]entities.User, error) {
	return (u.db).GetUsers(ctx)
}

func (u User) Delete(userId string, ctx context.Context) error {
	return (u.db).DeleteUser(userId, ctx)
}
