package game_objects

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func (gs *GameState) DoDeals(myPropertyCardCollection *PropertyCollection) {
	// Show some helpful logging so we know the state of play
	fmt.Println("These other players own the following properties:")
	for i, j := range gs.AllPlayers {
		propNamesOwned, propDeeds := ShowPropertiesOfPlayer(j.PlayerNumber, myPropertyCardCollection)
		fullSetters := strings.Join(ownsFullSet(propDeeds, myPropertyCardCollection), " ")
		if j.PlayerNumber != gs.CurrentPlayer.PlayerNumber {
			fmt.Print("[", j.Name, " (", i, ")-> \"", strings.Join(propNamesOwned, "\",\""), "\"] Fullsets: "+fullSetters+"\n")
		} else {
			fmt.Print("ME: [", j.Name, " (", i, ")-> \"", strings.Join(propNamesOwned, "\",\""), "\"] Fullsets: "+fullSetters+"\n")
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
				fmt.Println("We will get the card", GetTheCurrentCardName(otherCardNeeded.PositionOnBoard, myPropertyCardCollection), "from", gs.AllPlayers[otherOwner].Name)
				// obtain what we need to fill this missing piece
				swapPropertyBetweenPlayers(&gs.AllPlayers[otherOwner], gs.CurrentPlayer, otherCardNeeded, myPropertyCardCollection)
				// now we have to give back to the swapper, preferably something they need
				// check which set the player has 2 or more of. If they are lucky enough, send them that card
				pdSetOfOtherPlayer := highestPartiallyCompleteSet(otherOwner, gs.AllPlayers, myPropertyCardCollection)
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
									swapPropertyBetweenPlayers(gs.CurrentPlayer, &gs.AllPlayers[otherOwner], pd2, myPropertyCardCollection)
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
					// we can swap 2 random properties
					_, propertiesToGiveOut := ShowPropertiesOfPlayer(gs.CurrentPlayer.PlayerNumber, myPropertyCardCollection) // can we cache this? we use it a lot
					// take out any full sets, we don't give those away
					fullSetsToTakeOut := ownsFullSet(propertiesToGiveOut, myPropertyCardCollection)
					for _, colour := range fullSetsToTakeOut {
						propertiesToGiveOut = removeProperties(colour, propertiesToGiveOut)
					}
					sort.Sort(propertiesToGiveOut) // highest starts at [0]

					// do the removal from
					fmt.Println(fullSetsToTakeOut)
					// FUTURE TODO: sort them and give out the highest property
					leng := len(propertiesToGiveOut)
					fmt.Println(leng, propertiesToGiveOut) // useless
					time.Sleep(1 * time.Millisecond)
					rand.Seed(time.Now().UnixNano())
					var usedup int = -1
					if leng >= 2 {
						for i := 0; i < 2; i++ { // 2 choices of properties. There could be, say, 5 or 6 to choose from
							println(leng)
							choice := rand.Intn(leng)
							if choice == usedup {
								choice = choice + 1%leng
							}
							if i == 0 {
								choice = 0
							} // first iteration we always give highest value card
							usedup = choice
							property := propertiesToGiveOut[choice]
							swapPropertyBetweenPlayers(gs.CurrentPlayer, &gs.AllPlayers[otherOwner], property, myPropertyCardCollection)
						}
					} else if leng == 1 {
						swapPropertyBetweenPlayers(gs.CurrentPlayer, &gs.AllPlayers[otherOwner], propertiesToGiveOut[0], myPropertyCardCollection)
						// make up for the shortfall
						gs.CurrentPlayer.CashAvailable -= 300
						gs.AllPlayers[otherOwner].CashAvailable += 300
					} else if leng == 0 {
						// TODO: availability check
						gs.CurrentPlayer.CashAvailable -= 700
						gs.AllPlayers[otherOwner].CashAvailable += 700
					}
					//first := maxdicelow + rand.Intn(maxdicehigh-maxdicelow+1)
					//second := maxdicelow + rand.Intn(maxdicehigh-maxdicelow+1)
					// if we don't have 2 properties, then swap 1 + $300
					// if we don't have any properties, then swap $700
					// if we don't have the money for this, we are out
				}
			}
		}
	}
}
