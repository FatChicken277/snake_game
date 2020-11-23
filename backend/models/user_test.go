package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayerToResponse(t *testing.T) {
	c := require.New(t)

	newPlayer := PlayerModel{
		PlayerID:        1,
		Username:        "example",
		Password:        "password",
		PasswordConfirm: "password",
		MaxScore:        23,
	}

	var newPlayer2 PlayerModel

	jsonPlayer, err := newPlayer.ToResponse()
	c.NoError(err)

	err = json.Unmarshal(jsonPlayer, &newPlayer2)
	c.NoError(err)

	c.Equal(newPlayer, newPlayer2)
}
