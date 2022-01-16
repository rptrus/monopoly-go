package game_objects

import (
	"errors"
	"fmt"
)

// every one will have a plus or minus side
type Transaction struct {
	Sender   *Player
	Receiver *Player
	Amount   int
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
	if txn.Sender.Active == true {
		if txn.Sender.CashAvailable < txn.Amount {
			fmt.Println("Paying a partial amount to Bank, player out of cash!")
			txn.Amount = txn.Sender.CashAvailable
			txn.Sender.Active = false
			BankGameState.RemoveToken(txn.Sender)
		}
		txn.Sender.CashAvailable -= txn.Amount
		TheBank.CashReservesInDollars += txn.Amount
		TheBank.TransactionLedger = append(TheBank.TransactionLedger, *txn)
	}
}

// Bank gives money to Player for things like passing go (universal basic income!) and returning houses
func (txn *Transaction) BankCheque() {
	TheBank.CashReservesInDollars -= txn.Amount
	txn.Receiver.CashAvailable += txn.Amount
	if txn.Sender != nil {
		fmt.Errorf("Cannot set the receiver when doing a bank cheque!")
	}
	if TheBank.CashReservesInDollars <= 0 {
		panic("The bank has gone bankrupt! Game is over")
	}
	TheBank.TransactionLedger = append(TheBank.TransactionLedger, *txn)
}

func (txn *Transaction) TransactWithPlayer(priority byte) (int, error) {
	var (
		err               error = nil
		moneyPaid               = txn.Amount
		haveEnoughToCover       = false
	)
	if txn.Sender.CashAvailable < txn.Amount {
		if priority == 'n' {
			err = errors.New("Insufficient cash!")
			return 0, err
		} else if priority == 'x' {
			haveEnoughToCover = txn.sellDownHouses()
			if !haveEnoughToCover {
				haveEnoughToCover = txn.mortgage()
			}
			if !haveEnoughToCover {
				txn.Sender.Active = false
				allOfIt := txn.Sender.CashAvailable
				txn.Sender.CashAvailable -= allOfIt
				txn.Receiver.CashAvailable += allOfIt
				AcquireAllMortgagedProperties(txn.Receiver, txn.Sender)
				fmt.Println("[TWP] Player", txn.Sender.Name, "is bankrupt!")
				BankGameState.RemoveToken(txn.Sender)
				moneyPaid = allOfIt
				err = errors.New("Partial-Payment")
			}
		}
	} else {
		txn.Sender.CashAvailable -= txn.Amount
		txn.Receiver.CashAvailable += txn.Amount
	}
	TheBank.TransactionLedger = append(TheBank.TransactionLedger, *txn)
	return moneyPaid, err
}

// true: if we have enough money to pay off debts, false otherwise
func (txn *Transaction) sellDownHouses() bool {
	// cycle through the properties of the debtor
	// if they have houses, sell them down at half price first
	props := ShowPropertyDeedsOfPlayer(txn.Sender.PlayerNumber, BankGameState.AllProperties)
	// represents taking off 1 house from each of the properties.
	var stillMoreHouses = true
	var noHouseCounter = 0
	howManyToLookThrough := len(props)
	for txn.Sender.CashAvailable < txn.Amount {
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
					Sender:   nil,
					Receiver: txn.Sender,
					Amount:   halfCostHouse,
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
	if txn.Sender.CashAvailable >= txn.Amount {
		fmt.Println("Debt can be paid off after selling houses. Needed", txn.Amount, "have", txn.Sender.CashAvailable)
		return true
	}
	// after all houses are down, then proceed to mortgaging properties
	// if still short then assign properties to debtee/creditor
	return false
}

func (txn *Transaction) mortgage() bool {
	_, props := ShowPropertiesOfPlayer(txn.Sender.PlayerNumber, BankGameState.AllProperties)
	for _, prop := range props {
		t := Transaction{
			Sender:   nil,
			Receiver: txn.Sender,
			Amount:   int(half * float64(prop.PurchaseCost)),
		}
		t.BankCheque()
		prop.Mortgaged = true
		fmt.Println("Mortgaged", GetTheCurrentCardName(prop.PositionOnBoard, BankGameState.AllProperties), "for", t.Amount)

		if txn.Sender.CashAvailable >= txn.Amount {
			fmt.Println("Debt can be paid off after mortgaging. Needed", txn.Amount, "have", txn.Sender.CashAvailable)
			t := Transaction{
				Sender:   txn.Sender,
				Receiver: txn.Receiver,
				Amount:   txn.Amount,
			}
			t.TransactWithPlayer('x')
			return true
		}
	}
	return false
}
