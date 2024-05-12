package sql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/service"
	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

var _ service.UserDb = (*User)(nil)

type User struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) Add(ctx context.Context, user entities.User) error {
	sqlUser := casts.UserEntitieToSql(ctx, user)

	_, err := u.db.ExecContext(ctx,
		`INSERT INTO Users (id, first_name, last_name, surname, email, password, user_role) 
                VALUES($1,$2,$3,$4,$5,$6,$7);`,
		sqlUser.Id, sqlUser.FirstName, sqlUser.LastName, sqlUser.Surname, sqlUser.Email, sqlUser.Password,
		sqlUser.Role,
	)

	return err
}

func (u *User) Update(ctx context.Context, user entities.User) error {
	sqlUser := casts.UserEntitieToSql(ctx, user)

	_, err := u.db.ExecContext(ctx,
		`UPDATE Users SET first_name=$1, last_name=$2, surname=$3, email=$4 
                 WHERE id = $5;`, sqlUser.FirstName, sqlUser.LastName, sqlUser.Surname, sqlUser.Email,
		sqlUser.Id,
	)

	return err
}

func (u *User) Auth(ctx context.Context, role, email string) (uuid.UUID, string, error) {
	var result sqlEntities.User

	err := u.db.GetContext(ctx, &result, "SELECT * FROM Users WHERE email=$1 AND user_role=$2;", email, role)
	return result.Id, result.Password, err
}

func (u *User) Get(ctx context.Context, id uuid.UUID) (entities.User, error) {
	var result sqlEntities.User

	err := u.db.GetContext(ctx, &result, `SELECT * FROM Users WHERE id = $1 LIMIT 1;`, id)

	return casts.UserSqlToEntitie(ctx, result), err
}

func (u *User) GetAll(ctx context.Context, role string) ([]entities.User, error) {
	sqlResults := make([]sqlEntities.User, 0)

	if err := u.db.SelectContext(ctx, &sqlResults, `SELECT * FROM Users WHERE user_role = $1;`, role); err != nil {
		return nil, err
	}

	results := make([]entities.User, 0, len(sqlResults))

	for _, value := range sqlResults {
		results = append(results, casts.UserSqlToEntitie(ctx, value))
	}

	return results, nil
}

func (u *User) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := u.db.ExecContext(ctx, `DELETE FROM Users WHERE id = $1;`, id)

	return err
}
