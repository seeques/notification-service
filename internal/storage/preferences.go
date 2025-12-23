package storage

import (
	"context"
)

func (s *PostgresStorage) SetPreference(ctx context.Context, userID string, channel string, enabled bool) (*Preference, error) {
	sql := `INSERT INTO preferences (user_id, channel, enabled)
	VALUES ($1, $2, $3)
	ON CONFLICT(user_id, channel)
	DO UPDATE SET
	enabled = EXCLUDED.enabled,
	updated_at = NOW()
	RETURNING id, user_id, channel, enabled, created_at, updated_at`

	pref := Preference{}

	err := s.pool.QueryRow(ctx, sql, userID, channel, enabled).Scan(
		&pref.ID, 
		&pref.UserID, 
		&pref.Channel, 
		&pref.Enabled, 
		&pref.CreatedAt,
		&pref.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &pref, nil
}

func (s *PostgresStorage) GetPreference(ctx context.Context, ID int, userID string) (map[string]bool, error) {
	query := `SELECT id, user_id, channel, enabled, created_at, updated_at
	FROM preferences 
	WHERE id = $1 AND user_id = $2`

	pref := Preference{}

	err := s.pool.QueryRow(ctx, query, ID, userID).Scan(
		&pref.ID, 
		&pref.UserID, 
		&pref.Channel, 
		&pref.Enabled, 
		&pref.CreatedAt,
		&pref.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	m := make(map[string]bool)
	m[pref.Channel] = pref.Enabled
	
	return m, nil
}