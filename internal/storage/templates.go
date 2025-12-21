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

func (s *PostgresStorage) GetTemp(ctx context.Context, name string) (*Template, error) {
	query := `SELECT id, name, subject, body, created_at, updated_at
	FROM templates
	WHERE name = $1`

	temp := Template{}

	err := s.pool.QueryRow(ctx, query, name).Scan(
		&temp.ID,
		&temp.Name,
		&temp.Subject,
		&temp.Body,
		&temp.CreatedAt,
		&temp.UpdatedAt,
	)
	return &temp, err
}