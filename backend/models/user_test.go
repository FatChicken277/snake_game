package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserVerifyParams(t *testing.T) {
	c := require.New(t)

	newUser := UserModel{
		UserID:   12,
		Username: "FatChicken",
		Score:    32,
	}

	err := newUser.VerifyParams()
	c.NoError(err)

	newUser.Username = ""
	err = newUser.VerifyParams()
	c.Error(err)
	c.Equal(ErrMissingUsername, err)
}

func TestUserToResponse(t *testing.T) {
	c := require.New(t)

	var newUser1 UserModel
	newUser2 := UserModel{
		UserID:   12,
		Username: "FatChicken",
		Score:    32,
	}

	resp, err := newUser2.ToResponse()
	c.NoError(err)

	err = json.Unmarshal(resp, &newUser1)
	c.NoError(err)
	c.Equal(newUser1, newUser2)
}
