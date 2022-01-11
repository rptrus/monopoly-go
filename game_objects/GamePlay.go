package game_objects

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	maxdicelow          int = 1
	maxdicehigh         int = 6
	placesonboard       int = 40
	totalPlayersPlaying int = 6
	minThresholdHouses  int = 500
)

var TheBank *Bank

type GameState struct {
	CurrentPlayer         *Player
	CurrentPropertyOfTurn *Property
	CurrentDiceRoll       int
	GlobalTurnsMade       int
	allPropsSold          bool
	AllPlayers            []Player
	AllProperties         *PropertyCollection
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

// substantially similar code reuse (I know) but I'd rather keep both for convenience, mmkay?
func rollToGetOutOfJail() (int, int) {
	time.Sleep(1 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	first := maxdicelow + rand.Intn(maxdicehigh-maxdicelow+1)
	second := maxdicelow + rand.Intn(maxdicehigh-maxdicelow+1)
	return first, second
}

func (gs *GameState) RollDice() {
	gs.GlobalTurnsMade++
	gs.CurrentDiceRoll = rollDice()
}

// true: if game has been won by a player
func (gs *GameState) NextPlayer() bool {
	//gs.CurrentPlayer = &gs.AllPlayers[(gs.CurrentPlayer.PlayerNumber+1)%totalPlayersPlaying]
	var countActive int
	for _, j := range gs.AllPlayers {
		if j.Active == true {
			countActive++
		}
	}

	switch countActive {
	case 1:
		gs.CurrentPlayer = &gs.AllPlayers[(gs.CurrentPlayer.PlayerNumber+1)%totalPlayersPlaying]
		if gs.CurrentPlayer.Active == true {
			break
		}
		fallthrough
	case 2:
		gs.CurrentPlayer = &gs.AllPlayers[(gs.CurrentPlayer.PlayerNumber+1)%totalPlayersPlaying]
		if gs.CurrentPlayer.Active == true {
			break
		}
		fallthrough
	case 3:
		gs.CurrentPlayer = &gs.AllPlayers[(gs.CurrentPlayer.PlayerNumber+1)%totalPlayersPlaying]
		if gs.CurrentPlayer.Active == true {
			break
		}
		fallthrough
	case 4:
		gs.CurrentPlayer = &gs.AllPlayers[(gs.CurrentPlayer.PlayerNumber+1)%totalPlayersPlaying]
		if gs.CurrentPlayer.Active == true {
			break
		}
		fallthrough
	case 5:
		gs.CurrentPlayer = &gs.AllPlayers[(gs.CurrentPlayer.PlayerNumber+1)%totalPlayersPlaying]
		if gs.CurrentPlayer.Active == true {
			break
		}
		fallthrough
	default:
		gs.CurrentPlayer = &gs.AllPlayers[(gs.CurrentPlayer.PlayerNumber+1)%totalPlayersPlaying]
		if gs.CurrentPlayer.Active == true {
			break
		}
	}
	fmt.Println(gs.CurrentPlayer.Name, "is now up")

	if countActive == 1 {
		fmt.Println("Player", gs.CurrentPlayer.Name, "has won the game!")
		return true
	}
	return false
}

// board position to property[28]
func GetTheCurrentCard(board int, pc *PropertyCollection) (string, *PropertyDeed) {
	for _, card := range (*pc).AllProperty {
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

// convenience if we just want the name, we can use directly in a fmt.println statement
func GetTheCurrentCardName(board int, pc *PropertyCollection) string {
	name, _ := GetTheCurrentCard(board, pc)
	return name
}

func (gs *GameState) ProcessNonPropertySquare(CurrentPlayer *Player, sqType int, tax int) {
	taxCollection := 0
	switch sqType {
	case Tax:
		taxCollection += tax
		t := Transaction{
			//gs: gs,
			sender:   CurrentPlayer,
			receiver: nil,
			amount:   tax,
		}
		// general tax need 200
		if CurrentPlayer.PositionOnBoard == 4 {
			t.amount += tax
		}
		t.TransactWithBank()
		fmt.Println("Collected Tax: $", t.amount)
	case Jail:
		if CurrentPlayer.PositionOnBoard == 30 {
			CurrentPlayer.PositionOnBoard = 10
			CurrentPlayer.JailTurns = 3
		}
	default:
	}
}

func (gs *GameState) RemoveToken(playerToRemove *Player) {
	fmt.Println("Removing token", playerToRemove.Token, "played by", playerToRemove.Name)
	//var playerToExterminate int
	for _, j := range gs.AllPlayers {
		if j.PlayerNumber == playerToRemove.PlayerNumber {
			//playerToExterminate = i
			j.Active = false
			break
		}
	}
}
