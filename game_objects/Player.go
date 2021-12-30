package game_objects

import (
	"errors"
	"fmt"
)

// This file uses receiver methods for structs

//var AllPlayers []Player // slice

// board position will be zero based. GO space is zero
type Player struct {
	PlayerNumber    int
	Name            string
	CashAvailable   int
	PositionOnBoard int
}

func (p *Player) AdvancePlayer(steps int) {
	p.PositionOnBoard += steps
	p.PositionOnBoard = p.PositionOnBoard % placesonboard
}

func (p *Player) BuyProperty(pd *PropertyDeed) (int, error) {
	if p.CashAvailable-pd.PurchaseCost < 0 {
		return 0, errors.New("Cannot afford property!")
	}
	(*p).CashAvailable -= (*pd).PurchaseCost
	pd.Owner = byte(p.PlayerNumber)
	fmt.Println("Player ", p.PlayerNumber, "now owns property ", pd.PositionOnBoard)
	return pd.PurchaseCost, nil
}
