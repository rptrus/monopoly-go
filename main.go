package main

import (
	"fmt"
	"github.com/rptrus/monopoly-go/game_objects"
	"github.com/rptrus/monopoly-go/setup"
	"strconv"
)

const (
	numberOfPlayers int = 6
	numberOfTurns       = 800
)

func main() {
	fmt.Println("Starting Monopoly Go SIM")

	board := setup.InitializeBoard()
	_ = setup.InitializeBank()
	propertyCardCollection := setup.InitializePropertyCards()
	allPlayers := setup.InitializePlayers(numberOfPlayers)
	firstUp := game_objects.RollToSeeWhoGoesFirst(allPlayers)
	println(firstUp.Name, " is going first...")
	gameState := game_objects.GameState{
		CurrentPlayer:   firstUp,
		CurrentDiceRoll: 0,
		GlobalTurnsMade: 0,
	}
	for {
		gameState.RollDice()
		prePosition := gameState.CurrentPlayer.PositionOnBoard // before we roll to our next place
		gameState.CurrentPlayer.AdvancePlayer(gameState.CurrentDiceRoll)
		// do some monopoly stuff here
		fmt.Println("\nTurn:", gameState.GlobalTurnsMade, "Current Player", gameState.CurrentPlayer.PlayerNumber, "rolled a", gameState.CurrentDiceRoll)
		thePropertyName, theDeed := game_objects.GetTheCurrentCard(gameState.CurrentPlayer.PositionOnBoard, propertyCardCollection)
		if theDeed != nil {
			// currently a bug (or design shortfall) that it will only display property cards and not other non-property cards. Will create another array for non-property cards
			preName, _ := game_objects.GetTheCurrentCard(prePosition, propertyCardCollection)
			str := "Moved from " + strconv.Itoa(prePosition) + " " + preName + " and Landed on space " + strconv.Itoa(gameState.CurrentPlayer.PositionOnBoard) + " " + string(thePropertyName) + " Owned by "
			if theDeed.Owner != 'u' {
				fmt.Println(str+"Player", int(theDeed.Owner))
			} else {
				fmt.Println(str + "Bank")
			}
			if theDeed.Owner == 'u' {
				gameState.CurrentPlayer.BuyProperty(theDeed)
				fmt.Println("Purchase $", theDeed.PurchaseCost, "by player", gameState.CurrentPlayer.Name, "who now has $", gameState.CurrentPlayer.CashAvailable)
			} else {
				_, err := theDeed.PayRent(&allPlayers[gameState.CurrentPlayer.PlayerNumber], &allPlayers[int(theDeed.Owner)], board)
				if err == nil {
					fmt.Println(allPlayers[gameState.CurrentPlayer.PlayerNumber].Name, gameState.CurrentPlayer.PlayerNumber, "paid $", theDeed.Rent, "rent to Player", allPlayers[int(theDeed.Owner)].Name, int(theDeed.Owner))
					fmt.Println(allPlayers[gameState.CurrentPlayer.PlayerNumber].Name, "now has $", allPlayers[gameState.CurrentPlayer.PlayerNumber].CashAvailable, "and", allPlayers[int(theDeed.Owner)].Name, "has $", allPlayers[int(theDeed.Owner)].CashAvailable)
				}
			}
		} else {
			sqType := board.MonopolySpace[gameState.CurrentPlayer.PositionOnBoard]
			fmt.Println("Landed on a non property square!", gameState.CurrentPlayer.PositionOnBoard, game_objects.GetPropertyType(sqType.SquareType))
		}
		gameState.NextPlayer(allPlayers)

		if gameState.GlobalTurnsMade == numberOfTurns {
			break
		}
	}

	fmt.Println("Finish.")
}
