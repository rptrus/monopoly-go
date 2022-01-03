package setup

// This file demonstrates the use of slice literals and arrays inside structs

import (
	"github.com/rptrus/monopoly-go/game_objects"
)

func InitializePropertyCards() *game_objects.PropertyCollection {

	var prop *game_objects.Property
	prop = new(game_objects.Property) // pointer to instance

	PropertyCollection := new(game_objects.PropertyCollection)

	// out of struct
	/*
		m:=make(map[string] *game_objects.PropertyDeed)
		m["S"] = &game_objects.PropertyDeed{
			PurchaseCost:   60,
			Rent:           2,
			RentWithHouses: []int{2,10,30,90,160,250}, // creates another slice literal by first creating the underlying array
		}
	*/
	// in struct

	// SIDE 1

	prop.Card = make(map[string]*game_objects.PropertyDeed)

	(*prop).Card["Old Kent Road"] = &game_objects.PropertyDeed{
		PositionOnBoard: 1,
		PurchaseCost:    60,
		Rent:            2,
		RentWithHouses:  []int{2, 10, 30, 90, 160, 250}, // creates another slice literal by first creating the underlying array
	}
	PropertyCollection.AllProperty[0] = *prop

	prop = new(game_objects.Property)                       // new pointer to instance
	prop.Card = make(map[string]*game_objects.PropertyDeed) // therefore new map needed each time
	prop.Card["Whitechapel Road"] = &game_objects.PropertyDeed{
		PositionOnBoard: 3,
		PurchaseCost:    60,
		Rent:            4,
		RentWithHouses:  []int{4, 20, 60, 180, 320, 450},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[1] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Kings Cross Station"] = &game_objects.PropertyDeed{
		PositionOnBoard: 5,
		PurchaseCost:    200,
		Rent:            25, // base cost
		RentWithHouses:  nil,
		Owner:           'u',
	}
	PropertyCollection.AllProperty[2] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["The Angel Islington"] = &game_objects.PropertyDeed{
		PositionOnBoard: 6,
		PurchaseCost:    100,
		Rent:            6,
		RentWithHouses:  []int{30, 90, 270, 400, 550},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[3] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Euston Road"] = &game_objects.PropertyDeed{
		PositionOnBoard: 8,
		PurchaseCost:    100,
		Rent:            6,
		RentWithHouses:  []int{30, 90, 270, 400, 550},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[4] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Pentonville Road"] = &game_objects.PropertyDeed{
		PositionOnBoard: 9,
		PurchaseCost:    120,
		Rent:            8,
		RentWithHouses:  []int{40, 100, 300, 450, 600},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[5] = *prop

	// SIDE 2

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Pall Mall"] = &game_objects.PropertyDeed{
		PositionOnBoard: 11,
		PurchaseCost:    140,
		Rent:            10,
		RentWithHouses:  []int{50, 140, 450, 625, 750},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[6] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Electric Company"] = &game_objects.PropertyDeed{
		PositionOnBoard: 12,
		PurchaseCost:    150,
		Rent:            -1, // special case
		RentWithHouses:  nil,
		Owner:           'u',
	}
	PropertyCollection.AllProperty[7] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Whitehall"] = &game_objects.PropertyDeed{
		PositionOnBoard: 13,
		PurchaseCost:    140,
		Rent:            10,
		RentWithHouses:  []int{50, 150, 450, 625, 750},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[8] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Northumberland Ave"] = &game_objects.PropertyDeed{
		PositionOnBoard: 14,
		PurchaseCost:    140,
		Rent:            10,
		RentWithHouses:  []int{50, 150, 450, 625, 750},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[9] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Marlyebone Station"] = &game_objects.PropertyDeed{
		PositionOnBoard: 15,
		PurchaseCost:    200,
		Rent:            25,
		RentWithHouses:  nil,
		Owner:           'u',
	}
	PropertyCollection.AllProperty[10] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Bow Street"] = &game_objects.PropertyDeed{
		PositionOnBoard: 16,
		PurchaseCost:    180,
		Rent:            14,
		RentWithHouses:  []int{70, 200, 550, 750, 950},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[11] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Marlborough Street"] = &game_objects.PropertyDeed{
		PositionOnBoard: 18,
		PurchaseCost:    180,
		Rent:            14,
		RentWithHouses:  []int{70, 200, 550, 750, 950},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[12] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Vine Street"] = &game_objects.PropertyDeed{
		PositionOnBoard: 19,
		PurchaseCost:    200,
		Rent:            16,
		RentWithHouses:  []int{80, 220, 600, 800, 1000},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[13] = *prop

	// SIDE 3

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["The Strand"] = &game_objects.PropertyDeed{
		PositionOnBoard: 21,
		PurchaseCost:    220,
		Rent:            18,
		RentWithHouses:  []int{90, 250, 700, 875, 1050},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[14] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Fleet Street"] = &game_objects.PropertyDeed{
		PositionOnBoard: 23,
		PurchaseCost:    220,
		Rent:            20,
		RentWithHouses:  []int{90, 250, 700, 875, 1050},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[15] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Trafalgar Square"] = &game_objects.PropertyDeed{
		PositionOnBoard: 24,
		PurchaseCost:    240,
		Rent:            20,
		RentWithHouses:  []int{100, 300, 750, 925, 1100},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[16] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Fenchurch Street Station"] = &game_objects.PropertyDeed{
		PositionOnBoard: 25,
		PurchaseCost:    200,
		Rent:            25,
		RentWithHouses:  nil,
		Owner:           'u',
	}
	PropertyCollection.AllProperty[17] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Leicester Square"] = &game_objects.PropertyDeed{
		PositionOnBoard: 26,
		PurchaseCost:    260,
		Rent:            22,
		RentWithHouses:  []int{110, 330, 800, 975, 1150},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[18] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Coventry Street"] = &game_objects.PropertyDeed{
		PositionOnBoard: 27,
		PurchaseCost:    260,
		Rent:            22,
		RentWithHouses:  []int{110, 330, 800, 975, 1150},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[19] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Water Works"] = &game_objects.PropertyDeed{
		PositionOnBoard: 28,
		PurchaseCost:    150,
		Rent:            -1,
		RentWithHouses:  nil,
		Owner:           'u',
	}
	PropertyCollection.AllProperty[20] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Picadilly"] = &game_objects.PropertyDeed{
		PositionOnBoard: 29,
		PurchaseCost:    280,
		Rent:            22,
		RentWithHouses:  []int{120, 360, 850, 1025, 1200},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[21] = *prop

	// SIDE 4

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Regent Street"] = &game_objects.PropertyDeed{
		PositionOnBoard: 31,
		PurchaseCost:    300,
		Rent:            26,
		RentWithHouses:  []int{130, 390, 900, 1100, 1275},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[22] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Oxford Street"] = &game_objects.PropertyDeed{
		PositionOnBoard: 32,
		PurchaseCost:    300,
		Rent:            26,
		RentWithHouses:  []int{130, 390, 900, 1100, 1275},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[23] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Bond Street"] = &game_objects.PropertyDeed{
		PositionOnBoard: 34,
		PurchaseCost:    320,
		Rent:            28,
		RentWithHouses:  []int{150, 450, 1000, 1200, 1400},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[24] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Liverpool Street Station"] = &game_objects.PropertyDeed{
		PositionOnBoard: 35,
		PurchaseCost:    200,
		Rent:            25,
		RentWithHouses:  nil,
		Owner:           'u',
	}
	PropertyCollection.AllProperty[25] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Park Lane"] = &game_objects.PropertyDeed{
		PositionOnBoard: 37,
		PurchaseCost:    350,
		Rent:            35,
		RentWithHouses:  []int{175, 500, 1100, 1300, 1500},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[26] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Mayfair"] = &game_objects.PropertyDeed{
		PositionOnBoard: 39,
		PurchaseCost:    400,
		Rent:            50,
		RentWithHouses:  []int{200, 600, 1400, 1700, 2000},
		Owner:           'u',
	}
	PropertyCollection.AllProperty[27] = *prop

	// another potential way to initiailze, would need individual vars though
	//another := game_objects.PropertyCollection{AllProperty: [28]game_objects.Property{*prop,*prop,*prop}}
	//fmt.Println(another)
	return PropertyCollection

}

