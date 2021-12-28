package game_objects

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	maxlow  int = 1
	maxhigh int = 6
)

type GameState struct {
	CurrentPlayer   *Player
	CurrentDiceRoll int
	GlobalTurnsMade int
}

// TODO: tie situation. At the moment first player with highest score wins the toss
func RollToSeeWhoGoesFirst() *Player {
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
	first := maxlow + rand.Intn(maxhigh-maxlow+1)
	second := maxlow + rand.Intn(maxhigh-maxlow+1)
	total := first + second
	return total
}

// acts as though it's a method on a class
func (gs *GameState) RollDice() {
	gs.GlobalTurnsMade++
	gs.CurrentDiceRoll = rollDice()
}

// For now we assume that there will be 6 players always
func (gs *GameState) NextPlayer() {
	gs.CurrentPlayer.PlayerNumber = (gs.CurrentPlayer.PlayerNumber + 1) % 6
}
