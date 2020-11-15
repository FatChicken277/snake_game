package handlers

import (
	"encoding/json"
	"net/http"
	"snake_game/backend/authentication"
	"snake_game/backend/models"

	"github.com/jackc/pgx"
)

func getAndVerifyLoginParams(dbConn *pgx.Conn, r *http.Request, player *models.PlayerModel) error {
	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		return ErrInvalidJSON
	}

	if player.Username == "" {
		return ErrMissingUsername
	}

	if player.Password == "" {
		return ErrMissingPassword
	}

	return nil
}

// PlayerLogin is in charge of a player login
func PlayerLogin(dbConn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var player models.PlayerModel

		err := getAndVerifyLoginParams(dbConn, r, &player)
		if err != nil {
			LogError(NewErrorResponse(ErrorResponseBadRequest, w, err.Error()))
			return
		}

		err = authentication.PlayerAuthVerification(dbConn, &player)
		if err != nil {
			LogError(NewErrorResponse(ErrorResponseUnauthorized, w, err.Error()))
			return
		}

		player.Password, player.PasswordHash = "", ""

		resp := ResponseOK
		resp.Data = player

		LogError(NewResponse(resp, w))
	}
}
