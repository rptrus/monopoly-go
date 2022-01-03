package game_objects

import (
	"errors"
)

// This file uses receiver methods for structs

const (
	roundTripPayment = 200
)

// board position will be zero based. GO space is zero
type Player struct {
	PlayerNumber    int
	Name            string
	CashAvailable   int
	PositionOnBoard int
}

// return reward if GO is passed, 0 otherwise. If return results need to be augmented will create a struct in future
func (p *Player) AdvancePlayer(steps int) int {
	prePosition := p.PositionOnBoard
	p.PositionOnBoard += steps
	p.PositionOnBoard = p.PositionOnBoard % placesonboard
	if p.PositionOnBoard < prePosition {
		p.pay200Dollars()
		return roundTripPayment
	}
	return 0
}

func (p *Player) BuyProperty(pd *PropertyDeed) (int, error) {
	if p.CashAvailable-pd.PurchaseCost < 0 {
		return 0, errors.New("Cannot afford property!")
	}
	(*p).CashAvailable -= (*pd).PurchaseCost
	pd.Owner = byte(p.PlayerNumber)
	return pd.PurchaseCost, nil
}

func (p *Player) pay200Dollars() {
	p.CashAvailable += roundTripPayment
}
