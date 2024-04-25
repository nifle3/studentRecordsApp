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
                VALUES(?,?,?,?,?,?,?,?)`,
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
		`UPDATE Users SET first_name=?, last_name=?, surname=?, email=?, medecine_book=? 
                 WHERE id = ?`, sqlUser.FirstName, sqlUser.LastName, sqlUser.Surname, sqlUser.Email, sqlUser.Medicine,
		sqlUser.Id,
	)

	return err
}

func (s *Storage) GetUserByEmail(email string, ctx context.Context) (entities.User, error) {
	var result entities.User

	err := s.db.QueryRowContext(ctx, `SELECT * FROM Users WHERE email = ?`, email).Scan(&result)

	return result, err
}

func (s *Storage) GetUser(id string, ctx context.Context) (entities.User, error) {
	var result entities.User

	err := s.db.QueryRowContext(ctx, `SELECT * FROM Users WHERE id = ?`, id).Scan(&result)

	return result, err
}

func (s *Storage) GetUsers(ctx context.Context) ([]entities.User, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT * FROM Users`)
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

	_, err = s.db.ExecContext(ctx, `DELETE FROM Users WHERE id = ?`, uuId)

	return err
}