// These are cards that typically cannot be owned, built upon and don't pay rent
func InitializeNonPropertyCards() *game_objects.OtherPropertyCollection {

	var nonprop *game_objects.OtherProperty
	OtherPropertyCollection := new(game_objects.OtherPropertyCollection)

	nonprop = new(game_objects.OtherProperty) // pointer to instance
	nonprop.Card = make(map[string]*game_objects.OtherPropertyDetail)
	(*nonprop).Card["GO"] = &game_objects.OtherPropertyDetail{
		PositionOnBoard: 0,
		PlayerTax:       0,
	}
	OtherPropertyCollection.AllProperty[0] = *nonprop

	nonprop = new(game_objects.OtherProperty) // pointer to instance
	nonprop.Card = make(map[string]*game_objects.OtherPropertyDetail)
	(*nonprop).Card["Community Chest"] = &game_objects.OtherPropertyDetail{
		PositionOnBoard: 2,
		PlayerTax:       0,
	}
	OtherPropertyCollection.AllProperty[1] = *nonprop

	nonprop = new(game_objects.OtherProperty) // pointer to instance
	nonprop.Card = make(map[string]*game_objects.OtherPropertyDetail)
	(*nonprop).Card["Income Tax"] = &game_objects.OtherPropertyDetail{
		PositionOnBoard: 4,
		PlayerTax:       200,
	}
	OtherPropertyCollection.AllProperty[2] = *nonprop

	nonprop = new(game_objects.OtherProperty) // pointer to instance
	nonprop.Card = make(map[string]*game_objects.OtherPropertyDetail)
	(*nonprop).Card["Chance"] = &game_objects.OtherPropertyDetail{
		PositionOnBoard: 7,
		PlayerTax:       0,
	}
	OtherPropertyCollection.AllProperty[3] = *nonprop

	nonprop = new(game_objects.OtherProperty) // pointer to instance
	nonprop.Card = make(map[string]*game_objects.OtherPropertyDetail)
	(*nonprop).Card["Just Visiting"] = &game_objects.OtherPropertyDetail{
		PositionOnBoard: 10,
		PlayerTax:       0,
	}
	OtherPropertyCollection.AllProperty[4] = *nonprop

	nonprop = new(game_objects.OtherProperty) // pointer to instance
	nonprop.Card = make(map[string]*game_objects.OtherPropertyDetail)
	(*nonprop).Card["Community Chest"] = &game_objects.OtherPropertyDetail{
		PositionOnBoard: 17,
		PlayerTax:       0,
	}
	OtherPropertyCollection.AllProperty[5] = *nonprop

	nonprop = new(game_objects.OtherProperty) // pointer to instance
	nonprop.Card = make(map[string]*game_objects.OtherPropertyDetail)
	(*nonprop).Card["Free Parking"] = &game_objects.OtherPropertyDetail{
		PositionOnBoard: 20,
		PlayerTax:       0,
	}
	OtherPropertyCollection.AllProperty[6] = *nonprop

	nonprop = new(game_objects.OtherProperty) // pointer to instance
	nonprop.Card = make(map[string]*game_objects.OtherPropertyDetail)
	(*nonprop).Card["Chance"] = &game_objects.OtherPropertyDetail{
		PositionOnBoard: 22,
		PlayerTax:       0,
	}
	OtherPropertyCollection.AllProperty[7] = *nonprop

	nonprop = new(game_objects.OtherProperty) // pointer to instance
	nonprop.Card = make(map[string]*game_objects.OtherPropertyDetail)
	(*nonprop).Card["Go To Jail"] = &game_objects.OtherPropertyDetail{
		PositionOnBoard: 30,
		PlayerTax:       0,
	}
	OtherPropertyCollection.AllProperty[8] = *nonprop

	nonprop = new(game_objects.OtherProperty) // pointer to instance
	nonprop.Card = make(map[string]*game_objects.OtherPropertyDetail)
	(*nonprop).Card["Community Chest"] = &game_objects.OtherPropertyDetail{
		PositionOnBoard: 33,
		PlayerTax:       0,
	}
	OtherPropertyCollection.AllProperty[9] = *nonprop

	nonprop = new(game_objects.OtherProperty) // pointer to instance
	nonprop.Card = make(map[string]*game_objects.OtherPropertyDetail)
	(*nonprop).Card["Chance"] = &game_objects.OtherPropertyDetail{
		PositionOnBoard: 36,
		PlayerTax:       0,
	}
	OtherPropertyCollection.AllProperty[10] = *nonprop

	nonprop = new(game_objects.OtherProperty) // pointer to instance
	nonprop.Card = make(map[string]*game_objects.OtherPropertyDetail)
	(*nonprop).Card["Chance"] = &game_objects.OtherPropertyDetail{
		PositionOnBoard: 38,
		PlayerTax:       100,
	}
	OtherPropertyCollection.AllProperty[11] = *nonprop

	return OtherPropertyCollection
}
