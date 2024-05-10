package sql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

type User struct {
	db *sqlx.DB
}

func (u *User) AddUser(ctx context.Context, user entities.User) error {
	sqlUser := casts.UserEntitieToSql(ctx, user)

	_, err := u.db.ExecContext(ctx,
		`INSERT INTO Users (id, first_name, last_name, surname, email, password, user_role) 
                VALUES($1,$2,$3,$4,$5,$6,$7);`,
		sqlUser.Id, sqlUser.FirstName, sqlUser.LastName, sqlUser.Surname, sqlUser.Email, sqlUser.Password,
		sqlUser.Role,
	)

	return err
}

func (u *User) UpdateUser(ctx context.Context, user entities.User) error {
	sqlUser := casts.UserEntitieToSql(ctx, user)

	_, err := u.db.ExecContext(ctx,
		`UPDATE Users SET first_name=$1, last_name=$2, surname=$3, email=$4 
                 WHERE id = $6;`, sqlUser.FirstName, sqlUser.LastName, sqlUser.Surname, sqlUser.Email,
		sqlUser.Id,
	)

	return err
}

func (u *User) GetUserByEmailAndRole(ctx context.Context, email, role string) (entities.User, error) {
	var result sqlEntities.User

	err := u.db.GetContext(ctx, &result, "SELECT * FROM Users WHERE email=$1 AND user_role=$2;", email, role)
	return casts.UserSqlToEntitie(ctx, result), err
}

func (u *User) GetUser(ctx context.Context, id uuid.UUID) (entities.User, error) {
	var result sqlEntities.User

	err := u.db.GetContext(ctx, &result, `SELECT * FROM Users WHERE id = $1 LIMIT 1;`, id)

	return casts.UserSqlToEntitie(ctx, result), err
}

func (u *User) GetUsersWorker(ctx context.Context) ([]entities.User, error) {
	sqlResults := make([]sqlEntities.User, 0)

	if err := u.db.SelectContext(ctx, &sqlResults, `SELECT * FROM Users WHERE user_role = $1;`, entities.UserWorker); err != nil {
		return nil, err
	}

	results := make([]entities.User, 0, len(sqlResults))

	for _, value := range sqlResults {
		results = append(results, casts.UserSqlToEntitie(ctx, value))
	}

	return results, nil
}

func (u *User) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := u.db.ExecContext(ctx, `DELETE FROM Users WHERE id = $1;`, id)

	return err
}
