package service

import (
	"time"

	"github.com/MrDHiggins/planning-poker-backend/internal/models"
	"github.com/MrDHiggins/planning-poker-backend/internal/store"
	"github.com/google/uuid"
)

type SessionService struct {
	store store.SessionStore
}

func NewSessionService(store store.SessionStore) *SessionService {
	return &SessionService{store: store}
}

func (s *SessionService) CreateSession(hostId string, ticket string) (*models.Session, error) {
	session := &models.Session{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		HostId:    hostId,
		Ticket:    ticket,
	}

	if err := s.store.Create(session); err != nil {
		return nil, err
	}
	return session, nil
}

func (s *SessionService) GetSession(id string) (*models.Session, error) {
	return s.store.Get(id)
}
