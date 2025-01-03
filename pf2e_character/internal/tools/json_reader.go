package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"pf2e-character-api/internal/character"
)

var ErrUserCannotBeFound = errors.New("user cannot be found")

type JsonFileInteractor struct {
	file_name string
}

func (d *JsonFileInteractor) SetupAuthDatabase() error {
	return nil
}

func (d *JsonFileInteractor) SetupCharacterDatabase() error {
	return nil
}

func (d *JsonFileInteractor) GetAuthDetails() ([]AuthDetails, error) {
	var file *os.File

	dir, _ := os.Getwd()
	fmt.Println(dir)

	if _, err := os.Stat(d.file_name); errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	file, err := os.Open(d.file_name)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := io.ReadAll(file)

	var auth_details AuthDetailsArray // read our opened jsonFile as a byte array.

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &auth_details)

	return auth_details.AuthArray, nil

}

// Assumes authorization and doesn't check credentials
func (d *JsonFileInteractor) GetCharacters(username string) ([]character.Character, error) {
	var file *os.File

	dir, _ := os.Getwd()
	fmt.Println(dir)

	if _, err := os.Stat(d.file_name); errors.Is(err, os.ErrNotExist) {
		return nil, err
	}
	file, err := os.Open(d.file_name)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := io.ReadAll(file)

	var details_arr AuthDetailsArray // read our opened jsonFile as a byte array.

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &details_arr)

	var selectedAuthDetails AuthDetails
	isFound := false
	for i := 0; i < len(details_arr.AuthArray); i++ {
		if details_arr.AuthArray[i].Username == username {
			selectedAuthDetails = details_arr.AuthArray[i]
			isFound = true
		}
	}

	if !isFound {
		return nil, ErrUserCannotBeFound
	}

	return selectedAuthDetails.CharacterArray, nil
}

func (d *JsonFileInteractor) AddCharacter(username string, character character.Character) error {
	var file *os.File

	if _, err := os.Stat(d.file_name); errors.Is(err, os.ErrNotExist) {
		return err
	}

	file, err := os.Open(d.file_name)

	if err != nil {
		return err
	}
	defer file.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := io.ReadAll(file)

	var details_arr AuthDetailsArray // read our opened jsonFile as a byte array.

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &details_arr)

	isFound := false
	for i := 0; i < len(details_arr.AuthArray); i++ {
		if details_arr.AuthArray[i].Username == username {
			details_arr.AuthArray[i].CharacterArray = append(details_arr.AuthArray[i].CharacterArray, character)
			isFound = true
		}
	}

	if !isFound {
		return ErrUserCannotBeFound
	}

	jsonString, _ := json.Marshal(details_arr)
	os.WriteFile(d.file_name, jsonString, os.ModePerm)

	return nil
}

func (d *JsonFileInteractor) DeleteCharacter(username string, characterName string) (character.Character, error) {
	var file *os.File

	if _, err := os.Stat(d.file_name); errors.Is(err, os.ErrNotExist) {
		return character.Character{}, err
	}

	file, err := os.Open(d.file_name)

	if err != nil {
		return character.Character{}, err
	}
	defer file.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := io.ReadAll(file)

	var details_arr AuthDetailsArray // read our opened jsonFile as a byte array.

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &details_arr)

	isFound := false
	var details_index uint16
	for details_index = 0; int(details_index) < len(details_arr.AuthArray); details_index++ {
		if details_arr.AuthArray[details_index].Username == username {
			isFound = true
			break
		}
	}

	if !isFound {
		return character.Character{}, ErrUserCannotBeFound
	}

	var deleted_character character.Character
	for i := 0; i < len(details_arr.AuthArray[details_index].CharacterArray); i++ {
		if details_arr.AuthArray[details_index].CharacterArray[i].Name == characterName {
			// Remove element
			deleted_character = details_arr.AuthArray[details_index].CharacterArray[i]
			details_arr.AuthArray[details_index].CharacterArray[i] = details_arr.AuthArray[details_index].CharacterArray[len(details_arr.AuthArray[details_index].CharacterArray)-1]
			details_arr.AuthArray[details_index].CharacterArray = details_arr.AuthArray[details_index].CharacterArray[:len(details_arr.AuthArray[details_index].CharacterArray)-1]
		}
	}

	jsonString, _ := json.Marshal(details_arr)
	os.WriteFile(d.file_name, jsonString, os.ModePerm)

	return deleted_character, nil
}

func (d *JsonFileInteractor) AddAuthDetail(auth AuthDetails) error {

	var err error
	var authDetailsArray AuthDetailsArray
	authDetailsArray.AuthArray, err = d.GetAuthDetails()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	} else if err != nil && errors.Is(err, os.ErrNotExist) {
		authDetailsArray = AuthDetailsArray{
			AuthArray: []AuthDetails{},
		}
	}

	authDetailsArray.AuthArray = append(authDetailsArray.AuthArray, auth)

	jsonString, _ := json.Marshal(authDetailsArray)
	os.WriteFile(d.file_name, jsonString, os.ModePerm)

	/* 	file, err := os.OpenFile(d.file_name, os.O_CREATE, os.ModePerm)
	   	if err != nil {

	   	}

	   	defer file.Close()

	   	encoder := json.NewEncoder(file)
	   	encoder.Encode(authDetailsArray) */

	return nil
}
