package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"pf2e-character-api/api"
	"pf2e-character-api/internal/character"
	"pf2e-character-api/internal/tools"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	var params = api.CharacterParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	var char_database *tools.CharacterDatabase

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	char_database, err = tools.NewCharacterDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	err = (*char_database).AddCharacter(params.Username, character.Character{
		Name: params.CharacterName,
	})

	if err != nil {
		log.Error(err)
		if errors.Is(err, tools.ErrUserCannotBeFound) {
			api.RequestErrorHandler(w, err)
		} else {
			api.InternalErrorHandler(w)
		}
		return
	}

	var response api.GenericCodeResponse = api.GenericCodeResponse{
		Code: 200,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
