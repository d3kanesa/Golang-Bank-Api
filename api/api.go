package api

import (
	"encoding/json"
	"net/http"
	"apiProject/internal/tools"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code int
	Balance int64
}

type TransactionHistoryResponse struct {
	Username string
	Transactions []tools.TransactionDetails
	Code int
}

type ModifyCoinParams struct {
	Username string `json:"username" validate:"required"`
	ModifyAmount int64 `json:"modifyAmount" validate:"required"`
}

type TransferCoinParams struct {
	Username string `json:"username" validate:"required"`
	Receiver string `json:"receiver" validate:"required"`
	AddAmount int64 `json:"addAmount" validate:"required"`
}

type AddUserParams struct {
	Username string `json:"username" validate:"required"`
	AuthToken string `json:"authtoken" validate:"required"`
	Coins int64 `json:"coins" validate:"required"`
} 

type Error struct {
	Code int
	Message string
}

type UserCreateResponse struct {
	Username string
	Code int
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter){
		writeError(w, "An Unexpected ERror Occured", http.StatusInternalServerError)
	}
)