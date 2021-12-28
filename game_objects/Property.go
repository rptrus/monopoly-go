package game_objects

// using http://www.jdawiseman.com/papers/trivia/monopoly-rents.html

type PropertyDeed struct {
	PurchaseCost   int
	Rent           int
	RentWithHouses []int
}

type Property struct {
	Card map[string]*PropertyDeed
}

type PropertyCollection struct {
	AllProperty [28]Property // there are 12 non property cards
}
