package main

import (
	"fmt"
	"github.com/rptrus/monopoly-go/game_objects"
	"github.com/rptrus/monopoly-go/setup"
)

const numberOfPlayers int = 6

func main() {
	fmt.Println("Starting Monopoly Go SIM")

	a := []string{"a", "b", "c"}
	b := []string{"x", "y", "z"}
	a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
	//... when calling a function does the opposite: if you have several arguments in a slice, it will unpack them and pass as separate arguments to a variadic function.

	setup.InitializeBoard()
	setup.InitializeBank()
	setup.InitializePropertyCards()
	setup.InitializePlayers(numberOfPlayers)
	firstUp := game_objects.RollToSeeWhoGoesFirst()
	println(firstUp.PlayerNumber, " is going first...")
	gameState := game_objects.GameState{
		CurrentPlayer:   firstUp,
		CurrentDiceRoll: 0,
		GlobalTurnsMade: 0,
	}
	for {
		gameState.RollDice() // gs updated
		gameState.CurrentPlayer.AdvancePlayer(gameState.CurrentDiceRoll)
		// do some monopoly stuff here
		fmt.Println("Current Player ", gameState.CurrentPlayer.PlayerNumber, " rolled a ", gameState.CurrentDiceRoll)
		// ...
		gameState.NextPlayer()

		if gameState.GlobalTurnsMade > 50 {
			break
		}
	}

	fmt.Println("Finish.")
}
