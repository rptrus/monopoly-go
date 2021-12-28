package game_objects

type Board struct {
	MonopolySpace [40]Square
}

type Square struct {
	SquareType int
}

type GameState struct {
	CurrentPlayer   *Player
	CurrentDiceRoll int
}

const (
	buildableProperty int = iota
	chanceChest
	tax
	station
	noAction
	utility
	action
)
