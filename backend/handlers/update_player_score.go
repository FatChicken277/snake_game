package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"snake_game/backend/models"
	"snake_game/backend/storage"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx"
)

func getAndVerifyUpdateScoreParams(dbConn *pgx.Conn, r *http.Request, player *models.PlayerModel) error {
	playerID, err := strconv.Atoi(chi.URLParam(r, "player_id"))
	if err != nil {
		return ErrInvalidPlayer
	}

	player.PlayerID = playerID

	query := "SELECT username FROM players WHERE player_id = $1;"
	row := dbConn.QueryRow(context.Background(), query, player.PlayerID)
	err = row.Scan(player.Username)
	if err == pgx.ErrNoRows {
		return ErrInvalidPlayer
	}

	err = json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		return ErrInvalidJSON
	}

	if player.MaxScore == 0 {
		return ErrMissingMaxScore
	}

	return nil
}

// UpdatePlayerScore is used to update the score of a player
func UpdatePlayerScore(dbConn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var player models.PlayerModel

		err := getAndVerifyUpdateScoreParams(dbConn, r, &player)
		if err != nil {
			LogError(NewErrorResponse(ErrorResponseBadRequest, w, err.Error()))
			return
		}

		err = storage.UpdatePlayerScore(dbConn, &player)
		if err != nil {
			LogError(NewErrorResponse(ErrorResponseBadRequest, w, err.Error()))
			return
		}

		resp := ResponseOK
		resp.Message = "player score was successfuly updated"

		LogError(NewResponse(resp, w))
	}
}
