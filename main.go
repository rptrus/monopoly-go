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
		gameState.RollDice()
		gameState.CurrentPlayer.AdvancePlayer(gameState.CurrentDiceRoll)
		// do some monopoly stuff here
		fmt.Println("\nCurrent Player ", gameState.CurrentPlayer.PlayerNumber, " rolled a ", gameState.CurrentDiceRoll)
		theDeed := getTheCurrentCard(gameState.CurrentPlayer.PositionOnBoard, propertyCardCollection)
		fmt.Println("Landed on space ", gameState.CurrentPlayer.PositionOnBoard)
		if theDeed != nil {
			if theDeed.Owner == 'u' {
				fmt.Println("Purchase by player ", gameState.CurrentPlayer.Name, " who now has $", gameState.CurrentPlayer.CashAvailable)
				gameState.CurrentPlayer.BuyProperty(theDeed)
			} else {
				_, err := theDeed.PayRent(&allPlayers[gameState.CurrentPlayer.PlayerNumber], &allPlayers[int(theDeed.Owner)])
				if err != nil {
					fmt.Println(allPlayers[gameState.CurrentPlayer.PlayerNumber].Name, gameState.CurrentPlayer.PlayerNumber, " pay $", theDeed.Rent, " rent to Player ", allPlayers[int(theDeed.Owner)].Name, int(theDeed.Owner))
					fmt.Println(allPlayers[gameState.CurrentPlayer.PlayerNumber].Name, " now has $", allPlayers[gameState.CurrentPlayer.PlayerNumber].CashAvailable, " and ", allPlayers[int(theDeed.Owner)].Name, " has $", allPlayers[int(theDeed.Owner)].CashAvailable)
				}
			}
		} else {
			sqType := board.MonopolySpace[gameState.CurrentPlayer.PositionOnBoard]
			fmt.Println("Landed on a non property square! ", game_objects.GetPropertyType(sqType.SquareType))
		}
		// ...
		gameState.NextPlayer(allPlayers)

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
				return v
			} // do something with error
		}
	}
	return nil
}
