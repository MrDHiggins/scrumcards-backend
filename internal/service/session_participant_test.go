package service

import (
	"testing"

	"github.com/MrDHiggins/planning-poker-backend/internal/store/memory"
)

func TestAddParticipant(t *testing.T) {
	store := memory.NewSessionMemoryStore()
	sessionService := NewSessionService(store)

	// Create a session first
	session, err := sessionService.CreateSession("host1", "JIRA-123")
	if err != nil {
		t.Fatalf("expected no error creating session, got %v", err)
	}

	// Add participant
	participant, err := sessionService.AddParticipant(session.ID, "Fabio")
	if err != nil {
		t.Fatalf("expected no error adding participant, got %v", err)
	}

	if participant.Name != "Fabio" {
		t.Errorf("expected participant name 'Fabio', got '%v'", participant.Name)
	}

	// Verify participant stored in session
	storedSession, _ := sessionService.GetSession(session.ID)
	if len(storedSession.Participants) != 1 {
		t.Errorf("expected 1 participant, got %d", len(storedSession.Participants))
	}

	// Add second participant
	_, err = sessionService.AddParticipant(session.ID, "Bob")
	if err != nil {
		t.Fatalf("expected no error adding second participant, got %v", err)
	}

	storedSession, _ = sessionService.GetSession(session.ID)
	if len(storedSession.Participants) != 2 {
		t.Errorf("expected 2 participants, got %d", len(storedSession.Participants))
	}
}

func TestAddParticipantInvalidSession(t *testing.T) {
	store := memory.NewSessionMemoryStore()
	sessionService := NewSessionService(store)

	_, err := sessionService.AddParticipant("nonexistent-id", "Alice")
	if err == nil {
		t.Errorf("expected error when adding participant to non-existent session, got nil")
	}
}
