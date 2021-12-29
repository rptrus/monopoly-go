package game_objects

type Board struct {
	MonopolySpace [40]Square
}

type Square struct {
	SquareType int
}

const (
	BuildableProperty int = iota
	ChanceChest
	Tax
	Station
	NoAction
	Utility
	Action
)
