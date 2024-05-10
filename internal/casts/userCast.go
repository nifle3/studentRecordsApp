package casts

import (
	"context"
	entities "studentRecordsApp/internal/service/entites"

	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func UserSqlToEntitie(_ context.Context, user sqlEntities.User) entities.User {
	return entities.User{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Surname:   user.Surname,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
	}
}

func UserEntitieToSql(_ context.Context, user entities.User) sqlEntities.User {
	return sqlEntities.User{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Surname:   user.Surname,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
	}
}
