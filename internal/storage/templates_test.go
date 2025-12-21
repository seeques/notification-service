package storage

import (
	"testing"
	"context"
)

func TestCreateTemp(t *testing.T) {
	s := setupTestStorage(t)
	ctx := context.Background()

	template := &Template{
		Name: "test",
		Subject: "test_subject",
		Body: "test_body",
	}

	err := s.CreateTemp(ctx, template)
	if err != nil {
		t.Fatalf("failed to create a template: %v", err)
	}

	if template.ID == 0 {
		t.Error("expected ID to be set")
	}

	if template.CreatedAt.IsZero() || template.UpdatedAt.IsZero() {
		t.Error("expected timestamp to be set")
	}
}