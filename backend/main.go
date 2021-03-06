package main

import (
	"context"
	"fmt"
	"net/http"
	"snake_game/backend/handlers"
	"snake_game/backend/storage"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
)

const (
	// DatabaseSource reference to the database path
	DatabaseSource = "postgresql://admin@localhost:26257/snake?sslmode=disable"
)

func main() {
	port := ":3000"
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	dbConn, err := storage.DBConection(DatabaseSource)
	if err != nil {
		handlers.LogError(err)
	}
	defer dbConn.Close(context.Background())

	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)

	router.Route("/v1/players", func(r chi.Router) {
		r.Post("/register", handlers.PlayerRegister(dbConn))
		r.Post("/login", handlers.PlayerLogin(dbConn, tokenAuth))
		r.Get("/leaderboard", handlers.PlayerLeaderboard(dbConn))

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Put("/score", handlers.UpdatePlayerScore(dbConn))
		})
	})

	fmt.Println("Server running in localhost" + port)
	err = http.ListenAndServe(port, router)
	if err != nil {
		handlers.LogError(err)
	}
}
