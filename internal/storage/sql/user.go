package sql

import (
	"context"
	"studentRecordsApp/internal/storage/sql/sqlEntities"

	"github.com/google/uuid"

	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/service/entites"
)

func (s *Storage) AddUser(user entities.User, ctx context.Context) error {
	sqlUser := casts.UserEntiteToSqlWithoutId(user, ctx)
	sqlUser.Id = uuid.New()

	_, err := s.db.ExecContext(ctx,
		`INSERT INTO Users (id, first_name, last_name, surname, email, password, medecine_book, user_role) 
                VALUES($1,$2,$3,$4,$5,$6,$7,$8);`,
		sqlUser.Id, sqlUser.FirstName, sqlUser.LastName, sqlUser.Surname, sqlUser.Email, sqlUser.Password,
		sqlUser.Medicine, sqlUser.Role,
	)

	return err
}

func (s *Storage) UpdateUser(user entities.User, ctx context.Context) error {
	sqlUser, err := casts.UserEntiteToSql(user, ctx)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx,
		`UPDATE Users SET first_name=$1, last_name=$2, surname=$3, email=$4, medecine_book=$5 
                 WHERE id = $6;`, sqlUser.FirstName, sqlUser.LastName, sqlUser.Surname, sqlUser.Email, sqlUser.Medicine,
		sqlUser.Id,
	)

	return err
}

func (s *Storage) GetUserByEmailAndRole(email, role string, ctx context.Context) (entities.User, error) {
	var result entities.User

	err := s.db.QueryRowContext(ctx, `SELECT * FROM Users WHERE email = $1 AND user_role = $2;`, email, role).
		Scan(&result)

	return result, err
}

func (s *Storage) GetUser(id string, ctx context.Context) (entities.User, error) {
	var result entities.User

	err := s.db.QueryRowContext(ctx, `SELECT * FROM Users WHERE id = $1;`, id).Scan(&result)

	return result, err
}

func (s *Storage) GetUsers(ctx context.Context) ([]entities.User, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT * FROM Users;`)
	if err != nil {
		return nil, err
	}

	results := make([]entities.User, 0)

	for rows.Next() {
		var result sqlEntities.User
		err := rows.Scan(&result)
		if err != nil {
			return nil, err
		}

		results = append(results, casts.UserSqlToEntite(result, ctx))
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
