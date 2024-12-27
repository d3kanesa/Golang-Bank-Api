package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code int
	Balance int64
}

type AddCoinParams struct {
	Username string
	AddAmount int64
}

type TransferCoinParams struct {
	Username string
	Receiver string
	AddAmount int64
}

type AddUserParams struct {
	Username string
	AuthToken string
	Coins int64
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