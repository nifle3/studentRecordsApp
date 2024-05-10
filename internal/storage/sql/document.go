package sql

import (
	"context"
	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/storage/sql/sqlEntities"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"studentRecordsApp/internal/service/entites"
)

type Document struct {
	db *sqlx.DB
}

func NewDocument(db *sqlx.DB) *Document {
	return &Document{
		db: db,
	}
}

func (d *Document) GetById(ctx context.Context, id uuid.UUID) (entities.Document, error) {
	var result sqlEntities.StudentsDocument
	err := d.db.SelectContext(ctx, &result, `SELECT * FROM StudentsDocuments WHERE id = $1`, id)
	if err != nil {
		return entities.Document{}, err
	}

	return casts.DocumentSqlToEntite(ctx, result), nil
}

func (d *Document) GetByUserId(ctx context.Context, userId uuid.UUID) ([]entities.Document, error) {
	var sqlResults []sqlEntities.StudentsDocument
	err := d.db.SelectContext(ctx, &sqlResults, `SELECT * FROM StudentsDocuments WHERE student_id = $1`, userId)
	if err != nil {
		return nil, err
	}

	documents := make([]entities.Document, 0, len(sqlResults))
	for idx := range sqlResults {
		documents = append(documents, casts.DocumentSqlToEntite(ctx, sqlResults[idx]))
	}

	return documents, nil
}

func (d *Document) Add(ctx context.Context, document entities.Document) error {
	_, err := d.db.ExecContext(ctx, `INSERT INTO StudentsDocuments (id, student_id, document_name, document_type, document_link_s3) 
				VALUES ($1, $2, $3, $4, $5)`, document.Id, document.StudentId, document.Name, document.Type, document.Link)
	return err
}

func (d *Document) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := d.db.ExecContext(ctx, `DELETE FROM StudentsDocuments WHERE id = $1`, id)
	return err
}

func (d *Document) DeleteWithUserId(ctx context.Context, id, userId uuid.UUID) error {
	_, err := d.db.ExecContext(ctx, `DELETE FROM StudentsDocuments WHERE id = $1 AND student_id = $2`, id, userId)
	return err
}

func (d *Document) Update(ctx context.Context, document entities.Document) error {
	_, err := d.db.ExecContext(ctx, `UPDATE StudentsDocuments SET document_name = $1, document_type = $2, document_link_s3 = $3 
                         WHERE id = $4`, document.Name, document.Type, document.Link, document.Id)
	return err
}
