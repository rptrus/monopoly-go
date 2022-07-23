package game_objects

import (
	"errors"
	"fmt"
)

// This file uses receiver methods for structs

const (
	roundTripPayment = 200
)

// board position will be zero based. GO space is zero
type Player struct {
	PlayerNumber    int
	Name            string
	CashAvailable   int
	PositionOnBoard int
	Active          bool
	Turns           int
	JailTurns       int
	JailCards       []byte
	Token           string
}

// return reward if GO is passed, 0 otherwise. If return results need to be augmented will create a struct in future
func (p *Player) AdvancePlayer(steps int, cc *CardCollection) int {
	p.Turns++
	prePosition := p.PositionOnBoard
	if p.JailTurns == 0 {
		p.PositionOnBoard += steps
	} else if p.JailTurns > 0 && len(p.JailCards) > 0 {
		p.JailTurns = 0
		backIntoStack := p.JailCards[0]
		if backIntoStack == 'H' {
			cc.ShuffleOrderH = append(cc.ShuffleOrderH, 15) // 15 is chance card for jail free
		} else {
			cc.ShuffleOrderO = append(cc.ShuffleOrderO, 11) // 11 is community chest card for jail free
		}
		p.JailCards = p.JailCards[1:]
		p.PositionOnBoard += steps
	} else {
		firstRoll, secondRoll := rollToGetOutOfJail()
		if firstRoll == secondRoll {
			fmt.Println("Rolled a double! lets get out of Jail")
			p.JailTurns = 0
			p.PositionOnBoard += firstRoll + secondRoll
		} else {
			if p.JailTurns == 1 {
				fmt.Println("Exhausted all rolls, pay $50 to get out and roll", firstRoll+secondRoll, "spaces")
				t := Transaction{
					Sender:   p,
					Receiver: nil,
					Amount:   50,
				}
				t.TransactWithBank()
				p.JailTurns = 0
				p.PositionOnBoard += firstRoll + secondRoll
			} else {
				p.JailTurns--
				p.PositionOnBoard += 0
				fmt.Println("Rolled a ", firstRoll, "and", secondRoll, ". Not succesful.", p.JailTurns, "more tries available")
			}
		}
	}
	p.PositionOnBoard = p.PositionOnBoard % placesonboard
	if p.PositionOnBoard < prePosition && p.JailTurns == 0 {
		p.pay200Dollars()
		return roundTripPayment
	}
	return 0
}

func (p *Player) BuyProperty(pd *PropertyDeed) (int, error) {
	fmt.Println("5. Buy property if available")
	if p.CashAvailable-pd.PurchaseCost < 0 {
		return 0, errors.New("Cannot afford property!")
	}
	t := Transaction{
		Sender:   p,
		Receiver: nil,
		Amount:   (*pd).PurchaseCost,
	}
	t.TransactWithBank()
	pd.Owner = byte(p.PlayerNumber)
	return pd.PurchaseCost, nil
}

func (p *Player) pay200Dollars() {
	t := Transaction{
		Sender:   nil,
		Receiver: p,
		Amount:   200,
	}
	t.BankCheque()
}

func (p *Player) PutUpHouses(gs *GameState) {
	fmt.Println("3. Put up houses if possible")
	deeds := ShowPropertyDeedsOfPlayer(p.PlayerNumber, gs)
	colour := ownsFullSet(deeds, gs.AllProperties)
	// buy houses of these colours, 1 lot at a time
	maxSoFar := 5 // start with MAX:5 (hotel assumption) and work down to the lowest housed property
	nextPropertyToBuildInColourSet := -1
	for _, aFullSetColour := range colour {
		for _, deed := range deeds {
			if deed.Set == aFullSetColour {
				if deed.Set == "Utility" || deed.Set == "Train" {
					continue
				} // no houses for these
				if deed.HousesOwned >= 5 {
					continue
				}
				if deed.HousesOwned < maxSoFar {
					nextPropertyToBuildInColourSet = deed.PositionOnBoard
					maxSoFar = deed.HousesOwned
				}
			}
		}
	}
	if p.CashAvailable > minThresholdHouses && nextPropertyToBuildInColourSet != -1 {
		_, deed := GetTheCurrentCard(nextPropertyToBuildInColourSet, gs)
		t := Transaction{
			Sender:   p,
			Receiver: nil,
			Amount:   deed.HouseCost,
		}
		t.TransactWithBank()
		deed.HousesOwned++
		structureType := "House/s"
		structures := deed.HousesOwned
		if deed.HousesOwned == 5 {
			structureType = "Hotel"
			structures = 1
		}
		fmt.Println(structureType, "purchased for", GetTheCurrentCardName(deed.PositionOnBoard, gs), "by", p.Name, ". ", structureType, "on this property are: ", structures)
	}
}

// if we have mortgaged properties and the requisite cash, we can umortgage them and make them productive!
func (p *Player) CheckToUnmortgage(player *Player, pc arrayOfPropertyDeed) {
	for _, prop := range pc {
		if player.CashAvailable > cashBufferThreshold {
			if prop.Mortgaged == true {
				fmt.Println("1. Unmortgage check")
				t := Transaction{
					Sender:   player,
					Receiver: nil,
					Amount:   int(float64(prop.PurchaseCost) * tenPercent),
				}
				prop.Mortgaged = false
				fmt.Println("Unmortgaged: ", GetTheCurrentCardName(prop.PositionOnBoard, BankGameState))
				t.TransactWithBank()
			}
		}
	}
}
