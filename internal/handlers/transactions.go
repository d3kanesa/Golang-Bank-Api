package handlers

import (
	"apiProject/api"
	"apiProject/internal/tools"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)



func TransactionDetails(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	err = decoder.Decode(&params, r.URL.Query())

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
    transactionHistory := (*database).GetTransactionHistory(params.Username)
	if transactionHistory == nil {
		api.RequestErrorHandler(w, errors.New(fmt.Sprintf("No Transaction History found for user %s", params.Username)))
		return
	}

    var response = api.TransactionHistoryResponse{
        Username: params.Username,
		Transactions: transactionHistory,
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