package game_objects

import "fmt"

// This file uses receiver methods for structs

//var AllPlayers []Player // slice

// board position will be zero based. GO space is zero
type Player struct {
	PlayerNumber    int
	CashAvailable   int
	PositionOnBoard int
}

func (p *Player) AdvancePlayer(steps int) {
	p.PositionOnBoard += steps
	p.PositionOnBoard = p.PositionOnBoard % placesonboard
}

func (p *Player) BuyProperty(pd *PropertyDeed) {
	//p.CashAvailable -= pd.PurchaseCost // this is also ok
	(*p).CashAvailable -= (*pd).PurchaseCost
	(*pd).Owner = byte(p.PlayerNumber)
	fmt.Println("Player ", p.PlayerNumber, "now owns property ", pd.PositionOnBoard)
}
