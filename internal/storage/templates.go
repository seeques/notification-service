package storage

import (
	"context"
)

func (s *PostgresStorage) CreateTemp(ctx context.Context, temp *Template) error {
	sql := `INSERT INTO templates (name, subject, body)
	VALUES ($1, $2, $3)
	RETURNING id, created_at, updated_at`

	err := s.pool.QueryRow(ctx, sql, temp.Name, temp.Subject, temp.Body).Scan(
		&temp.ID, 
		&temp.CreatedAt, 
		&temp.UpdatedAt,
	)

	return err
}
