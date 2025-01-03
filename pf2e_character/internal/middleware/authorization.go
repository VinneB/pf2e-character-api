package middleware

import (
	"errors"
	"net/http"

	"pf2e-character-api/api"
	"pf2e-character-api/internal/tools"

	log "github.com/sirupsen/logrus"
)

var UnauthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			username string = r.URL.Query().Get("Username")
			token           = r.Header.Get("Authorization")
			err      error
		)

		if username == "" || token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		var database *tools.AuthDatabase
		database, err = tools.NewAuthDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails []tools.AuthDetails
		loginDetails, _ = (*database).GetAuthDetails()

		isFound := false
		var auth_index int = 0
		for auth_index = 0; auth_index < len(loginDetails); auth_index++ {
			if loginDetails[auth_index].Username == username {
				isFound = true
				break
			}
		}

		if !isFound || (token != loginDetails[auth_index].AuthToken) {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		// This calls the next piece of middleware on the request or the Handler func if there is no middleware left
		next.ServeHTTP(w, r)

	})

}
