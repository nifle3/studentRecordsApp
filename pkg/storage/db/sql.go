package db

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"studentRecordsApp/internal/service/entities"
	"studentRecordsApp/pkg/password"
)

var instance *sqlx.DB
var once sync.Once

func MustNewSqlConnection(ctx context.Context, connection string) *sqlx.DB {
	once.Do(func() {
		db, err := sqlx.ConnectContext(ctx, "postgres", connection)
		if err != nil {
			panic(err)
		}

		if err := db.PingContext(ctx); err != nil {
			panic(err.Error())
		}

		mustStartData(db)
		instance = db
	})

	return instance
}

func mustStartData(db *sqlx.DB) {
	var exist bool
	err := db.GetContext(context.Background(), &exist, `SELECT EXISTS(SELECT * FROM Users)`)

	if err != nil {
		panic(err.Error())
	}

	if exist {
		return
	}

	pass, err := password.Hash("qwe")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.ExecContext(context.Background(),
		`INSERT INTO Users VALUES ($1,$2,$3,$4,$5,$6,$7),($8,$9,$10,$11,$12,$13,$14),($15,$16,$17,$18,$19,$20,$21)`,
		uuid.New(), "Артём", "Куприянов", "Сергеевич", "nifle3@gmail.com", pass, entities.UserAdmin,
		uuid.New(), "Раниль", "Закиров", "Ильдусович", "homya@gmail.com", pass, entities.UserWorker,
		uuid.New(), "Антон", "Яковлев", "Дмитриевич", "mersya@gmail.com", pass, entities.UserWorker)

	if err != nil {
		panic(err.Error())
	}
}
