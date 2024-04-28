package sql

import (
	"context"
	"log"
	"studentRecordsApp/internal/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"

	"github.com/google/uuid"

	"studentRecordsApp/internal/casts"
)

func (s *Storage) AddUser(user entities.User, ctx context.Context) error {
	sqlUser := casts.UserEntiteToSqlWithoutId(user, ctx)
	sqlUser.Id = uuid.New()

	_, err := s.db.ExecContext(ctx,
		`INSERT INTO Users (id, first_name, last_name, surname, email, password, user_role) 
                VALUES($1,$2,$3,$4,$5,$6,$7);`,
		sqlUser.Id, sqlUser.FirstName, sqlUser.LastName, sqlUser.Surname, sqlUser.Email, sqlUser.Password,
		sqlUser.Role,
	)

	return err
}

func (s *Storage) UpdateUser(user entities.User, ctx context.Context) error {
	sqlUser, err := casts.UserEntiteToSql(user, ctx)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx,
		`UPDATE Users SET first_name=$1, last_name=$2, surname=$3, email=$4 
                 WHERE id = $6;`, sqlUser.FirstName, sqlUser.LastName, sqlUser.Surname, sqlUser.Email,
		sqlUser.Id,
	)

	return err
}

func (s *Storage) GetUserByEmailAndRole(email, role string, ctx context.Context) (entities.User, error) {
	var result sqlEntities.User

	var err = s.db.GetContext(ctx, &result, "SELECT * FROM Users WHERE email=$1 AND user_role=$2;", email, role)
	log.Printf("%#v", result)
	return casts.UserSqlToEntite(result, ctx), err
}

func (s *Storage) GetUser(id string, ctx context.Context) (entities.User, error) {
	var result sqlEntities.User

	err := s.db.GetContext(ctx, &result, `SELECT * FROM Users WHERE id = $1;`, id)

	return casts.UserSqlToEntite(result, ctx), err
}

func (s *Storage) GetUsers(ctx context.Context) ([]entities.User, error) {
	sqlResults := make([]sqlEntities.User, 0)

	err := s.db.SelectContext(ctx, &sqlResults, `SELECT * FROM Users WHERE user_role = $1;`, entities.entities.UserWorker)
	if err != nil {
		return nil, err
	}

	results := make([]entities.User, 0, len(sqlResults))

	for _, value := range sqlResults {
		results = append(results, casts.UserSqlToEntite(value, ctx))
	}

	return results, nil
}

func (s *Storage) DeleteUser(id string, ctx context.Context) error {
	uuId, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, `DELETE FROM Users WHERE id = $1;`, uuId)

	return err
}
