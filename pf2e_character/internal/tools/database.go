package tools

import (
	"pf2e-character-api/internal/character"

	log "github.com/sirupsen/logrus"
)

const auth_filename string = "./auth.json"
const character_filename string = "./auth.json"

type AuthDetails struct {
	AuthToken      string                `json:"auth_token"`
	Username       string                `json:"username"`
	Password       string                `json:"password"`
	CharacterArray []character.Character `json:"characters"`
}

type AuthDetailsArray struct {
	AuthArray []AuthDetails `json:"auth"`
}

type CharacterArray struct {
	Array []character.Character `json:"character"`
}

type AuthDatabase interface {
	GetAuthDetails() ([]AuthDetails, error)
	AddAuthDetail(auth AuthDetails) error
	SetupAuthDatabase() error
}

type CharacterDatabase interface {
	GetCharacters(username string) ([]character.Character, error)
	SetupCharacterDatabase() error
	AddCharacter(username string, character character.Character) error
	DeleteCharacter(username string, characterName string) (character.Character, error)
}

func NewAuthDatabase() (*AuthDatabase, error) {
	var database AuthDatabase = &JsonFileInteractor{
		file_name: auth_filename,
	}

	var err error = database.SetupAuthDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}

func NewCharacterDatabase() (*CharacterDatabase, error) {
	var database CharacterDatabase = &JsonFileInteractor{
		file_name: character_filename,
	}

	var err error = database.SetupCharacterDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
