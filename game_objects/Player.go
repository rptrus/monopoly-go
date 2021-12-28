package game_objects

// board position will be zero based. GO space is zero
type Player struct {
	PlayerNumber    int
	CashAvailable   int
	PositionOnBoard int
}
