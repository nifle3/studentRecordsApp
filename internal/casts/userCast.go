package casts

import (
	"context"

	"studentRecordsApp/internal/service/entities"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
	"studentRecordsApp/internal/transport/server/httpEntity"
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

func UserEntitieToHttp(_ context.Context, user entities.User) (httpEntity.User, error) {
	id, err := UuidToString(user.Id)
	if err != nil {
		return httpEntity.User{}, err
	}

	return httpEntity.User{
		Id:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Surname:   user.Surname,
		Email:     user.Email,
		Role:      user.Role,
	}, nil
}
