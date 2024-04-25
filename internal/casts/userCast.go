package casts

import (
	"context"

	"github.com/google/uuid"

	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func UserEntiteToSql(user entities.User, _ context.Context) (sqlEntities.User, error) {
	id, err := uuid.Parse(user.Id)
	if err != nil {
		return sqlEntities.User{}, err
	}

	return sqlEntities.User{
		Id:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Surname:   user.Surname,
		Email:     user.Email,
		Password:  user.Password,
		Medicine:  user.Medicine,
	}, nil
}

func UserEntiteToSqlWithoutId(user entities.User, _ context.Context) sqlEntities.User {
	return sqlEntities.User{
		Id:        uuid.UUID{},
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Surname:   user.Surname,
		Email:     user.Email,
		Password:  user.Password,
		Medicine:  user.Medicine,
	}
}

func UserSqlToEntite(user sqlEntities.User, _ context.Context) entities.User {
	return entities.User{
		Id:        user.Id.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Surname:   user.Surname,
		Email:     user.Email,
		Password:  user.Password,
		Medicine:  user.Medicine,
	}
}
