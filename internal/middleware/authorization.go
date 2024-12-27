package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"bytes"
	"apiProject/api"
	"apiProject/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New(fmt.Sprintf("Invalid username or token."))

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {	
		var username string
		var token = r.Header.Get("Authorization")
		var err error

		if r.Method == http.MethodGet {
			username = r.URL.Query().Get("username")
		} else{

			bodyBytes, err := ioutil.ReadAll(r.Body)
            if err != nil {
                log.Error(err)
                api.InternalErrorHandler(w)
                return
            }

			var params map[string]interface{}
            err = json.Unmarshal(bodyBytes, &params)
            if err != nil {
                log.Error(err)
                api.InternalErrorHandler(w)
                return
            }
            username, _ = params["username"].(string)

			r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}
		if username == "" {
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)

	})
}