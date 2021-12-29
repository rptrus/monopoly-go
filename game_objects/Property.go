package game_objects

// using http://www.jdawiseman.com/papers/trivia/monopoly-rents.html

type PropertyDeed struct {
	PositionOnBoard int
	PurchaseCost    int
	Rent            int
	RentWithHouses  []int
	Owner           byte // ['1'-'6'] or 'b' for bank. 'u' if unowned
}

type Property struct {
	Card map[string]*PropertyDeed
}

type PropertyCollection struct {
	AllProperty [28]Property // there are 12 non property cards
}
