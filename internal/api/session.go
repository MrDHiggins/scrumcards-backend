package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/MrDHiggins/planning-poker-backend/internal/models"
	"github.com/MrDHiggins/planning-poker-backend/internal/service"
	"github.com/MrDHiggins/planning-poker-backend/internal/utils"

	"github.com/go-chi/chi/v5"
)

type SessionHandler struct {
	service *service.SessionService
}

func NewSessionHandler(s *service.SessionService) *SessionHandler {
	return &SessionHandler{service: s}
}

func (h *SessionHandler) RegisterRoutes(r chi.Router) {
	r.Post("/sessions", h.CreateSession)
	r.Get("/sessions/{id}", h.GetSession)
	r.Post("/sessions/{id}/participants", h.JoinSession)
	r.Post("/sessions/{id}/votes", h.CastVote)
	r.Post("/sessions/{id}/reveal", h.RevealVotes)
}

func (h *SessionHandler) CreateSession(w http.ResponseWriter, r *http.Request) {
	var req struct {
		HostID string `json:"host_id"`
		Ticket string `json:"ticket"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	session, err := h.service.CreateSession(req.HostID, req.Ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("[SESSION CREATED] ID=%s | HostID=%s | Ticket=%s | CreatedAt=%s\n",
		session.ID, session.HostId, session.Ticket, session.CreatedAt.Format(time.RFC3339))
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(session)
}

func (h *SessionHandler) GetSession(w http.ResponseWriter, r *http.Request) {
	sessionID := chi.URLParam(r, "id")
	session, err := h.service.GetSession(sessionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := models.SessionResponse{
		ID:        session.ID,
		CreatedAt: session.CreatedAt,
		HostID:    session.HostId,
		Ticket:    session.Ticket,
	}

	for _, p := range session.Participants {
		resp.Participants = append(resp.Participants, p)
	}

	for _, v := range session.Votes {
		resp.Votes = append(resp.Votes, v)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *SessionHandler) JoinSession(w http.ResponseWriter, r *http.Request) {
	sessionID := chi.URLParam(r, "id")

	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	participant, err := h.service.AddParticipant(sessionID, req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(participant)
}

func (h *SessionHandler) CastVote(w http.ResponseWriter, r *http.Request) {
	sessionID := chi.URLParam(r, "id")

	var req struct {
		ParticipantID string `json:"participant_id"`
		Value         string `json:"value"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.ParticipantID == "" || req.Value == "" {
		http.Error(w, "invalid participant vote request body", http.StatusBadRequest)
		return
	}

	vote, err := h.service.CastVote(sessionID, req.ParticipantID, req.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vote)
}

func (h *SessionHandler) RevealVotes(w http.ResponseWriter, r *http.Request) {
	sessionID := chi.URLParam(r, "id")
	session, err := h.service.RevealVotes(sessionID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := models.SessionResponse{
		ID:        session.ID,
		CreatedAt: session.CreatedAt,
		HostID:    session.HostId,
		Ticket:    session.Ticket,
		Revealed:  session.Revealed,
	}

	for _, p := range session.Participants {
		resp.Participants = append(resp.Participants, p)
	}

	// May need to create a util for the below when
	// once the websocket has been implemented.
	var enrichedVotes []models.VoteResponse
	for _, v := range session.Votes {
		if participant, ok := session.Participants[v.ParticipantID]; ok {
			enrichedVotes = append(enrichedVotes, models.VoteResponse{
				ParticipantID:   v.ParticipantID,
				ParticipantName: participant.Name,
				Value:           v.Value,
			})
		}
	}

	respData := map[string]any{
		"session":     resp,
		"votes":       enrichedVotes,
		"VoteAverage": utils.CalculateVoteAverage(session.Votes),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respData)
}
