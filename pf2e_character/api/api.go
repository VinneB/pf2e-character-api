package api

import (
	"encoding/json"
	"net/http"
	"pf2e-character-api/internal/character"
)

// API Endpoint Parameters
// These are the paramaters passed in the url of the api call

// Token Parameters
type TokenParams struct {
	Username string
	Password string
}

type CharacterParams struct {
	Username      string
	CharacterName string
}

// API Response
// This is what is returned back to the client

// Response for CreateUser
type GenericCodeResponse struct {
	Code int
}

// Response for GetToken
type GetTokenResponse struct {
	Code  int
	Token string
}

type GetCharacterResponse struct {
	Code      int
	Character character.Character
}

type Error struct {
	Code int

	Message string
}

var TokenLength uint8 = 16

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)
