package game_objects

import (
	"errors"
	"fmt"
	"strings"
)

// using http://www.jdawiseman.com/papers/trivia/monopoly-rents.html

type PropertyDeed struct {
	PositionOnBoard int
	PurchaseCost    int
	Rent            int
	RentWithHouses  []int
	Owner           byte // ['1'-'6'] or 'b' for bank. 'u' if unowned
}

type Property struct {
	Card map[string]*PropertyDeed
}

type PropertyCollection struct {
	AllProperty [28]Property // there are 12 non property cards
}

func (pd *PropertyDeed) PayRent(from *Player, to *Player, board *Board) (int, error) {
	if (*from).PlayerNumber == (*to).PlayerNumber {
		fmt.Println("Don't pay rent to ourselves")
		return 0, errors.New("RentToOurself") // *not really* an error, but a way to suppress output
	}
	if (*from).CashAvailable-pd.Rent < 0 {
		str := []string{"Player does not have enough funds to cover rent ", string(pd.Rent)}
		return 0, errors.New(strings.Join(str, " "))
	}
	switch board.MonopolySpace[from.PositionOnBoard].SquareType {
	case Utility:
		// just implementing single utility for now
		roll := rollDice()
		fmt.Println("Utility re-roll of", roll)
		pd.Rent = 4 * roll
		from.CashAvailable -= pd.Rent
		to.CashAvailable += pd.Rent
	case BuildableProperty:
		fallthrough
	default:
		from.CashAvailable -= pd.Rent
		to.CashAvailable += pd.Rent
	}
	return pd.Rent, nil
}
