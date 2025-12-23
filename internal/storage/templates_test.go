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

func TestGetTemp(t *testing.T) {
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

	retTemp, err := s.GetTemp(ctx, template.Name)
	if err != nil {
		t.Fatalf("failed to get the template: %v", err)
	}

	if retTemp.ID != template.ID {
		t.Error("ID does not match")
	}

	if retTemp.Subject != template.Subject {
		t.Error("Subject does not match")
	} 

	if retTemp.CreatedAt != template.CreatedAt {
		t.Error("creation time does not match")
	}
}

func TestListAllTemp(t *testing.T) {
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

	templates, err := s.ListAllTemp(ctx)
	if err != nil {
		t.Fatalf("failed to retrieve templates: %v", err)
	}

	temp := templates[0]

	if temp.ID != template.ID {
		t.Error("ID does not match")
	}
}

func TestUpdateTemp(t *testing.T) {
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

	name := "test_update"
	subject := "test_subject_update"
	body := "test_body_update"

	tempUpd, err := s.UpdateTemp(ctx, name, subject, body, template.ID)
	if err != nil {
		t.Fatalf("template update failed: %v", err)
	}

	if tempUpd.ID != template.ID {
		t.Error("ids mismatch")
	}

	if tempUpd.Name != "test_update" {
		t.Error("incorrect name after update")
	}
}

func TestDeleteTemp(t *testing.T) {
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

	err = s.DeleteTemp(ctx, template.ID)
	if err != nil {
		t.Fatalf("failed to delete a template: %v", err)
	}

	temp, _ := s.GetTemp(ctx, template.Name)
	if temp.Name != "" {
		t.Error("template should be deleted")
	}
}