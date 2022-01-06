package game_objects

import (
	"fmt"
	"strings"
)

func (gs *GameState) DoDeals(allPlayers []Player, myPropertyCardCollection *PropertyCollection) {
	// TODO
	// Show some helpful logging so we know the state of play
	var propsOwned []string
	fmt.Println("These other players own the following properties:")
	for i, j := range allPlayers {
		if j.PlayerNumber != gs.CurrentPlayer.PlayerNumber {
			propsOwned, _ = ShowPropertiesOfPlayer(j.PlayerNumber, myPropertyCardCollection)
			fmt.Print("[", j.Name, " (", i, ")-> \"", strings.Join(propsOwned, "\",\""), "\"]\n")
		}
	}
	gs.UnownedProperties(myPropertyCardCollection)
	// work out if we have anything that we (the current player) have anything viable to trade to the player we just got our card from
	_, propertyDeeds := ShowPropertiesOfPlayer(gs.CurrentPlayer.PlayerNumber, myPropertyCardCollection)
	for _, pd := range propertyDeeds {
		myCount, totalCount := propsOwnedByPlayerInASet(pd, myPropertyCardCollection)
		if (len(myCount) == 1 && totalCount == 2) || (len(myCount) == 2 && totalCount == 3) {
			// majority ownership in a 3 card set or half in a 2 card set
			name, cardToSwap := GetTheCurrentCard(pd.PositionOnBoard, myPropertyCardCollection)
			// find other player who owns the card so we can complete it, and make sure the bank owns none of them (all bought by players)
			owners, bank := ownersOfASet(pd.Set, myPropertyCardCollection)
			if bank == false {
				fmt.Println("Have a candidate here:", pd.Set, ":", name)
				fmt.Println("Owners (with no bank as owner): ", owners)
				// simply look for player who isn't our player number
				otherOwner := OtherOwnerOfSet(gs.CurrentPlayer.PlayerNumber, owners)
				_, propsAll := ShowPropertiesOfPlayer(int(otherOwner), myPropertyCardCollection)
				var otherCardNeeded *PropertyDeed = nil
				for _, listOfCardsOtherPlayer := range propsAll {
					if listOfCardsOtherPlayer.Set == cardToSwap.Set {
						otherCardNeeded = listOfCardsOtherPlayer
					}
				}
				fmt.Println("We will get the card", GetTheCurrentCardName(otherCardNeeded.PositionOnBoard, myPropertyCardCollection), "from", allPlayers[otherOwner].Name)
				// obtain what we need to fill this missing piece
				swapPropertyBetweenPlayers(&allPlayers[otherOwner], gs.CurrentPlayer, otherCardNeeded, myPropertyCardCollection)
				// now we have to give back to the swapper, preferably something they need
				// check which set the player has 2 or more of. If they are lucky enough, send them that card
				pdSetOfOtherPlayer := highestPartiallyCompleteSet(otherOwner, allPlayers, myPropertyCardCollection)
				// for each high available (generally 2+) set another player owns, check if we own it by cycling through the owners to see if we're there
				dealDone := false
			out:
				for _, pd2 := range pdSetOfOtherPlayer {
					owners, _ = ownersOfASet(pd2.Set, myPropertyCardCollection)
					for owner := range owners {
						if owner == gs.CurrentPlayer.PlayerNumber {
							// we are an owner of something they need, lets give it to them as a good steward
							// get the card by cycling through our cards with this colour
							_, propertyOfCurrentPlayer := ShowPropertiesOfPlayer(gs.CurrentPlayer.PlayerNumber, myPropertyCardCollection)
							for _, j := range propertyOfCurrentPlayer {
								if j.Set == pd2.Set { // colour of other player who will get the magic card to fill their set
									swapPropertyBetweenPlayers(gs.CurrentPlayer, &allPlayers[otherOwner], pd2, myPropertyCardCollection)
									dealDone = true
									break out
									// then we can swap this
								}
							}

						}
					}
				}
				if !dealDone {
					fmt.Println("Will do another deal. TBD.")
					// we can swap 2 random properties
					// if we don't have 2 properties, then swap 1 + $300
					// if we don't have any properties, then swap $700
					// if we don't have the money for this, we are out
				}
			}
		}
	}
}
