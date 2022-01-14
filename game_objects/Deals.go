package game_objects

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func (gs *GameState) DoDeals(pc *PropertyCollection) {
	// Show some helpful logging so we know the state of play
	fmt.Println("These other players own the following properties:")
	for i, j := range gs.AllPlayers {
		propNamesOwned, propDeeds := ShowPropertiesOfPlayer(j.PlayerNumber, pc)
		fullSetters := strings.Join(ownsFullSet(propDeeds, pc), " ")
		if j.PlayerNumber != gs.CurrentPlayer.PlayerNumber {
			if j.Active {
				fmt.Print("[", j.Name, " (", i, ")-> \"", strings.Join(propNamesOwned, "\",\""), "\"] Fullsets: "+fullSetters+" CASH: $", gs.AllPlayers[i].CashAvailable, gs.CurrentPlayer.Active, "\n")
			}
		} else {
			//if gs.CurrentPlayer.Active {
			fmt.Print("THIS TURN: [", j.Name, " (", i, ")-> \"", strings.Join(propNamesOwned, "\",\""), "\"] Fullsets: "+fullSetters+" CASH: $", gs.AllPlayers[i].CashAvailable, "\n")
			//}
		}
	}
	gs.UnownedProperties(pc) // needs to set AllPropsSold when applicable
	// work out if we have anything that we (the current player) have anything viable to trade to the player we just got our card from
	propertyDeeds := ShowPropertyDeedsOfPlayer(gs.CurrentPlayer.PlayerNumber, pc)
	for _, pd := range propertyDeeds {
		myCount, totalCount := propsOwnedByPlayerInASet(pd, pc)
		if (len(myCount) == 1 && totalCount == 2) || (len(myCount) == 2 && totalCount == 3) {
			// majority ownership in a 3 card set or half in a 2 card set
			name, cardToSwap := GetTheCurrentCard(pd.PositionOnBoard, pc)
			// find other player who owns the card so we can complete it, and make sure the bank owns none of them (all bought by players)
			owners, bank := ownersOfASet(pd.Set, pc)
			if bank == false {
				fmt.Println("Have a candidate here:", pd.Set, ":", name)
				fmt.Println("Owners (with no bank as owner): ", owners)
				// simply look for player who isn't our player number
				otherOwner := OtherOwnerOfSet(gs.CurrentPlayer.PlayerNumber, owners)
				propsAll := ShowPropertyDeedsOfPlayer(int(otherOwner), pc)
				var otherCardNeeded *PropertyDeed = nil
				for _, listOfCardsOtherPlayer := range propsAll {
					if listOfCardsOtherPlayer.Set == cardToSwap.Set {
						otherCardNeeded = listOfCardsOtherPlayer
					}
				}
				fmt.Println("We will get the card", GetTheCurrentCardName(otherCardNeeded.PositionOnBoard, pc), "from", gs.AllPlayers[otherOwner].Name)
				// obtain what we need to fill this missing piece
				swapPropertyBetweenPlayers(&gs.AllPlayers[otherOwner], gs.CurrentPlayer, otherCardNeeded, pc)
				// now we have to give back to the swapper, preferably something they need
				// check which set the player has 2 or more of. If they are lucky enough, send them that card
				pdSetOfOtherPlayer := highestPartiallyCompleteSet(otherOwner, gs.AllPlayers, pc)
				// for each high available (generally 2+) set another player owns, check if we own it by cycling through the owners to see if we're there
				dealDone := false
			out:
				for _, pd2 := range pdSetOfOtherPlayer {
					owners, _ = ownersOfASet(pd2.Set, pc)
					for owner := range owners {
						if owner == gs.CurrentPlayer.PlayerNumber {
							// we are an owner of something they need, lets give it to them as a good steward
							// get the card by cycling through our cards with this colour
							_, propertyOfCurrentPlayer := ShowPropertiesOfPlayer(gs.CurrentPlayer.PlayerNumber, pc)
							for _, j := range propertyOfCurrentPlayer {
								if j.Set == pd2.Set { // colour of other player who will get the magic card to fill their set
									swapPropertyBetweenPlayers(gs.CurrentPlayer, &gs.AllPlayers[otherOwner], pd2, pc)
									dealDone = true
									break out
								}
							}

						}
					}
				}
				// We can't give them the property they need. Will need to contend with giving them 2 of ours. One should be high value property.
				if !dealDone {
					fmt.Println("Will do another deal. TBD.")
					_, propertiesToGiveOut := ShowPropertiesOfPlayer(gs.CurrentPlayer.PlayerNumber, pc)
					// take out any full sets, we don't give those away
					fullSetsToTakeOut := ownsFullSet(propertiesToGiveOut, pc)
					for _, colour := range fullSetsToTakeOut {
						propertiesToGiveOut = removeProperties(colour, propertiesToGiveOut)
					}
					sort.Sort(propertiesToGiveOut) // highest starts at [0]
					leng := len(propertiesToGiveOut)
					time.Sleep(1 * time.Millisecond)
					rand.Seed(time.Now().UnixNano())
					var usedup int = -1
					if leng >= 2 {
						for i := 0; i < 2; i++ { // 2 choices of properties. There could be, say, 5 or 6 to choose from
							choice := rand.Intn(leng)
							if choice == usedup {
								choice = choice + 1%leng
							}
							if i == 0 {
								choice = 0
							} // first iteration we always give highest value card
							usedup = choice
							property := propertiesToGiveOut[choice]
							swapPropertyBetweenPlayers(gs.CurrentPlayer, &gs.AllPlayers[otherOwner], property, pc)
						}
					} else if leng == 1 {
						swapPropertyBetweenPlayers(gs.CurrentPlayer, &gs.AllPlayers[otherOwner], propertiesToGiveOut[0], pc)
						// make up for the shortfall
						t := Transaction{
							//gs: gs,
							Sender:   gs.CurrentPlayer,
							Receiver: &gs.AllPlayers[otherOwner],
							Amount:   300,
						}
						t.TransactWithPlayer('x')

					} else if leng == 0 {
						t := Transaction{
							//gs:       gs,
							Sender:   gs.CurrentPlayer,
							Receiver: &gs.AllPlayers[otherOwner],
							Amount:   700,
						}
						t.TransactWithPlayer('x')
					}
					// if we don't have 2 properties, then swap 1 + $300
					// if we don't have any properties, then swap $700
					// if we don't have the money for this, we are out
				}
			}
		}
	}
}
