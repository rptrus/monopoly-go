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

func (pd *PropertyDeed) PayRent(from *Player, to *Player) (int, error) {
	if (*from).PlayerNumber == (*to).PlayerNumber {
		fmt.Println("Don't pay rent to ourselves") // not really an error
		return 0, nil
	}
	if (*from).CashAvailable-pd.Rent < 0 {
		str := []string{"Player does not have enough funds to cover rent ", string(pd.Rent)}
		return 0, errors.New(strings.Join(str, " "))
	}
	from.CashAvailable -= pd.Rent
	to.CashAvailable += pd.Rent
	return pd.Rent, nil
}
