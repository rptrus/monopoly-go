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
	Jail
	Payment
)

func GetPropertyType(number int) string {
	var propType string
	switch number {
	case BuildableProperty:
		propType = "Property"
	case ChanceChest:
		propType = "Community Chest / Chance"
	case Tax:
		propType = "Tax collection"
	case Station:
		propType = "Station"
	case NoAction:
		propType = "Nothing"
	case Utility:
		propType = "Utility"
	case Jail:
		propType = "Jail"
	case Payment:
		propType = "Payment"

	}
	return propType
}
