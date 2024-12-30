package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	AuthToken string
	Username  string
}

type CoinDetails struct {
	Coins    int64
	Username string
}


type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	ModifyUserCoins(username string, amount int64, isDeposit bool) (*CoinDetails, string)
	TransferCoins(sender string, receiver string, amount int64) (*CoinDetails, string)
	SetupDatabase() error
	CreateUser(username string, authtoken string, coins int64) (*LoginDetails, string)
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}