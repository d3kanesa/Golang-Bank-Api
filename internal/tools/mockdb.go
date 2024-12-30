package tools

import (
	"fmt"
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"deshna": {
		AuthToken: "abc",
		Username:  "deshna",
	},
	"dhirhan": {
		AuthToken: "def",
		Username:  "dhirhan",
	},
	"darun": {
		AuthToken: "ghi",
		Username:  "darun",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"deshna": {
		Coins:    10,
		Username: "deshna",
	},
	"dhirhan": {
		Coins:    100,
		Username: "dhirhan",
	},
	"darun": {
		Coins:    1000,
		Username: "darun",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}

func (d *mockDB) ModifyUserCoins(username string, amount int64, isDeposit bool) (*CoinDetails, string) {
	time.Sleep(time.Second * 1)
	if amount <= 0 {
		return nil, "amount must be positive"
	}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil, "user not found"
	}
	if !isDeposit && clientData.Coins < amount {
		return nil, "insufficient funds"
	}

	if isDeposit {
		clientData.Coins += amount
	} else {
		clientData.Coins -= amount
	}
	mockCoinDetails[username] = clientData
	return &clientData, ""
}

func (d *mockDB) TransferCoins(sender string, receiver string, amount int64) (*CoinDetails, string) {
	time.Sleep(time.Second * 1)
	fmt.Println(sender, receiver, amount)
	if amount <= 0 {
		return nil, "amount must be positive"
	}
	fmt.Println("reaches")
	var senderData = CoinDetails{}
	senderData, ok := mockCoinDetails[sender]
	if !ok {
		return nil, "sender not found"
	}
	fmt.Println("reaches")
	if (senderData.Coins - amount) < 0 {
		return nil, "insufficient balance"
	}
	fmt.Println("reaches")
	var receiverData = CoinDetails{}
	receiverData, ok = mockCoinDetails[receiver]
	if !ok {
		return nil, "receiver not found"
	}
	fmt.Println("reaches")
	senderData.Coins -= amount
	receiverData.Coins += amount
	mockCoinDetails[sender] = senderData
	mockCoinDetails[receiver] = receiverData
	return &senderData, ""
}

func (d *mockDB) CreateUser(username string, authtoken string, coins int64) (*LoginDetails, string) {
	time.Sleep(time.Second * 1)
	if coins < 0 {
		return nil, "coins can't be negative"
	}
	_, ok := mockLoginDetails[username]
	if ok {
		return nil, "User already exists"
	}
	mockLoginDetails[username] = LoginDetails{
		Username:  username,
		AuthToken: authtoken,
	}
	mockCoinDetails[username] = CoinDetails{
		Username: username,
		Coins:    coins,
	}
	clientData := mockLoginDetails[username]
	return &clientData, ""
}
