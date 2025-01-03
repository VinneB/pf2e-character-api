package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"pf2e-character-api/api"
	"pf2e-character-api/internal/tools"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

var ErrTooShort = errors.New("username and password must be greater than 8 characters")
var ErrAlreadyExists error = errors.New("user with this username already exists")

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var params = api.TokenParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	var auth_database *tools.AuthDatabase
	var auths []tools.AuthDetails

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Ensure Proper length
	if len(params.Username) <= 8 || len(params.Password) <= 8 {
		log.Error(ErrTooShort)
		api.RequestErrorHandler(w, ErrTooShort)
		return
	}

	auth_database, err = tools.NewAuthDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	auths, err = (*auth_database).GetAuthDetails()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	// Check if user already exists
	for i := 0; i < len(auths); i++ {
		if (auths[i].Username == params.Username) && (auths[i].Password == params.Password) {
			log.Error(ErrAlreadyExists)
			api.RequestErrorHandler(w, ErrAlreadyExists)
			return
		}
	}

	var authDetail tools.AuthDetails = tools.AuthDetails{
		Username:  params.Username,
		Password:  params.Password,
		AuthToken: tools.RandSeq(api.TokenLength),
	}

	auth_database, err = tools.NewAuthDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	err = (*auth_database).AddAuthDetail(authDetail)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
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
