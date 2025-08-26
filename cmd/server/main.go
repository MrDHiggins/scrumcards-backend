package main

import (
	"log"
	"net/http"

	"github.com/MrDHiggins/planning-poker-backend/internal/api"
	"github.com/MrDHiggins/planning-poker-backend/internal/service"
	"github.com/MrDHiggins/planning-poker-backend/internal/store/memory"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	sessionStore := memory.NewSessionMemoryStore()
	sessionService := service.NewSessionService(sessionStore)
	sessionHandler := api.NewSessionHandler(sessionService)

	sessionHandler.RegisterRoutes(r)

	log.Println("Planning Poker backend started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
