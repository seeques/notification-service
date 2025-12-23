package storage

import (
	"testing"
	"context"
)

func TestCreatePreference(t *testing.T) {
	s := setupTestStorage(t)
	ctx := context.Background()

	userID := "abcUser"
	channel := "email"
	enabled := true

	pref, err := s.SetPreference(ctx, userID, channel, enabled)
	if err != nil {
		t.Fatalf("failed to set a preference: %v", err)
	}

	if pref.UserID != userID {
		t.Error("userID mismatch")
	}

	if pref.Enabled != true {
		t.Error("channel should be enabled")
	}
}

func TestUpdatePreference(t *testing.T) {
	s := setupTestStorage(t)
	ctx := context.Background()

	userID := "abcUser"
	channel := "email"
	enabled := true

	_, err := s.SetPreference(ctx, userID, channel, enabled)
	if err != nil {
		t.Fatalf("failed to set a preference: %v", err)
	}

	enabled = false
	pref2, err := s.SetPreference(ctx, userID, channel, enabled)
	if err != nil {
		t.Fatalf("failed to set a preference: %v", err)
	}

	if pref2.Enabled != false {
		t.Error("channel should be disabled")
	}

}

func TestGetPreference(t *testing.T) {
	s := setupTestStorage(t)
	ctx := context.Background()

	userID := "abcUser"
	channel := "email"
	enabled := true

	pref, err := s.SetPreference(ctx, userID, channel, enabled)
	if err != nil {
		t.Fatalf("failed to set a preference: %v", err)
	}

	m, err := s.GetPreference(ctx, pref.ID, pref.UserID)
	if err != nil {
		t.Fatalf("failed to get a preference: %v", err)
	}

	if m[channel] != true {
		t.Error("channel should be enabled")
	}
}