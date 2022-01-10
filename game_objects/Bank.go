package game_objects

import (
	"errors"
	"fmt"
)

// every one will have a plus or minus side
type Transaction struct {
	gs       *GameState
	sender   *Player
	receiver *Player
	amount   int
}

var BankGameState *GameState

type Bank struct {
	CashReservesInDollars int
	TransactionLedger     []Transaction
	TotalHouses           int
	TotalHotels           int
}

func (txn *Transaction) TransactWithBank() {
	if txn.sender.CashAvailable < txn.amount {
		txn.amount = txn.sender.CashAvailable
		txn.sender.CashAvailable -= txn.amount
		fmt.Println("Player", txn.sender.Name, "is bankrupt!")
		txn.sender.Active = false
		BankGameState.RemoveToken(txn.sender)
	}
	txn.sender.CashAvailable -= txn.amount
	TheBank.CashReservesInDollars += txn.amount
}

func (txn *Transaction) BankCheque() {
	TheBank.CashReservesInDollars -= txn.amount
	txn.receiver.CashAvailable += 200
	if TheBank.CashReservesInDollars <= 0 {
		panic("The bank has gone bankrupt! Game is over")
	}
}

func (txn *Transaction) TransactWithPlayer(priority byte) error {
	var err error = nil
	if txn.sender.CashAvailable < txn.amount {
		if priority == 'n' {
			err = errors.New("Insufficient cash!")
			return err
		} else if priority == 'x' {
			txn.sender.Active = false
			txn.sender.CashAvailable -= txn.amount
			txn.receiver.CashAvailable += txn.amount
			fmt.Println("Player", txn.sender.Name, "is bankrupt!")
			BankGameState.RemoveToken(txn.sender)
		}
	} else {
		txn.sender.CashAvailable -= txn.amount
		txn.receiver.CashAvailable += txn.amount
	}
	return nil
}
