package authentication

import (
	"context"
	"snake_game/backend/models"
	"snake_game/backend/storage"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayerAuthVerification(t *testing.T) {
	c := require.New(t)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	query := "INSERT INTO players (username, password) VALUES ($1, $2);"
	_, err = dbConn.Exec(context.Background(), query, "example", "$2a$10$l8.rKqZmp9fgfX5KqOqy3uFO1UrYPrxZO5SFrzU3ykSYVOObz6A0u")
	c.NoError(err)

	newPlayer := models.PlayerModel{
		Username: "example",
		Password: "password",
	}

	err = PlayerAuthVerification(dbConn, &newPlayer)
	c.NoError(err)

	query = "DELETE FROM players WHERE username = 'example';"
	_, err = dbConn.Exec(context.Background(), query)
	c.NoError(err)
}

func TestPlayerAuthVerificationInvalidPlayer(t *testing.T) {
	c := require.New(t)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	query := "INSERT INTO players (username, password) VALUES ($1, $2);"
	_, err = dbConn.Exec(context.Background(), query, "example", "$2a$10$l8.rKqZmp9fgfX5KqOqy3uFO1UrYPrxZO5SFrzU3ykSYVOObz6A0u")
	c.NoError(err)

	newPlayer := models.PlayerModel{
		Username: "invalid",
		Password: "password",
	}

	err = PlayerAuthVerification(dbConn, &newPlayer)
	c.Error(err)
	c.Equal(ErrInvalidPlayer, err)

	query = "DELETE FROM players WHERE username = 'example';"
	_, err = dbConn.Exec(context.Background(), query)
	c.NoError(err)
}

func TestPlayerAuthVerificationInvalidPassword(t *testing.T) {
	c := require.New(t)

	dbConn, err := storage.DBConection("postgresql://admin@localhost:26257/snake_test?sslmode=disable")
	c.NoError(err)
	defer dbConn.Close(context.Background())

	query := "INSERT INTO players (username, password) VALUES ($1, $2);"
	_, err = dbConn.Exec(context.Background(), query, "example", "$2a$10$l8.rKqZmp9fgfX5KqOqy3uFO1UrYPrxZO5SFrzU3ykSYVOObz6A0u")
	c.NoError(err)

	newPlayer := models.PlayerModel{
		Username: "example",
		Password: "invalid",
	}

	err = PlayerAuthVerification(dbConn, &newPlayer)
	c.Error(err)
	c.Equal(ErrInvalidPassword, err)

	query = "DELETE FROM players WHERE username = 'example';"
	_, err = dbConn.Exec(context.Background(), query)
	c.NoError(err)
}
