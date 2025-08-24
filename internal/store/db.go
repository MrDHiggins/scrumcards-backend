package store

import "github.com/MrDHiggins/planning-poker-backend/internal/models"

// SessionStore defines the storage operations for sessions
type SessionStore interface {
	Create(session *models.Session) error
	Get(id string) (*models.Session, error)
}
