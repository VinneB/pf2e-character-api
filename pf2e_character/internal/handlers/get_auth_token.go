package handlers

import (
	"net/http"

	"encoding/json"
	"errors"
	"pf2e-character-api/api"
	"pf2e-character-api/internal/tools"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

var ErrNotFound error = errors.New("username or password is incorrect")

func GetAuthToken(w http.ResponseWriter, r *http.Request) {
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

	auth_database, err = tools.NewAuthDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	auths, err = (*auth_database).GetAuthDetails()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var selectedAuthDetails tools.AuthDetails
	isFound := false
	for i := 0; i < len(auths); i++ {
		if (auths[i].Username == params.Username) && (auths[i].Password == params.Password) {
			selectedAuthDetails = auths[i]
			isFound = true
		}
	}

	if !isFound {
		log.Error(ErrNotFound)
		api.RequestErrorHandler(w, ErrNotFound)
		return
	}

	var response = api.GetTokenResponse{
		Token: selectedAuthDetails.AuthToken,
		Code:  200,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
