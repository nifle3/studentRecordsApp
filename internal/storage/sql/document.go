package sql

import (
	"context"

	"github.com/google/uuid"

	"studentRecordsApp/internal/casts"
	"studentRecordsApp/internal/service/entites"
	"studentRecordsApp/internal/storage/sql/sqlEntities"
)

func (s *Storage) GetStudentsDocumentById(id string, userId string, ctx context.Context) (entities.Document, error) {
	uuId, err := uuid.Parse(id)
	if err != nil {
		return entities.Document{}, err
	}

	uuUserId, err := uuid.Parse(userId)
	if err != nil {
		return entities.Document{}, err
	}

	var result sqlEntities.StudentsDocument

	err = s.db.GetContext(ctx, &result,
		`SELECT * FROM StudentsDocuments WHERE id = $1 AND student_id = $2;`, uuUserId, uuId)

	if err != nil {
		return entities.Document{}, err
	}

	return casts.DocumentSqlToEntite(result, ctx), nil
}

func (s *Storage) GetStudentsDocumentsForUser(userId string, ctx context.Context) ([]entities.Document, error) {
	uuUserId, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	sqlResults := make([]sqlEntities.StudentsDocument, 0)
	err = s.db.SelectContext(ctx, &sqlResults, `SELECT * FROM StudentsDocuments WHERE student_id = $1;`, uuUserId)
	if err != nil {
		return nil, err
	}

	results := make([]entities.Document, 0, len(sqlResults))
	for _, result := range sqlResults {
		results = append(results, casts.DocumentSqlToEntite(result, ctx))
	}

	return results, nil
}

func (s *Storage) AddStudentsDocument(document entities.Document, ctx context.Context) error {
	sqlDocument, err := casts.DocumentEntiteToSqlWithoutId(document, ctx)
	if err != nil {
		return err
	}

	sqlDocument.Id = uuid.New()
	_, err = s.db.ExecContext(ctx,
		`INSERT INTO StudentsDocuments (id, student_id, document_name, document_type, document_link_s3, created_at) 
				VALUES ($1,$2,$3,$4,$5, $6);`,
		sqlDocument.Id, sqlDocument.StudentId, sqlDocument.Name, sqlDocument.Type, sqlDocument.Link, sqlDocument.CreatedAt)

	return err
}

func (s *Storage) DeleteStudentsDocument(id string, userId string, ctx context.Context) error {
	uuId, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	uuUserId, err := uuid.Parse(userId)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx,
		`DELETE FROM StudentsDocuments WHERE id = $1 AND student_id = $2;`, uuId, uuUserId)

	return err
}

func (s *Storage) UpdateStudentsDocument(document entities.Document, ctx context.Context) error {
	sqlDocument, err := casts.DocumentEntiteToSql(document, ctx)
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx,
		`UPDATE StudentsDocuments SET document_name =$1, document_type =$2, document_link_s3 =$3 WHERE id =$4 AND student_id = $5;`,
		sqlDocument.Name, sqlDocument.Type, sqlDocument.Link, sqlDocument.Id, sqlDocument.StudentId)

	return err
}
