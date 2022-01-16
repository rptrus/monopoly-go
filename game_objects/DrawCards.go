package game_objects

import (
	"fmt"
	"math/rand"
	"time"
)

type DrawCard struct {
	Id             int
	Designator     byte // 'H' = chance, 'O' = community chest
	Content        string
	PlayerToPlayer *Transaction
	PlayerToBank   *Transaction
	BankToPlayer   *Transaction
	MoveToSpace    *int // always check if they pass Go, except for Jail or Mayfair
	RelativeMove   *int // -3 = go back 3 spaces
	NearestType    *int
	PlayerPaysAll  *Transaction
	AllPaysPlayer  *Transaction
}

type CardCollection struct {
	AllDrawCards  [32]DrawCard
	ShuffleOrderH []int
	ShuffleOrderO []int
	CurrentCardH  int // Current card in Chance
	CurrentCardO  int // Current card in Community Chest
}

func GenerateOrderForChanceCommunityChestCards() []int {
	cardsToDeal := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(cardsToDeal)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cardsToDeal), func(i, j int) { cardsToDeal[i], cardsToDeal[j] = cardsToDeal[j], cardsToDeal[i] })
	fmt.Println(cardsToDeal)
	return cardsToDeal
}

func processDrawCardInternal(card *DrawCard, gs *GameState, cc *CardCollection) {
	if card.MoveToSpace != nil {
		fmt.Println("Move to space")
		if *card.MoveToSpace == 10 {
			gs.CurrentPlayer.JailTurns = 3
		} // special case
		gs.GoToSquare(*card.MoveToSpace, true)
	} else if card.RelativeMove != nil {
		fmt.Println("Relative move")
		gs.GoToSquare(gs.CurrentPlayer.PositionOnBoard-3, false)
	} else if card.NearestType != nil {
		fmt.Println("Move to nearest type")
		toPos := -1
		if *card.NearestType == Station {
			fmt.Println("Move to Train")
			currentPos := gs.CurrentPlayer.PositionOnBoard
			if gs.CurrentPlayer.PositionOnBoard == 36 {
				toPos = 5
			} else {
				for currentPos%5 != 0 || currentPos%10 == 0 {
					currentPos++
				}
				toPos = currentPos
			}
		} else if *card.NearestType == Utility {
			fmt.Println("Move to Utility")
			if gs.CurrentPlayer.PositionOnBoard > 12 && gs.CurrentPlayer.PositionOnBoard < 28 {
				toPos = 28
			} else {
				toPos = 12
			}
		} else {
			fmt.Errorf("Unknown type!", *card.NearestType)
		}
		gs.GoToSquare(toPos, true)
	} else if card.BankToPlayer != nil {
		fmt.Println("Bank pays player")
		card.BankToPlayer.Receiver = gs.CurrentPlayer
		card.BankToPlayer.BankCheque()
	} else if card.PlayerToBank != nil {
		fmt.Println("Player pays bank")
		card.PlayerToBank.Sender = gs.CurrentPlayer
		card.PlayerToBank.TransactWithBank()
	} else if card.PlayerToPlayer != nil {
		card.PlayerToPlayer.Sender = gs.CurrentPlayer
		card.PlayerToPlayer.TransactWithPlayer('x')
		fmt.Println("Player pays other players")
	} else if card.PlayerPaysAll != nil {
		card.PlayerPaysAll.Sender = gs.CurrentPlayer
		for i, j := range gs.AllPlayers {
			if j.PlayerNumber == gs.CurrentPlayer.PlayerNumber { // skip  ourself
			} else {
				if j.Active == true {
					card.PlayerPaysAll.Receiver = &gs.AllPlayers[i]
				}
			}
		}
		fmt.Println("Current player pays all other players")
	} else if card.AllPaysPlayer != nil {
		fmt.Println("All players pay current player")
		for i, j := range gs.AllPlayers {
			if j.PlayerNumber == gs.CurrentPlayer.PlayerNumber {
				continue
			}
			if gs.AllPlayers[i].Active == true {
				fmt.Println("Player", i, gs.AllPlayers[i])
				card.AllPaysPlayer.Sender = &gs.AllPlayers[i]
				card.AllPaysPlayer.Receiver = gs.CurrentPlayer
				card.AllPaysPlayer.TransactWithPlayer('x')
			}
		}
	} else {
		fmt.Println("More complex card processing", card.Content)
		var (
			countHousesOwned = 0
			countHotelsOwned = 0
			costHouseRepair  = 0
			costHotelRepair  = 0
		)

		repairs := (card.Designator == 'O' && card.Id == 8) || (card.Designator == 'H' && card.Id == 6)
		jailcard := (card.Designator == 'O' && card.Id == 11) || (card.Designator == 'H' && card.Id == 15)
		if repairs {
			deeds := ShowPropertyDeedsOfPlayer(gs.CurrentPlayer.PlayerNumber, gs.AllProperties)
			switch card.Id {
			case 8:
				costHouseRepair = 40
				costHotelRepair = 115
			case 6:
				costHouseRepair = 25
				costHotelRepair = 100
			}
			// street repairs 40 & 115
			for _, j := range deeds {
				if j.HousesOwned > 0 {
					if j.HousesOwned == 5 {
						countHotelsOwned++
					} else {
						countHousesOwned += j.HousesOwned
					}
				}
			}
			repairCost := (countHousesOwned * costHouseRepair) + (countHotelsOwned * costHotelRepair)
			if repairCost > 0 {
				t := Transaction{
					Sender: gs.CurrentPlayer,
					Amount: repairCost,
				}
				t.TransactWithBank()
			}
		}
		if jailcard {
			gs.CurrentPlayer.JailCards = append(gs.CurrentPlayer.JailCards, card.Designator)
			// remove card 11 from deck
			if card.Id == 11 {
				for i, j := range cc.ShuffleOrderH {
					if j == 11 {
						cc.ShuffleOrderH = append(cc.ShuffleOrderH[:i], cc.ShuffleOrderH[i+1:]...)
					}
				}
			}
			// remove card 15 from deck
			if card.Id == 15 {
				for i, j := range cc.ShuffleOrderO {
					if j == 15 {
						cc.ShuffleOrderH = append(cc.ShuffleOrderH[:i], cc.ShuffleOrderH[i+1:]...)
					}
				}
			}
		}
	}
}
