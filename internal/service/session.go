package service

import (
	"fmt"
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
		ID:           uuid.NewString(),
		CreatedAt:    time.Now(),
		HostId:       hostId,
		Ticket:       ticket,
		Participants: make(map[string]*models.Participant),
		Votes:        make(map[string]*models.Vote),
		Revealed:     false,
	}

	if err := s.store.Create(session); err != nil {
		return nil, err
	}
	return session, nil
}

func (s *SessionService) GetSession(id string) (*models.Session, error) {
	return s.store.Get(id)
}

func (s *SessionService) AddParticipant(sessionID, name string) (*models.Participant, error) {
	session, err := s.GetSession(sessionID)
	if err != nil {
		return nil, err
	}

	participant := &models.Participant{
		ID:   uuid.NewString(),
		Name: name,
	}

	if session.Participants == nil {
		session.Participants = make(map[string]*models.Participant)
	}
	session.Participants[participant.ID] = participant

	return participant, nil
}

func (s *SessionService) CastVote(sessionID, participantID, value string) (*models.Vote, error) {
	session, err := s.GetSession(sessionID)
	if err != nil {
		return nil, err
	}

	if _, ok := session.Participants[participantID]; !ok {
		return nil, fmt.Errorf("participant %s not found in session", participantID)
	}

	vote := &models.Vote{
		ParticipantID: participantID,
		Value:         value,
	}

	if session.Votes == nil {
		session.Votes = make(map[string]*models.Vote)
	}

	session.Votes[participantID] = vote

	return vote, nil
}

func (s *SessionService) RevealVotes(sessionID string) (*models.Session, error) {
	session, err := s.GetSession(sessionID)

	if err != nil {
		return nil, err
	}

	session.Revealed = true

	return session, nil
}
