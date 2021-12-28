package game_objects

type Board struct {
	MonopolySpace [40]Square
}

type Square struct {
	SquareType int
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
