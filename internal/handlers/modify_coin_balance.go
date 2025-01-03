package handlers

import (
    "encoding/json"
    "net/http"
    "apiProject/api"
    "apiProject/internal/tools"
    log "github.com/sirupsen/logrus"
    "errors"
    "fmt"
)

func DepositCoins(w http.ResponseWriter, r *http.Request) {
    var params api.ModifyCoinParams
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
	
    tokenDetails, e := (*database).ModifyUserCoins(params.Username, params.ModifyAmount, true)
	if tokenDetails == nil {
		log.Error(err)
		api.RequestErrorHandler(w, errors.New(fmt.Sprintf(e)))
		return
	}

    var response = api.CoinBalanceResponse{
        Balance: tokenDetails.Coins,
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

func WithdrawCoins(w http.ResponseWriter, r *http.Request) {
    var params api.ModifyCoinParams
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
	
    tokenDetails, e := (*database).ModifyUserCoins(params.Username, params.ModifyAmount, false)
	if tokenDetails == nil {
		log.Error(err)
		api.RequestErrorHandler(w, errors.New(fmt.Sprintf(e)))
		return
	}

    var response = api.CoinBalanceResponse{
        Balance: tokenDetails.Coins,
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