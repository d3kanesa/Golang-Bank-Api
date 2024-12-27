package handlers

import (
    "encoding/json"
    "net/http"
    "apiProject/api"
    "apiProject/internal/tools"
    log "github.com/sirupsen/logrus"
	"fmt"
)

func TransferCoins(w http.ResponseWriter, r *http.Request) {
    var params api.TransferCoinParams
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
	fmt.Println(params)
    var tokenDetails *tools.CoinDetails = (*database).TransferCoins(params.Username, params.Receiver, params.AddAmount)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
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