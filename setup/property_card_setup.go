package setup

// This file demonstrates the use of slice literals and arrays inside structs

import (
	"fmt"
	"github.com/rptrus/monopoly-go/game_objects"
)

func SetupPropertyCards() {

	var prop *game_objects.Property
	prop = new(game_objects.Property) // pointer to instance

	propertyCollection := new(game_objects.PropertyCollection)

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
	prop.Card = make(map[string]*game_objects.PropertyDeed)

	(*prop).Card["Old Kent Road"] = &game_objects.PropertyDeed{
		PurchaseCost:   60,
		Rent:           2,
		RentWithHouses: []int{2, 10, 30, 90, 160, 250}, // creates another slice literal by first creating the underlying array
	}
	propertyCollection.AllProperty[0] = *prop

	prop = new(game_objects.Property)                       // new pointer to instance
	prop.Card = make(map[string]*game_objects.PropertyDeed) // therefore new map needed each time
	prop.Card["Whitechapel Road"] = &game_objects.PropertyDeed{
		PurchaseCost:   60,
		Rent:           4,
		RentWithHouses: []int{4, 20, 60, 180, 320, 450},
	}
	propertyCollection.AllProperty[1] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["The Angel Islington"] = &game_objects.PropertyDeed{
		PurchaseCost:   100,
		Rent:           6,
		RentWithHouses: []int{30, 90, 270, 400, 550},
	}
	propertyCollection.AllProperty[2] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Euston Road"] = &game_objects.PropertyDeed{
		PurchaseCost:   100,
		Rent:           6,
		RentWithHouses: []int{30, 90, 270, 400, 550},
	}
	propertyCollection.AllProperty[3] = *prop

	prop = new(game_objects.Property)
	prop.Card = make(map[string]*game_objects.PropertyDeed)
	prop.Card["Pentonville Road"] = &game_objects.PropertyDeed{
		PurchaseCost:   120,
		Rent:           8,
		RentWithHouses: []int{40, 100, 300, 450, 600},
	}
	propertyCollection.AllProperty[4] = *prop

	// another way to initiailze
	another := game_objects.PropertyCollection{AllProperty: [40]game_objects.Property{*prop}}
	fmt.Println(another)

}
