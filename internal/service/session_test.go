package service

import (
	"testing"

	"github.com/MrDHiggins/scrumdcards-backend/internal/store/memory"
)

func TestCreateAndGetSession(t *testing.T) {
	store := memory.NewSessionMemoryStore()
	sessionService := NewSessionService(store)

	// Test Create session
	createdSession, err := sessionService.CreateSession("host1", "JIRA-123")
	if err != nil {
		t.Fatalf("expected no error creating session, got error: %v", err)
	}

	if createdSession.HostId != "host1" {
		t.Errorf("expected hostIdHostId 'host1', got '%v'", createdSession.HostId)
	}
	if createdSession.Ticket != "JIRA-123" {
		t.Errorf("expected ticket 'JIRA-123', got '%v'", createdSession.Ticket)
	}

	// Test Get session
	fetchedSession, err := sessionService.GetSession(createdSession.ID)
	if err != nil {
		t.Fatalf("expected to fetch session successfully, got error: %v", err)
	}
	if fetchedSession.ID != createdSession.ID {
		t.Errorf("expected session ID '%v', got '%v'", createdSession.ID, fetchedSession.ID)
	}
}
