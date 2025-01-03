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

var ErrCharNotFound error = errors.New("character not found")

func GetCharacter(w http.ResponseWriter, r *http.Request) {
	var params = api.CharacterParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	var database *tools.CharacterDatabase

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	database, err = tools.NewCharacterDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var characters []character.Character
	var character_index uint16
	characters, err = (*database).GetCharacters(params.Username)
	if err != nil {
		log.Error(err)
		if errors.Is(err, tools.ErrUserCannotBeFound) {
			api.RequestErrorHandler(w, err)
		} else {
			api.InternalErrorHandler(w)
		}
		return
	}

	isFound := false
	for character_index = 0; int(character_index) < len(characters); character_index++ {
		if characters[character_index].Name == params.CharacterName {
			isFound = true
			break
		}
	}

	if !isFound {
		log.Error(ErrCharNotFound)
		api.RequestErrorHandler(w, ErrCharNotFound)
		return
	}

	var response api.GetCharacterResponse = api.GetCharacterResponse{
		Code:      200,
		Character: characters[character_index],
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
