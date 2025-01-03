package handlers

import (
    "encoding/json"
    "net/http"
    "apiProject/api"
    "apiProject/internal/tools"
    log "github.com/sirupsen/logrus"
	"fmt"
	"errors"
)



func CreateAccount(w http.ResponseWriter, r *http.Request) {
    var params api.AddUserParams
    err := json.NewDecoder(r.Body).Decode(&params)
    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }

    var database *tools.DatabaseInterface
    database, err = tools.NewDatabase()
    if err != nil {
        api.InternalErrorHandler(w)
        return
    }
    tokenDetails, e := (*database).CreateUser(params.Username, params.AuthToken, params.Coins)
	if tokenDetails == nil {
		log.Error(err)
		api.RequestErrorHandler(w, errors.New(fmt.Sprintf(e)))
		return
	}

    var response = api.UserCreateResponse{
        Username: tokenDetails.Username,
        Code:    http.StatusOK,
    }

    w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(w).Encode(response)
    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }
}