package memory

import (
	"errors"
	"sync"

	"github.com/MrDHiggins/scrumdcards-backend/internal/models"
)

type SessionMemoryStore struct {
	mu       sync.RWMutex
	sessions map[string]*models.Session
}

func NewSessionMemoryStore() *SessionMemoryStore {
	return &SessionMemoryStore{
		sessions: make(map[string]*models.Session),
	}
}

func (s *SessionMemoryStore) Create(session *models.Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.sessions[session.ID]; exists {
		return errors.New("session already exists")
	}

	s.sessions[session.ID] = session
	return nil
}

func (s *SessionMemoryStore) Get(id string) (*models.Session, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if session, exists := s.sessions[id]; exists {
		return session, nil
	}

	return nil, errors.New("session not found")
}
