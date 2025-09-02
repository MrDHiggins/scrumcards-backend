package memory

import (
	"testing"
	"time"

	"github.com/MrDHiggins/planning-poker-backend/internal/models"
)

func TestSessionManagementStore(t *testing.T) {
	store := NewSessionMemoryStore()
	testSession := &models.Session{ID: "123", CreatedAt: time.Now(), HostId: "anon1"}

	// Test create session successfully
	if err := store.Create(testSession); err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Test duplicate session ID
	if err := store.Create(testSession); err == nil {
		t.Errorf("expected error when creating duplicate session ID, but got no error")
	}

	// Test get existing session
	fetchedSession, err := store.Get("123")
	if err != nil {
		t.Errorf("expected to fetch session successfully, got error: %v", err)
	}

	if fetchedSession.HostId != "anon1" {
		t.Errorf("expected hostID 'host1', got '%v'", fetchedSession.HostId)
	}

	// Test invalid session
	_, err = store.Get("999")
	if err == nil {
		t.Errorf("expected error fetching non-existent session, but got none")
	}
}
