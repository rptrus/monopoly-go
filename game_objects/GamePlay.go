package game_objects

import (
	"errors"
	"fmt"
	"math/rand"
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
	AllPlayers            []Player
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

// acts as though it's a method on a class
func (gs *GameState) RollDice() {
	gs.GlobalTurnsMade++
	gs.CurrentDiceRoll = rollDice()
}

// For now we assume that there will be 6 players always
// TODO: no need to pass in the allPlayers. See if other instances like this....
func (gs *GameState) NextPlayer() {
	gs.CurrentPlayer = &gs.AllPlayers[(gs.CurrentPlayer.PlayerNumber+1)%totalPlayersPlaying]
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

// convenience if we just want the name, we can use directly in a fmt.println statement
func GetTheCurrentCardName(board int, MyPropertyCardCollection *PropertyCollection) string {
	name, _ := GetTheCurrentCard(board, MyPropertyCardCollection)
	return name
}

func (gs *GameState) ProcessNonPropertySquare(CurrentPlayer *Player, sqType int, tax int) {
	taxCollection := 0
	switch sqType {
	case Tax:
		taxCollection += tax
		TheBank.CashReservesInDollars += taxCollection
		CurrentPlayer.CashAvailable -= taxCollection
		// general tax need 200
		if CurrentPlayer.PositionOnBoard == 4 {
			taxCollection += tax
			TheBank.CashReservesInDollars += taxCollection
			CurrentPlayer.CashAvailable -= taxCollection
		}
		fmt.Println("Collected Tax: $", taxCollection)
	//implement Go To Jail
	case Jail:
		if CurrentPlayer.PositionOnBoard == 30 {
			CurrentPlayer.PositionOnBoard = 10
			CurrentPlayer.JailTurns = 3
		}
	default:
	}
}

// 'n' means it will just not allow, 'x' will mean that if the payer can't pay they will go bankrupt
func transact(from *Player, to *Player, fromAmount int, priority byte) error {
	var err error = nil
	if from.CashAvailable-fromAmount < 0 {
		if priority == 'n' {
			err = errors.New("Insufficient cash!")
			return err
		} else if priority == 'x' { // on reciprocal deals we don't muck around, you gotta have it
			from.CashAvailable -= from.CashAvailable
			to.CashAvailable += from.CashAvailable
			from.Active = false
		}
	} else {
		from.CashAvailable -= fromAmount
		to.CashAvailable += fromAmount
	}
	return nil
}

func (gs *GameState) removeToken(playerToRemove *Player) {
	// something to do with AllPlayers.
	// TEST THIS
	newPlayerList := append(gs.AllPlayers[:playerToRemove.PlayerNumber], gs.AllPlayers[playerToRemove.PlayerNumber+1:]...)
	fmt.Println(newPlayerList)
}
