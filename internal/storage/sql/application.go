package sql

import (
	"context"
	"studentRecordsApp/internal/service"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/service/entities"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

var _ service.ApplicationDb = (*Application)(nil)

type Application struct {
	db *sqlx.DB
}

func (a *Application) GetWithInfo(ctx context.Context) ([]entities.ApplicationWithInfo, error) {
	result := make([]sqlEntities.ApplicationWithInfo, 0)
	err := a.db.SelectContext(ctx, &result, `SELECT applications.id, applications._name, applications.created_at,
		applications.link, applications.contact_info,
		applications._text, students.first_name, students.last_name, students.surname, students._group, students.course
    	FROM Applications 
        INNER JOIN Students on applications.student_id = students.id;`)
	if err != nil {
		return nil, err
	}

	entitieResult := make([]entities.ApplicationWithInfo, 0, len(result))
	for _, value := range result {
		entitieResult = append(entitieResult, casts.ApplicationWithInfoToEntite(ctx, value))
	}

	return entitieResult, err
}

func (a *Application) CheckLinkWithUserId(ctx context.Context, id uuid.UUID, link string) (bool, error) {
	var result bool
	err := a.db.GetContext(ctx, &result, "SELECT EXISTS(SELECT * FROM Applications WHERE link = $1 AND student_id = $2)",
		link, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func NewApplication(db *sqlx.DB) *Application {
	return &Application{
		db: db,
	}
}

func (a *Application) Get(ctx context.Context) ([]entities.Application, error) {
	var sqlResults []sqlEntities.Application
	err := a.db.SelectContext(ctx, &sqlResults, "SELECT * FROM Applications ORDER BY created_at")
	if err != nil {
		return nil, err
	}

	results := make([]entities.Application, 0, len(sqlResults))
	for idx := range sqlResults {
		results = append(results, casts.ApplicationSqlToService(ctx, sqlResults[idx]))
	}

	return results, nil
}

func (a *Application) GetForUser(ctx context.Context, userId uuid.UUID) ([]entities.Application, error) {
	var result []sqlEntities.Application
	err := a.db.SelectContext(ctx, &result, "SELECT * FROM Applications WHERE student_id = $1 ORDER BY created_at",
		userId)
	if err != nil {
		return nil, err
	}

	results := make([]entities.Application, 0, len(result))
	for idx := range result {
		results = append(results, casts.ApplicationSqlToService(ctx, result[idx]))
	}

	return results, nil
}

func (a *Application) GetById(ctx context.Context, id, userID uuid.UUID) (entities.Application, error) {
	var result sqlEntities.Application
	err := a.db.GetContext(ctx, &result, "SELECT * FROM Applications WHERE id = $1 AND student_id = $2",
		id, userID)
	if err != nil {
		return entities.Application{}, err
	}

	return casts.ApplicationSqlToService(ctx, result), nil
}

func (a *Application) GetOne(ctx context.Context, id uuid.UUID) (entities.Application, error) {
	var result sqlEntities.Application
	err := a.db.GetContext(ctx, &result, "SELECT * FROM Applications WHERE id = $1",
		id)
	if err != nil {
		return entities.Application{}, err
	}

	return casts.ApplicationSqlToService(ctx, result), nil
}

func (a *Application) Add(ctx context.Context, app entities.Application) error {
	_, err := a.db.ExecContext(ctx, `INSERT INTO Applications (id, student_id, _name, 
                          contact_info, _text, status, created_at, link) 
    										VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		app.Id, app.StudentId, app.Name, app.ContactInfo, app.Text, app.Status, app.CreatedAt, app.Link)

	return err
}

func (a *Application) Update(ctx context.Context, app entities.Application) error {
	_, err := a.db.ExecContext(ctx, "UPDATE Applications SET _name = $1, contact_info = $2, _text = $3 WHERE id = $4",
		app.Name, app.ContactInfo, app.Text, app.Id)

	return err
}

func (a *Application) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	_, err := a.db.ExecContext(ctx, `UPDATE Applications SET status = $1 WHERE id = $2`,
		status, id)

	return err
}

func (a *Application) Delete(ctx context.Context, id, userId uuid.UUID) error {
	_, err := a.db.ExecContext(ctx, "DELETE FROM Applications WHERE id = $1 AND student_id = $2", id, userId)

	return err
}
