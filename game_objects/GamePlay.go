package game_objects

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	maxdicelow          int = 1
	maxdicehigh         int = 6
	placesonboard       int = 40
	totalPlayersPlaying int = 6
)

var TheBank *Bank

type GameState struct {
	CurrentPlayer         *Player
	CurrentPropertyOfTurn *Property
	CurrentDiceRoll       int
	GlobalTurnsMade       int
	allPropsSold          bool
}

// TODO: tie situation. At the moment first player with highest score wins the toss
func RollToSeeWhoGoesFirst(AllPlayers []Player) *Player {
	// in the same directory / package so no need to qualify it
	var (
		highestSoFarPlayer int = 0
		highestSoFarScore  int = 0
	)

	for i, _ := range AllPlayers {
		total := rollDice()
		if total > highestSoFarScore {
			highestSoFarScore = total
			highestSoFarPlayer = i
		}
	}
	fmt.Println(highestSoFarPlayer, " wins the toss")
	//return highestSoFarPlayer
	return &AllPlayers[highestSoFarPlayer]
}

func rollDice() int {
	time.Sleep(1 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	first := maxdicelow + rand.Intn(maxdicehigh-maxdicelow+1)
	second := maxdicelow + rand.Intn(maxdicehigh-maxdicelow+1)
	total := first + second
	return total
}

// acts as though it's a method on a class
func (gs *GameState) RollDice() {
	gs.GlobalTurnsMade++
	gs.CurrentDiceRoll = rollDice()
}

// For now we assume that there will be 6 players always
func (gs *GameState) NextPlayer(allPlayers []Player) {
	//gs.CurrentPlayer.PlayerNumber = (gs.CurrentPlayer.PlayerNumber + 1) % totalPlayersPlaying
	gs.CurrentPlayer = &allPlayers[(gs.CurrentPlayer.PlayerNumber+1)%totalPlayersPlaying]
}

// board position to property[28]
func GetTheCurrentCard(board int, MyPropertyCardCollection *PropertyCollection) (string, *PropertyDeed) {
	for _, card := range (*MyPropertyCardCollection).AllProperty {
		// j is single entry map of name:Property
		aSingularCardMap := card.Card
		for n, v := range aSingularCardMap {
			//fmt.Println(v)
			if v.PositionOnBoard == board {
				return n, v
			}
		}
	}
	return "", nil
}

func (gs *GameState) DoDeals(allPlayers []Player, myPropertyCardCollection *PropertyCollection) {
	// TODO
	// Show some helpful logging so we know the state of play
	var propsOwned []string
	fmt.Println("These other players own the following properties:")
	for _, j := range allPlayers {
		if j.PlayerNumber != gs.CurrentPlayer.PlayerNumber {
			propsOwned, _ = gs.ShowPropertiesOfPlayer(j.PlayerNumber, myPropertyCardCollection)
			fmt.Print("[", j.Name, "-> \"", strings.Join(propsOwned, "\",\""), "\"]\n")
		}
	}
	gs.UnownedProperties(myPropertyCardCollection)
	// work out if we have anything that we (the current player) have anything viable to trade
	_, propertyDeeds := gs.ShowPropertiesOfPlayer(gs.CurrentPlayer.PlayerNumber, myPropertyCardCollection)
	for _, pd := range propertyDeeds {
		myCount, totalCount := propsOwnedByPlayerInASet(pd, myPropertyCardCollection)
		if (len(myCount) == 1 && totalCount == 2) || (len(myCount) == 2 && totalCount == 3) {
			// majority ownership in a 3 card set or half in a 2 card set
			name, _ := GetTheCurrentCard(pd.PositionOnBoard, myPropertyCardCollection)
			fmt.Println("Have a candidate here: ", pd.Set, ":", name)
		}

	}
}
