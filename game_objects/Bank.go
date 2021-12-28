package game_objects

// every one will have a plus or minus side
type Transaction struct {
	sender   *Player
	receiver *Player
	amount   int
}

type Bank struct {
	CashReservesInDollars int
	TransactionLedger     []Transaction
}
