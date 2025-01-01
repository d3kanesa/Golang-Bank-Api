package tools

import (
	"sync"
	"time"
	"github.com/google/uuid"
)

type mockDB struct{
	loginMutex sync.RWMutex
}

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

var mockTransactionHistory = make(map[string][]TransactionDetails)
var userMutexes map[string]*sync.RWMutex = make(map[string]*sync.RWMutex)

func init(){
	mockTransactionHistory["deshna"] = []TransactionDetails{}
	mockTransactionHistory["dhirhan"] = []TransactionDetails{}
	mockTransactionHistory["darun"] = []TransactionDetails{}
	userMutexes["deshna"] = &sync.RWMutex{}
	userMutexes["dhirhan"] = &sync.RWMutex{}
	userMutexes["darun"] = &sync.RWMutex{}
}

func (d *mockDB) RecordTransaction(username string, transType string, receiver string, amount int64) {
	var transaction = TransactionDetails{
		id: uuid.New().String(),
		Username:  username,
		Type:      transType,
		Receiver:  receiver,
		Amount:    amount,
		Timestamp: time.Now(),
	}
	mockTransactionHistory[username] = append([]TransactionDetails{transaction}, mockTransactionHistory[username]...)
}

func (d *mockDB) GetTransactionHistory(username string) []TransactionDetails {
	time.Sleep(time.Second * 1)
	userMutexes[username].RLock()
	defer userMutexes[username].RUnlock()
	var clientData = []TransactionDetails{}
	clientData, ok := mockTransactionHistory[username]
	if !ok {
		return nil
	}
	return clientData
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
	userMutexes[username].RLock()
	defer userMutexes[username].RUnlock()
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
	userMutexes[username].Lock()
	defer userMutexes[username].Unlock()
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
		d.RecordTransaction(username, "deposit", "", amount)
		clientData.Coins += amount
	} else {
		d.RecordTransaction(username, "withdrawl", "", amount)
		clientData.Coins -= amount
	}
	mockCoinDetails[username] = clientData
	return &clientData, ""
}

func (d *mockDB) TransferCoins(sender string, receiver string, amount int64) (*CoinDetails, string) {
	time.Sleep(time.Second * 1)
	userMutexes[sender].Lock()
	userMutexes[receiver].Lock()
	defer userMutexes[sender].Unlock()
	defer userMutexes[receiver].Unlock()

	if amount <= 0 {
		return nil, "amount must be positive"
	}
	var senderData = CoinDetails{}
	senderData, ok := mockCoinDetails[sender]
	if !ok {
		return nil, "sender not found"
	}
	if (senderData.Coins - amount) < 0 {
		return nil, "insufficient balance"
	}
	var receiverData = CoinDetails{}
	receiverData, ok = mockCoinDetails[receiver]
	if !ok {
		return nil, "receiver not found"
	}

	senderData.Coins -= amount
	receiverData.Coins += amount
	mockCoinDetails[sender] = senderData
	mockCoinDetails[receiver] = receiverData
	d.RecordTransaction(sender, "transfer", receiver, amount)
	return &senderData, ""
}

func (d *mockDB) CreateUser(username string, authtoken string, coins int64) (*LoginDetails, string) {
	time.Sleep(time.Second * 1)
	d.loginMutex.Lock()
	defer d.loginMutex.Unlock()
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
	mockTransactionHistory[username] = []TransactionDetails{}
	userMutexes[username] = &sync.RWMutex{}
	clientData := mockLoginDetails[username]
	return &clientData, ""
}
