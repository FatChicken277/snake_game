package main

import (
	"context"
	"fmt"
	"net/http"
	"snake_game/backend/handlers"
	"snake_game/backend/storage"

	"github.com/go-chi/chi"
)

const (
	// DatabaseSource reference to the database path
	DatabaseSource = "postgresql://admin@localhost:26257/snake?sslmode=disable"
)

func main() {
	port := ":3000"
	router := chi.NewRouter()

	dbConn, err := storage.DBConection(DatabaseSource)
	if err != nil {
		handlers.LogError(err)
	}
	defer dbConn.Close(context.Background())

	router.Route("/v1/players", func(r chi.Router) {
		r.Post("/register", handlers.PlayerRegister(dbConn))
		r.Post("/login", handlers.PlayerLogin(dbConn))
		r.Get("/leaderboard", handlers.PlayerLeaderboard(dbConn))

		r.Route("/{player_id}", func(r chi.Router) {
			r.Put("/", handlers.UpdatePlayerScore(dbConn))
		})
	})

	fmt.Println("Server running in localhost" + port)
	err = http.ListenAndServe(port, router)
	if err != nil {
		handlers.LogError(err)
	}
}
