package main

import (
	"fmt"
	"github.com/rptrus/monopoly-go/game_objects"
	"github.com/rptrus/monopoly-go/setup"
)

const numberOfPlayers int = 6

func main() {
	fmt.Println("Starting Monopoly Go SIM")

	board := setup.InitializeBoard()
	_ = setup.InitializeBank()
	propertyCardCollection := setup.InitializePropertyCards()
	allPlayers := setup.InitializePlayers(numberOfPlayers)
	firstUp := game_objects.RollToSeeWhoGoesFirst(allPlayers)
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
		theDeed := getTheCurrentCard(gameState.CurrentPlayer.PositionOnBoard, propertyCardCollection)
		fmt.Println("Landed on space ", gameState.CurrentPlayer.PositionOnBoard)
		if theDeed != nil {
			if theDeed.Owner == 'u' {
				fmt.Println("Purchase by player ", gameState.CurrentPlayer.PlayerNumber)
				gameState.CurrentPlayer.BuyProperty(theDeed)
			} else {
				fmt.Println("Pay rent")
			}
		} else {
			// TODO: print out the square type by english name not number
			sqType := board.MonopolySpace[gameState.CurrentPlayer.PositionOnBoard]
			fmt.Println("Landed on a non property square! ", sqType.SquareType)
		}
		// ...
		gameState.NextPlayer()

		if gameState.GlobalTurnsMade > 50 {
			break
		}
	}

	fmt.Println("Finish.")
}

// board position to property[28]
func getTheCurrentCard(board int, MyPropertyCardCollection *game_objects.PropertyCollection) *game_objects.PropertyDeed {
	for _, card := range (*MyPropertyCardCollection).AllProperty {
		// j is single entry map of name:Property
		aSingularCardMap := card.Card
		for _, v := range aSingularCardMap {
			//fmt.Println(v)
			if v.PositionOnBoard == board {
				fmt.Println("found it")
				return v
			} // do something with error
		}
	}
	fmt.Println("not found!!")
	return nil
}
