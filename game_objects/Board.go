package game_objects

type Board struct {
	MonopolySpace [40]Square
}

type Square struct {
	SquareType int
}

const (
	BuildableProperty int = iota
	CommunityChest
	Chance
	Tax
	Station
	JustVisiting
	Utility
	Jail
	Payment
	FreeParking
)

func GetPropertyType(number int) string {
	var propType string
	switch number {
	case BuildableProperty:
		propType = "Property"
	case CommunityChest:
		propType = "Community Chest"
	case Chance:
		propType = "Chance"
	case Tax:
		propType = "Tax collection"
	case Station:
		propType = "Station"
	case JustVisiting:
		propType = "Just Visiting"
	case FreeParking:
		propType = "Free Parking"
	case Utility:
		propType = "Utility"
	case Jail:
		propType = "Jail"
	case Payment:
		propType = "Payment"

	}
	return propType
}
