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

// input: player number
// output: properties owned
func (gs *GameState) ShowPropertiesOfPlayer(playerNumber int, myPropertyCardCollection *PropertyCollection) []string {
	propsOwned := []string{}
	for _, card := range myPropertyCardCollection.AllProperty {
		aSingularCardMap := card.Card
		for _, v := range aSingularCardMap {
			if int(v.Owner) == /*gs.CurrentPlayer.PlayerNumber*/ playerNumber {
				n, _ := GetTheCurrentCard(v.PositionOnBoard, myPropertyCardCollection)
				propsOwned = append(propsOwned, n)
			}
		}
	}
	return propsOwned
}

func (gs *GameState) DoDeals(allPlayers []Player, myPropertyCardCollection *PropertyCollection) {
	// TODO
	var propsOwned []string
	fmt.Println("These other players own the following properties:")
	for _, j := range allPlayers {
		if j.PlayerNumber != gs.CurrentPlayer.PlayerNumber {
			propsOwned = gs.ShowPropertiesOfPlayer(j.PlayerNumber, myPropertyCardCollection)
			fmt.Print("[", j.Name, "-> \"", strings.Join(propsOwned, "\",\""), "\"]\n")
		}
	}
	gs.UnownedProperties(myPropertyCardCollection)
}

func (gs *GameState) UnownedProperties(myPropertyCardCollection *PropertyCollection) {
	var propsSpare []string
	for _, props := range myPropertyCardCollection.AllProperty {
		for _, k := range props.Card { // 1 element map
			if k.Owner == 'u' {
				name, _ := GetTheCurrentCard(k.PositionOnBoard, myPropertyCardCollection)
				propsSpare = append(propsSpare, name)
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
