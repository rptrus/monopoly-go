package game_objects

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// using http://www.jdawiseman.com/papers/trivia/monopoly-rents.html

type PropertyDeed struct {
	Set             string
	PositionOnBoard int
	PurchaseCost    int
	Rent            int
	RentWithHouses  []int
	Owner           byte // [1-6] or 'u' for bank. 'u' is unowned
	HouseCost       int
	HousesOwned     int
	Mortgaged		bool
}

type arrayOfPropertyDeed []*PropertyDeed

func (p arrayOfPropertyDeed) Len() int {
	return len(p)
}

func (p arrayOfPropertyDeed) Less(i, j int) bool {
	return p[i].PurchaseCost > p[j].PurchaseCost
}

func (p arrayOfPropertyDeed) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Property struct {
	Card map[string]*PropertyDeed
}

type PropertyCollection struct {
	AllProperty [28]Property // there are 12 non property cards
}

// encapsulates things like tax, go to jail and eventually community chest/chance
type OtherPropertyDetail struct {
	PositionOnBoard int
	PlayerTax       int // if applicable (e.g. tax, supertax)
	moveToSquare    int // if defined, it will move a player to paricluar square (e.g. go to jail) -1 if undefined
}

type OtherProperty struct {
	Card map[string]*OtherPropertyDetail
}

type OtherPropertyCollection struct {
	AllProperty [12]OtherProperty
}

func (pd *PropertyDeed) PayRent(from *Player, to *Player, board *Board, pc *PropertyCollection) (int, error) {
	if pd.Mortgaged {
		return 0, nil
	}
	t := Transaction{
		sender:   from,
		receiver: to,
		amount:   0,
	}
	var actualPaid = 0 // a player may not be able to pay the full amount
	if (*from).PlayerNumber == (*to).PlayerNumber {
		fmt.Println("Don't pay rent to ourselves")
		return 0, errors.New("RentToOurself") // *not really* an error, but a way to suppress output
	}
	switch board.MonopolySpace[from.PositionOnBoard].SquareType {
	case Utility:
		ownsBoth := len(findSameType(board, pd, pc)) == 2
		roll := rollDice()
		fmt.Println("Utility re-roll of", roll)
		if ownsBoth {
			pd.Rent = 10 * roll
		} else {
			pd.Rent = 4 * roll
		}
		t.amount = pd.Rent
		actualPaid, _ = t.TransactWithPlayer('x')
	case Station:
		stationsOwnedByPlayer := len(findSameType(board, pd, pc))
		pd.Rent = stationsOwnedByPlayer * 25
		t.amount = pd.Rent
		actualPaid, _ = t.TransactWithPlayer('x')
	case BuildableProperty:
		// check if the property landed on is a complete set
		moneyOwing := pd.Rent
		multiplyFactor := 1
		hasAllSet := checkCompleteSet(pd, pc)
		if hasAllSet && pd.HousesOwned == 0 {
			multiplyFactor = 2
		}
		if pd.HousesOwned > 0 {
			moneyOwing = pd.RentWithHouses[pd.HousesOwned-1]
		}
		t.amount = moneyOwing * multiplyFactor
		actualPaid, _ = t.TransactWithPlayer('x')
	default:
		fmt.Println("Unknown or not implemented", board.MonopolySpace[from.PositionOnBoard].SquareType)
	}
	var addition = ""
	if !(pd.Set == "Train" || pd.Set == "Utility") {
		addition = "with "+strconv.Itoa(pd.HousesOwned)+" houses"
	}
	fmt.Println("Cost of landing on property", GetTheCurrentCardName(pd.PositionOnBoard, pc), "is: $", t.amount,addition)
	return actualPaid, nil
}

func swapPropertyBetweenPlayers(from *Player, to *Player, card *PropertyDeed, pc *PropertyCollection) {
	fmt.Println("Player", from.Name, "Will give property", GetTheCurrentCardName(card.PositionOnBoard, pc), "to", to.Name)
	// since we will be swapping properties later, we don't need to adjust cash here
	card.Owner = byte(to.PlayerNumber)
	fmt.Println("Now assigned")
}

// Given a square of a particular type, find all the others of that type
// This is useful for rent calculations for utilities and stations
func findSameType(board *Board, pd *PropertyDeed, pc *PropertyCollection) []byte {
	var similars []int
	var singleOwnerCount []byte
	for i, j := range board.MonopolySpace {
		if j.SquareType == board.MonopolySpace[pd.PositionOnBoard].SquareType {
			similars = append(similars, i) // remember the position on the board
		}
	}
	for _, l := range similars {
		for _, n := range pc.AllProperty {
			aSingularCardMap := n.Card
			for _, v := range aSingularCardMap {
				if v.PositionOnBoard == l && v.Owner == pd.Owner {
					singleOwnerCount = append(singleOwnerCount, v.Owner) // will increment if only if owned by same player
				}
			}
		}
	}
	return singleOwnerCount
}

// true if all properties are owned by one and only one owner
// At most we would expect 2 false, since the owner owns at least one!
// easier to handle the case by exception. Assume owns all, until a counterexample emerges
func checkCompleteSet(pd *PropertyDeed, pc *PropertyCollection) bool {
	var ownsAll = true
	propsInSet, setCounter := propsOwnedByPlayerInASet(pd, pc)
	if len(propsInSet) == setCounter {
		ownsAll = true
	}
	return ownsAll
}

