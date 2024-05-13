package service

import (
	"context"
	"net/http"
	"studentRecordsApp/pkg/customError"

	"github.com/google/uuid"

    "studentRecordsApp/internal/service/entities"
	"studentRecordsApp/pkg/password"
)

type UserDb interface {
	Add(ctx context.Context, user entities.User) error
	Update(ctx context.Context, user entities.User) error
	Get(ctx context.Context, id uuid.UUID) (entities.User, error)
	GetAll(ctx context.Context, role string) ([]entities.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type User struct {
	db *UserDb
}

func NewUser(db UserDb) User {
	return User{
		db: &db,
	}
}

func (u User) Add(ctx context.Context, user entities.User) *customError.Http {
	if !user.IsNotEmpty() || !user.CheckEmail() {
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

	err = (*u.db).Add(ctx, user)
	if err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (u User) Update(ctx context.Context, user entities.User) *customError.Http {
	if !user.IsNotUpdateEmpty() || !user.CheckEmail() {
		return customError.New(http.StatusBadRequest, "Has some invalid fields")
	}

	if err := (*u.db).Update(ctx, user); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (u User) Get(ctx context.Context, userId uuid.UUID) (entities.User, *customError.Http) {
	result, err := (*u.db).Get(ctx, userId)
	if err != nil {
		return entities.User{}, customError.New(http.StatusInternalServerError, err.Error())
	}

	return result, nil
}

func (u User) GetAllWorker(ctx context.Context) ([]entities.User, *customError.Http) {
	result, err := (*u.db).GetAll(ctx, entities.UserWorker)
	if err != nil {
		return nil, customError.New(http.StatusInternalServerError, err.Error())
	}

	return result, nil
}

func (u User) Delete(ctx context.Context, userId uuid.UUID) *customError.Http {
	if err := (*u.db).Delete(ctx, userId); err != nil {
		return customError.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}
