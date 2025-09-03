package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/MrDHiggins/planning-poker-backend/internal/service"
	"github.com/MrDHiggins/planning-poker-backend/internal/store/memory"

	"testing"

	"github.com/go-chi/chi/v5"
)

func setupTestRouter() (*chi.Mux, *service.SessionService) {
	store := memory.NewSessionMemoryStore()
	sessionService := service.NewSessionService(store)
	handler := NewSessionHandler(sessionService)

	r := chi.NewRouter()
	handler.RegisterRoutes(r)

	return r, sessionService
}

func TestJoinSessionAPI(t *testing.T) {
	r, svc := setupTestRouter()

	// Create a session first
	session, _ := svc.CreateSession("host1", "JIRA-123")

	// Create request body
	body := []byte(`{"name":"Alice"}`)

	req := httptest.NewRequest("POST", "/sessions/"+session.ID+"/participants", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp["name"] != "Alice" {
		t.Errorf("expected participant name 'Alice', got '%v'", resp["name"])
	}
	if resp["id"] == "" {
		t.Error("expected participant ID, got empty string")
	}
}

func TestJoinSessionAPINotFound(t *testing.T) {
	r, _ := setupTestRouter()

	body := []byte(`{"name":"Alice"}`)
	req := httptest.NewRequest("POST", "/sessions/nonexistent-id/participants", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", w.Code)
	}
}

func TestJoinSessionAPIInvalidBody(t *testing.T) {
	r, svc := setupTestRouter()

	// Create a session
	session, _ := svc.CreateSession("host1", "JIRA-123")

	// Missing name
	body := []byte(`{}`)
	req := httptest.NewRequest("POST", "/sessions/"+session.ID+"/participants", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400 for invalid body, got %d", w.Code)
	}
}
