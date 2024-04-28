package casts

import (
	"context"
	"studentRecordsApp/internal/entites"
	"studentRecordsApp/internal/transport/server/jsonStruct"

	"github.com/google/uuid"

	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func UserEntiteToSql(user entities.entities, _ context.Context) (sqlEntities.User, error) {
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
	}
}

func UserEntitieToJson(user entities.User, _ context.Context) jsonStruct.User {
	return jsonStruct.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Surname:   user.Surname,
		Email:     user.Email,
	}
}

func UserJsonToEntitie(user jsonStruct.User, role, id string, _ context.Context) entities.User {
	return entities.User{
		Id:        id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Surname:   user.Surname,
		Email:     user.Email,
		Role:      role,
	}
}

func UserEntitieToJsonWithId(user entities.User, _ context.Context) jsonStruct.UserWithId {
	return jsonStruct.UserWithId{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Surname:   user.Surname,
		Email:     user.Email,
	}
}

func UserJsonWithIdToEntitie(user jsonStruct.UserWithId, role string, _ context.Context) entities.User {
	return entities.User{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Surname:   user.Surname,
		Email:     user.Email,
		Role:      role,
	}
}
