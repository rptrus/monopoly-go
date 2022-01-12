package game_objects

import (
	"errors"
	"fmt"
)

// every one will have a plus or minus side
type Transaction struct {
	sender   *Player
	receiver *Player
	amount   int
}

const (
	half       float64 = 0.5
	tenPercent float64 = 0.1
)

var BankGameState *GameState

type Bank struct {
	CashReservesInDollars int
	TransactionLedger     []Transaction
	TotalHouses           int
	TotalHotels           int
}

// Player gives money to the bank
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

// Bank gives money to Player for things like passing go (universal basic income!) and returning houses
func (txn *Transaction) BankCheque() {
	TheBank.CashReservesInDollars -= txn.amount
	txn.receiver.CashAvailable += txn.amount
	if txn.sender != nil {
		fmt.Errorf("Cannot set the receiver when doing a bank cheque!")
	}
	if TheBank.CashReservesInDollars <= 0 {
		panic("The bank has gone bankrupt! Game is over")
	}
}

func (txn *Transaction) TransactWithPlayer(priority byte) (int, error) {
	var (
		err               error = nil
		moneyPaid               = txn.amount
		haveEnoughToCover       = false
	)
	if txn.sender.CashAvailable < txn.amount {
		if priority == 'n' {
			err = errors.New("Insufficient cash!")
			return 0, err
		} else if priority == 'x' {
			haveEnoughToCover = txn.sellDownHouses()
			if !haveEnoughToCover {
				haveEnoughToCover = txn.mortgage()
			}
			if !haveEnoughToCover {
				txn.sender.Active = false
				allOfIt := txn.sender.CashAvailable
				txn.sender.CashAvailable -= allOfIt
				txn.receiver.CashAvailable += allOfIt
				AcquireAllMortgagedProperties(txn.receiver, txn.sender)
				fmt.Println("Player", txn.sender.Name, "is bankrupt!")
				BankGameState.RemoveToken(txn.sender)
				moneyPaid = allOfIt
			}
		}
	} else {
		txn.sender.CashAvailable -= txn.amount
		txn.receiver.CashAvailable += txn.amount
	}
	return moneyPaid, nil
}

// true: if we have enough money to pay off debts, false otherwise
func (txn *Transaction) sellDownHouses() bool {
	// cycle through the properties of the debtor
	// if they have houses, sell them down at half price first
	props := ShowPropertyDeedsOfPlayer(txn.sender.PlayerNumber, BankGameState.AllProperties)
	// represents taking off 1 house from each of the properties.
	var stillMoreHouses = true
	var noHouseCounter = 0
	howManyToLookThrough := len(props)
	for txn.sender.CashAvailable < txn.amount {
		if stillMoreHouses == false {
			break
		}
		noHouseCounter = 0
		for _, prop := range props {
			if prop.HousesOwned > 0 {
				prop.HousesOwned--
				halfCostHouse := int(half * float64(prop.HouseCost))
				// don't use the txn transaction object, we need another for working this out
				t := Transaction{
					sender:   nil,
					receiver: txn.sender,
					amount:   halfCostHouse,
				}
				t.BankCheque()
			} else {
				noHouseCounter++
			}
		}
		if noHouseCounter == howManyToLookThrough {
			stillMoreHouses = false
		}
	}
	if txn.sender.CashAvailable >= txn.amount {
		fmt.Println("Debt can be paid off after selling houses. Needed", txn.amount, "have", txn.sender.CashAvailable)
		return true
	}
	// after all houses are down, then proceed to mortgaging properties
	// if still short then assign properties to debtee/creditor
	return false
}

func (txn *Transaction) mortgage() bool {
	// TODO Check balance after each unmortgage

	_, props := ShowPropertiesOfPlayer(txn.sender.PlayerNumber, BankGameState.AllProperties)
	for _, prop := range props {
		t := Transaction{
			sender:   nil,
			receiver: txn.sender,
			amount:   int(half * float64(prop.PurchaseCost)),
		}
		t.BankCheque()
		prop.Mortgaged = true
		fmt.Println("Mortgaged", GetTheCurrentCardName(prop.PositionOnBoard, BankGameState.AllProperties), "for", t.amount)

		if txn.sender.CashAvailable >= txn.amount {
			fmt.Println("Debt can be paid off after mortgaging. Needed", txn.amount, "have", txn.sender.CashAvailable)
			return true
		}
	}
	return false
}
