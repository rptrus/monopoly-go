package game_objects

import (
	"errors"
	"fmt"
	"github.com/rptrus/monopoly-go/utils"
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
	Mortgaged       bool
}

const cashBufferThreshold = 600

var ErrR2O = errors.New("RentToOurself")

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
	fmt.Println("6. Charge rent if applicable")
	if !(pd.Set == "Train" || pd.Set == "Utility") && pd.HousesOwned > 0 {
		fmt.Println("Cost of landing on builtup property:", GetTheCurrentCardName(pd.PositionOnBoard, BankGameState), "is: $", pd.RentWithHouses[pd.HousesOwned-1])
	}
	if pd.Mortgaged {
		return 0, nil
	}
	t := Transaction{
		Sender:   from,
		Receiver: to,
		Amount:   0,
	}
	var (
		actualPaid       = 0 // a player may not be able to pay the full amount
		err        error = nil
	)
	if (*from).PlayerNumber == (*to).PlayerNumber {
		fmt.Println("Don't pay rent to ourselves")
		return 0, ErrR2O // *not really* an error, but a way to suppress output
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
		t.Amount = pd.Rent
		actualPaid, err = t.TransactWithPlayer('x')
	case Station:
		stationsOwnedByPlayer := len(findSameType(board, pd, pc))
		pd.Rent = stationsOwnedByPlayer * 25
		t.Amount = pd.Rent
		actualPaid, err = t.TransactWithPlayer('x')
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
		t.Amount = moneyOwing * multiplyFactor
		actualPaid, err = t.TransactWithPlayer('x')
	default:
		fmt.Println("Unknown or not implemented", board.MonopolySpace[from.PositionOnBoard].SquareType)
	}
	var addition = ""
	if !(pd.Set == "Train" || pd.Set == "Utility") {
		addition = "with " + strconv.Itoa(pd.HousesOwned) + " houses"
	}
	fmt.Println("Invoice for landing", GetTheCurrentCardName(pd.PositionOnBoard, BankGameState), "is: $", t.Amount, addition)
	if err == nil {
		err = errors.New("Full-Payment")
	} // not really an error, but more like a status
	return actualPaid, err
}

func swapPropertyBetweenPlayers(from *Player, to *Player, card *PropertyDeed, pc *PropertyCollection) {
	fmt.Println("Player", from.Name, "Will give property", GetTheCurrentCardName(card.PositionOnBoard, BankGameState), "to", to.Name)
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
func ShowPropertiesOfPlayer(playerNumber int, gs *GameState) ([]string, arrayOfPropertyDeed) {
	var M = ""
	propsOwnedNameOnly := []string{}
	propDeeds := []*PropertyDeed{}
	for _, card := range gs.AllProperties.AllProperty {
		aSingularCardMap := card.Card
		for _, v := range aSingularCardMap {
			if int(v.Owner) == playerNumber {
				n, pd := GetTheCurrentCard(v.PositionOnBoard, gs)
				if pd.Mortgaged == true {
					M = " (M)"
				}
				propsOwnedNameOnly = append(propsOwnedNameOnly, n+M)
				propDeeds = append(propDeeds, pd)
			}
		}
	}
	return propsOwnedNameOnly, propDeeds
}

// convenience that we can use this in function calls without needing a preceeding variable first
func ShowPropertyDeedsOfPlayer(playerNumber int, gs *GameState) arrayOfPropertyDeed {
	_, properties := ShowPropertiesOfPlayer(playerNumber, gs)
	return properties
}

func (gs *GameState) UnownedProperties(pc *PropertyCollection) {
	var propsSpare []string
	for _, props := range pc.AllProperty {
		for _, k := range props.Card { // 1 element map
			if k.Owner == 'u' {
				propsSpare = append(propsSpare, GetTheCurrentCardName(k.PositionOnBoard, gs))
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
func highestPartiallyCompleteSet(otherPlayer byte, gs *GameState) []*PropertyDeed {
	var setsWithMostPropertiesOwned []*PropertyDeed
	_, deeds := ShowPropertiesOfPlayer(int(otherPlayer), gs)
	for _, pd := range deeds {
		owned, totalInSet := propsOwnedByPlayerInASet(pd, gs.AllProperties)
		if len(owned) == 2 && totalInSet == 3 || len(owned) == 1 && totalInSet == 2 {
			setsWithMostPropertiesOwned = append(setsWithMostPropertiesOwned, pd)
		}
	}
	return setsWithMostPropertiesOwned
}

// needs work for things like utility / train station. works ok for coloured property sets
func ownsFullSet(properties arrayOfPropertyDeed, pc *PropertyCollection) []string {
	var setsOwned []string
	for _, pd := range properties {
		fullyOwned := true
		// we technically only need to do one of each colour, but we are checking others of a set again unnecessarily (tradeoff is not large enough to warrant concern and keep code simpler)
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
	setsOwned = utils.RemoveDuplicateStr(setsOwned)
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

func AcquireAllMortgagedProperties(playerToAcquire *Player, playerToRecoverFrom *Player) {
	_, props := ShowPropertiesOfPlayer(playerToRecoverFrom.PlayerNumber, BankGameState)
	// everything should be mortgaged at this point since we tried to mortgage everything prior to paying debts
	for _, prop := range props {
		fmt.Println("Player", playerToAcquire.Name, "has just acquired", GetTheCurrentCardName(prop.PositionOnBoard, BankGameState)+" (", prop.Set, ") ")
		//unMortgageOptions(playerToAcquire, prop)
		if playerToAcquire.CashAvailable >= 600 { // 600 just an aribtrary chosen buffer of cash to keep
			unMortgageCost := int(float64(prop.PurchaseCost) * half * (1 + tenPercent))
			t := Transaction{
				Sender:   playerToAcquire,
				Receiver: nil,
				Amount:   unMortgageCost,
			}
			t.TransactWithBank()
			fmt.Println("Unmortgage (full) cost:", unMortgageCost)
			prop.Owner = byte(playerToAcquire.PlayerNumber)
			prop.Mortgaged = false
		} else {
			unMortgageCost := int(float64(prop.PurchaseCost) * half * tenPercent)
			t := Transaction{
				Sender:   playerToAcquire,
				Receiver: nil,
				Amount:   unMortgageCost,
			}
			t.TransactWithBank()
			fmt.Println("Unmortgage (partial) cost:", unMortgageCost)
			prop.Owner = byte(playerToAcquire.PlayerNumber)
			prop.Mortgaged = true // redundant, just to document
		}
	}
}
