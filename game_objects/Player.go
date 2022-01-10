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
	JailTurns       int
	Token           string
}

// return reward if GO is passed, 0 otherwise. If return results need to be augmented will create a struct in future
func (p *Player) AdvancePlayer(steps int) int {
	prePosition := p.PositionOnBoard
	if p.JailTurns == 0 {
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
					sender:   p,
					receiver: nil,
					amount:   50,
				}
				t.TransactWithBank()
				//p.CashAvailable -= 50
				//TheBank.CashReservesInDollars += 50
				p.JailTurns = 0
				p.PositionOnBoard += firstRoll + secondRoll
			} else {
				p.JailTurns--
				p.PositionOnBoard += 0
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
	if p.CashAvailable-pd.PurchaseCost < 0 {
		return 0, errors.New("Cannot afford property!")
	}
	//(*p).CashAvailable -= (*pd).PurchaseCost
	t := Transaction{
		sender:   p,
		receiver: nil,
		amount:   (*pd).PurchaseCost,
	}
	t.TransactWithBank()
	pd.Owner = byte(p.PlayerNumber)
	return pd.PurchaseCost, nil
}

func (p *Player) pay200Dollars() {
	t := Transaction{
		sender:   nil,
		receiver: p,
		amount:   200,
	}
	t.BankCheque()
	//p.CashAvailable += roundTripPayment
	//TheBank.CashReservesInDollars -= 200
}

func (p *Player) PutUpHouses(pc *PropertyCollection) {
	_, deeds := ShowPropertiesOfPlayer(p.PlayerNumber, pc)
	colour := ownsFullSet(deeds, pc)
	// buy houses of these colours, 1 lot at a time
	for _, aFullSetColour := range colour {
		for _, deed := range deeds {
			if deed.Set == aFullSetColour {
				if deed.Set == "Utility" || deed.Set == "Train" {
					continue
				} // no houses for these
				if p.CashAvailable > minThresholdHouses {
					if deed.HousesOwned >= 5 {
						break
					}
					t := Transaction{
						sender:   p,
						receiver: nil,
						amount:   deed.HouseCost,
					}
					t.TransactWithBank()
					//p.CashAvailable -= deed.HouseCost
					//TheBank.CashReservesInDollars += deed.HouseCost
					deed.HousesOwned++
					fmt.Println("House purchased for", GetTheCurrentCardName(deed.PositionOnBoard, pc), "by", p.Name, ". Total houses on this property are: ", deed.HousesOwned)
				}
			}
		}
	}
}
