package game_objects

// This file uses receiver methods for structs

//var allPlayers [6]Player // array
var AllPlayers []Player // slice

// board position will be zero based. GO space is zero
type Player struct {
	PlayerNumber    int
	CashAvailable   int
	PositionOnBoard int
}

func (p *Player) AdvancePlayer(steps int) {
	p.PositionOnBoard += steps
}