// Given a property deed card of a particular set, find the positions (i.e. ownership) of other cards in a set
func propsOwnedByPlayerInASet(pd *PropertyDeed, pc *PropertyCollection) ([]int, int) {
	var propsInSet []int
	var setCounter int
	for _, property := range pc.AllProperty {
		for _, v := range property.Card {
			if v.Set == pd.Set { // same colour as our input property deed
				if v.Owner == pd.Owner {
					propsInSet = append(propsInSet, v.PositionOnBoard)
				}
				setCounter++
			}
		}
	}
	return propsInSet, setCounter
}

// input a set colour, get the owners
func ownersOfASet(setColour string, pc *PropertyCollection) ([]byte, bool) {
	var owners []byte
	var bankOwned bool = false
	for _, property := range pc.AllProperty { // over array
		for _, v := range property.Card { // over 1-map
			if v.Set == setColour {
				owners = append(owners, v.Owner)
				if v.Owner == 'u' {
					bankOwned = true
				}
			}
		}
	}
	return owners, bankOwned
}

// input: player number
// output: properties owned
func ShowPropertiesOfPlayer(playerNumber int, pc *PropertyCollection) ([]string, arrayOfPropertyDeed) {
	var M = ""
	propsOwnedNameOnly := []string{}
	propDeeds := []*PropertyDeed{}
	for _, card := range pc.AllProperty {
		aSingularCardMap := card.Card
		for _, v := range aSingularCardMap {
			if int(v.Owner) == playerNumber {
				n, pd := GetTheCurrentCard(v.PositionOnBoard, pc)
				if pd.Mortgaged == true {
					M=" (M)"
				}
				propsOwnedNameOnly = append(propsOwnedNameOnly, n+M)
				propDeeds = append(propDeeds, pd)
			}
		}
	}
	return propsOwnedNameOnly, propDeeds
}

func (gs *GameState) UnownedProperties(pc *PropertyCollection) {
	var propsSpare []string
	for _, props := range pc.AllProperty {
		for _, k := range props.Card { // 1 element map
			if k.Owner == 'u' {
				propsSpare = append(propsSpare, GetTheCurrentCardName(k.PositionOnBoard, pc))
			}
		}
	}
	if (len(propsSpare)) > 0 {
		fmt.Println("Outstanding properties to be purchased:")
		fmt.Print(len(propsSpare), ") -> \"", strings.Join(propsSpare, "\",\""), "\" \n")
	} else {
		if gs.allPropsSold == false {
			fmt.Println("* ALL PROPERTIES HAVE NOW SOLD! *")
			gs.allPropsSold = true
		}
	}
}

// who owns the set besides us when we have a majority
func OtherOwnerOfSet(playerNum int, owners []byte) byte {
	var otherPlayer byte
	for _, j := range owners {
		if int(j) != playerNum {
			otherPlayer = j
		}
	}
	return otherPlayer
}

// we consider 2 of a 3-set or 1 of a 2-set to be the highest partially completed set
func highestPartiallyCompleteSet(otherPlayer byte, AllPlayers []Player, pc *PropertyCollection) []*PropertyDeed {
	var setsWithMostPropertiesOwned []*PropertyDeed
	_, deeds := ShowPropertiesOfPlayer(int(otherPlayer), pc)
	for _, pd := range deeds {
		owned, totalInSet := propsOwnedByPlayerInASet(pd, pc)
		if len(owned) == 2 && totalInSet == 3 || len(owned) == 1 && totalInSet == 2 {
			setsWithMostPropertiesOwned = append(setsWithMostPropertiesOwned, pd)
		}
	}
	return setsWithMostPropertiesOwned
}

// needs work for things like utility / train station. works ok for coloured property sets
func ownsFullSet(properties []*PropertyDeed, pc *PropertyCollection) []string {
	var setsOwned []string
	// we are already using sort for purchase price, so we can't really sort by name now. We will just do a small hacky thing
	// for stations and utilities. It's a bit ugly, but this can be tweaked.
	var utilityAdded = false
	var stationAdded = false
	// for this, we just want one representative property for each set to iterate over
	var currentSetColour string
	var oneOfEach []*PropertyDeed = nil
	for _, pd := range properties {
		if (*pd).Set != currentSetColour {
			if pd.Set == GetPropertyType(Utility) && utilityAdded {
				continue
			}
			if pd.Set == GetPropertyType(Station) && stationAdded {
				stationAdded = true
				continue
			}
			currentSetColour = pd.Set
			if pd.Set == GetPropertyType(Utility) {
				utilityAdded = true
			}
			if pd.Set == GetPropertyType(Station) {
				stationAdded = true
			}
			oneOfEach = append(oneOfEach, pd)
		} else {
			continue
		}
	}
	//
	for _, pd := range oneOfEach {
		fullyOwned := true
		// we only need to do one of each colour
		owners, _ := ownersOfASet(pd.Set, pc)
		// check we own all of them
		for _, owner := range owners {
			if owner != properties[0].Owner { // we just need any card to establish our player number
				fullyOwned = false
			}
		}
		if fullyOwned {

			setsOwned = append(setsOwned, pd.Set)
		}
	}
	return setsOwned
}

func removeProperties(setColor string, propertiesToGiveOut []*PropertyDeed) arrayOfPropertyDeed {
	var noFullSets []*PropertyDeed
	for _, deed := range propertiesToGiveOut {
		if deed.Set == setColor {
			// don't add
		} else {
			noFullSets = append(noFullSets, deed)
		}
	}
	return noFullSets
}
