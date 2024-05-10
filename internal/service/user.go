package service

import (
	"context"
	"net/http"
	"studentRecordsApp/pkg/customError"

	"github.com/google/uuid"

	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/pkg/password"
)

type UserDb interface {
	AddUser(user entities.User, ctx context.Context) error
	UpdateUser(ctx context.Context, user entities.User) error
	GetUser(ctx context.Context, id uuid.UUID) (entities.User, error)
	GetUsers(ctx context.Context) ([]entities.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type User struct {
	db UserDb
}

func NewUser(db UserDb) User {
	return User{
		db: db,
	}
}

func (u User) Add(ctx context.Context, user entities.User) error {
	if !user.IsNotEmpty() {
		return customError.New(http.StatusBadRequest, "Has some invalid fields")
	}

	if !user.IsRoleCorrect() {
		return customError.New(http.StatusBadRequest, "Has an invalid role")
	}

	user.Id = uuid.New()

	var err error
	user.Password, err = password.Hash(user.Password)
	if err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	return (u.db).AddUser(user, ctx)
}

func (u User) Update(ctx context.Context, user entities.User) error {
	if !user.IsNotEmpty() {
		return customError.New(http.StatusBadRequest, "Has some invalid fields")
	}

	if !user.IsRoleCorrect() {
		return customError.New(http.StatusBadRequest, "Has an invalid role")
	}

	return (u.db).UpdateUser(ctx, user)
}

func (u User) Get(ctx context.Context, userId uuid.UUID) (entities.User, error) {
	return (u.db).GetUser(ctx, userId)
}

func (u User) GetAllWorker(ctx context.Context) ([]entities.User, error) {
	return (u.db).GetUsers(ctx)
}

func (u User) Delete(ctx context.Context, userId uuid.UUID) error {
	return (u.db).DeleteUser(ctx, userId)
}
