package models

import (
	"encoding/json"
	"errors"
)

var (
	//ErrMissingUsername indicates that the username is missing
	ErrMissingUsername = errors.New("'username' cannot be empty")
)

// UserModel defines the model for users
type UserModel struct {
	Username string `json:"username"`
	Score    uint   `json:"score"`
}

// UserList defines the structure of a user list
type UserList struct {
	Users []UserModel `json:"users"`
}

// VerifyParams is used to verify the user parameters
func (u *UserModel) VerifyParams() error {
	if u.Username == "" {
		return ErrMissingUsername
	}

	return nil
}

// ToResponse is used to convert the user object into a JSON response
func (u *UserModel) ToResponse() ([]byte, error) {
	resp, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
