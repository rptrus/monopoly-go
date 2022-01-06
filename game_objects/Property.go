package game_objects

import (
	"errors"
	"fmt"
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

func (pd *PropertyDeed) PayRent(from *Player, to *Player, board *Board, props *PropertyCollection) (int, error) {
	if (*from).PlayerNumber == (*to).PlayerNumber {
		fmt.Println("Don't pay rent to ourselves")
		return 0, errors.New("RentToOurself") // *not really* an error, but a way to suppress output
	}
	if (*from).CashAvailable-pd.Rent < 0 {
		str := []string{"Player does not have enough funds to cover rent ", string(pd.Rent)}
		return 0, errors.New(strings.Join(str, " "))
	}
	switch board.MonopolySpace[from.PositionOnBoard].SquareType {
	case Utility:
		ownsBoth := len(findSameType(board, pd, props)) == 2
		roll := rollDice()
		fmt.Println("Utility re-roll of", roll)
		if ownsBoth {
			pd.Rent = 10 * roll
		} else {
			pd.Rent = 4 * roll
		}
		from.CashAvailable -= pd.Rent
		to.CashAvailable += pd.Rent
	case Station:
		stationsOwnedByPlayer := len(findSameType(board, pd, props))
		pd.Rent = stationsOwnedByPlayer * 25
		from.CashAvailable -= pd.Rent
		to.CashAvailable += pd.Rent
	case BuildableProperty:
		// check if the property landed on is a complete set
		multiplyFactor := 1
		hasAllSet := checkCompleteSet(pd, props)
		if hasAllSet {
			multiplyFactor = 2
		}
		from.CashAvailable -= pd.Rent * multiplyFactor
		to.CashAvailable += pd.Rent * multiplyFactor
	default:
		fmt.Println("Unknown or not implemented", board.MonopolySpace[from.PositionOnBoard].SquareType)
	}
	return pd.Rent, nil
}

func swapPropertyBetweenPlayers(from *Player, to *Player, card *PropertyDeed, myPropertyCardCollection *PropertyCollection) {
	fmt.Println("Player", from.Name, "Will give property", GetTheCurrentCardName(card.PositionOnBoard, myPropertyCardCollection), "to", to.Name)
	// since we will be swapping properties later, we don't need to adjust cash here
	//from.CashAvailable -= card.PurchaseCost
	//to.CashAvailable += card.PurchaseCost
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
	/*
		for _, property := range pc.AllProperty {
			for _, v := range property.Card {
				if v.Set == pd.Set { // same colour as our input property deed
					if v.Owner != pd.Owner {
						ownsAll = false
					}
				}
			}
		}
	*/
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
func ShowPropertiesOfPlayer(playerNumber int, myPropertyCardCollection *PropertyCollection) ([]string, []*PropertyDeed) {
	propsOwnedNameOnly := []string{}
	propDeeds := []*PropertyDeed{}
	for _, card := range myPropertyCardCollection.AllProperty {
		aSingularCardMap := card.Card
		for _, v := range aSingularCardMap {
			if int(v.Owner) == playerNumber {
				n, pd := GetTheCurrentCard(v.PositionOnBoard, myPropertyCardCollection)
				propsOwnedNameOnly = append(propsOwnedNameOnly, n)
				propDeeds = append(propDeeds, pd)
			}
		}
	}
	return propsOwnedNameOnly, propDeeds
}

func (gs *GameState) UnownedProperties(myPropertyCardCollection *PropertyCollection) {
	var propsSpare []string
	for _, props := range myPropertyCardCollection.AllProperty {
		for _, k := range props.Card { // 1 element map
			if k.Owner == 'u' {
				propsSpare = append(propsSpare, GetTheCurrentCardName(k.PositionOnBoard, myPropertyCardCollection))
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
func highestPartiallyCompleteSet(otherPlayer byte, AllPlayers []Player, myPropertyCardCollection *PropertyCollection) []*PropertyDeed {
	var setsWithMostPropertiesOwned []*PropertyDeed
	_, deeds := ShowPropertiesOfPlayer(int(otherPlayer), myPropertyCardCollection)
	for _, pd := range deeds {
		owned, totalInSet := propsOwnedByPlayerInASet(pd, myPropertyCardCollection)
		if len(owned) == 2 && totalInSet == 3 || len(owned) == 1 && totalInSet == 2 {
			setsWithMostPropertiesOwned = append(setsWithMostPropertiesOwned, pd)
		}
	}
	return setsWithMostPropertiesOwned
}

// needs work for things like utility / train station. works ok for coloured property sets
func ownsFullSet(propertiesToGiveOut []*PropertyDeed, myPropertyCardCollection *PropertyCollection) []string {
	var setsOwned []string
	// for this, we just want one representative property for each set to iterate over
	var currentSetColour string
	var oneOfEach []*PropertyDeed = nil
	for _, pd := range propertiesToGiveOut {
		if (*pd).Set != currentSetColour {
			currentSetColour = pd.Set
			oneOfEach = append(oneOfEach, pd)
		} else {
			continue
		} // check utility. it's not contiguous like the colours.
	}
	//
	for _, pd := range oneOfEach {
		fullyOwned := true
		// we only need to do one of each colour
		owners, _ := ownersOfASet(pd.Set, myPropertyCardCollection)
		// check we own all of them
		for _, owner := range owners {
			if owner != propertiesToGiveOut[0].Owner { // we just need any card to establish our player number
				fullyOwned = false
			}
		}
		if fullyOwned {

			setsOwned = append(setsOwned, pd.Set)
		}
	}
	return setsOwned
}

func removeProperties(setColor string, propertiesToGiveOut []*PropertyDeed) []*PropertyDeed {
	//var start int = 0
	//var end int = 0
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
