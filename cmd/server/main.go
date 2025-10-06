package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MrDHiggins/scrumdcards-backend/internal/api"
	"github.com/MrDHiggins/scrumdcards-backend/internal/service"
	"github.com/MrDHiggins/scrumdcards-backend/internal/store/memory"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	sessionStore := memory.NewSessionMemoryStore()
	sessionService := service.NewSessionService(sessionStore)
	sessionHandler := api.NewSessionHandler(sessionService)
	sessionHandler.RegisterRoutes(r)

	// Get port from environment (Cloud Run sets this)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // local dev
	}

	log.Printf("🚀 Planning Poker backend started on :%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
