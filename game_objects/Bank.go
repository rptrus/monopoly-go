package game_objects

// every one will have a plus or minus side
type Transaction struct {
	sender   *player
	receiver *player
	amount   int
}

type Bank struct {
	CashReservesInDollars int
	TransactionLedger     []Transaction
}
