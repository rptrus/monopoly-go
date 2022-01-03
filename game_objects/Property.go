package game_objects

import (
	"errors"
	"fmt"
	"strings"
)

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
		// just implementing single utility for now
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
	case BuildableProperty:
		fallthrough
	default:
		from.CashAvailable -= pd.Rent
		to.CashAvailable += pd.Rent
	}
	return pd.Rent, nil
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
