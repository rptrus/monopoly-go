package main

import (
	"fmt"
	"github.com/rptrus/monopoly-go/game_objects"
	"github.com/rptrus/monopoly-go/setup"
	"strconv"
	"strings"
)

const (
	numberOfPlayers int = 6
	numberOfTurns       = 800
	tax                 = 100
)

func init() {
	fmt.Println("Starting Monopoly Go SIM")
}

func main() {
	// move to init function
	board := setup.InitializeBoard()
	game_objects.TheBank = setup.InitializeBank()
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
		prePosition := gameState.CurrentPlayer.PositionOnBoard // place before we advance to our roll
		passGoPayment := gameState.CurrentPlayer.AdvancePlayer(gameState.CurrentDiceRoll)
		// do some monopoly stuff here
		fmt.Println("\nTurn:", gameState.GlobalTurnsMade, "Current Player", gameState.CurrentPlayer.Name, gameState.CurrentPlayer.PlayerNumber, "rolled a", gameState.CurrentDiceRoll)
		if passGoPayment > 0 {
			fmt.Println("BANK PAYS PLAYER $", passGoPayment)
		}
		thePropertyName, theDeed := game_objects.GetTheCurrentCard(gameState.CurrentPlayer.PositionOnBoard, propertyCardCollection)
		if theDeed != nil {
			// currently a bug (or design shortfall) that it will only display property cards and not other non-property cards. Will create another array for non-property cards
			preName, _ := game_objects.GetTheCurrentCard(prePosition, propertyCardCollection)
			str := "Moved from space " + strconv.Itoa(prePosition) + " " + preName + " and Landed on space " + strconv.Itoa(gameState.CurrentPlayer.PositionOnBoard) + " " + string(thePropertyName) + " Owned by "
			if theDeed.Owner != 'u' {
				fmt.Println(str+"Player", int(theDeed.Owner))
			} else {
				fmt.Println(str+"Bank", theDeed.Owner)
			}
			if theDeed.Owner == 'u' {
				_, err := gameState.CurrentPlayer.BuyProperty(theDeed)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Purchase $", theDeed.PurchaseCost, "by player", gameState.CurrentPlayer.Name, "who now has $", gameState.CurrentPlayer.CashAvailable)
			} else {
				_, err := theDeed.PayRent(&allPlayers[gameState.CurrentPlayer.PlayerNumber], &allPlayers[int(theDeed.Owner)], board, propertyCardCollection)
				if err == nil { // no errors
					fmt.Println(allPlayers[gameState.CurrentPlayer.PlayerNumber].Name, gameState.CurrentPlayer.PlayerNumber, "paid $", theDeed.Rent, "rent to Player", allPlayers[int(theDeed.Owner)].Name, int(theDeed.Owner))
					fmt.Println(allPlayers[gameState.CurrentPlayer.PlayerNumber].Name, "now has $", allPlayers[gameState.CurrentPlayer.PlayerNumber].CashAvailable, "and", allPlayers[int(theDeed.Owner)].Name, "has $", allPlayers[int(theDeed.Owner)].CashAvailable)
				}
			}
			names, _ := game_objects.ShowPropertiesOfPlayer(gameState.CurrentPlayer.PlayerNumber, propertyCardCollection)
			fmt.Println("\n", allPlayers[gameState.CurrentPlayer.PlayerNumber].Name, "owns the following properties:", "[ \""+strings.Join(names, "\",\"")+"\" ]\n")

			gameState.DoDeals(allPlayers, propertyCardCollection)
		} else {
			sqType := board.MonopolySpace[gameState.CurrentPlayer.PositionOnBoard].SquareType
			fmt.Println("Landed on a non property square!", gameState.CurrentPlayer.PositionOnBoard, game_objects.GetPropertyType(sqType))
			gameState.ProcessNonPropertySquare(gameState.CurrentPlayer, sqType, tax)
		}
		gameState.NextPlayer(allPlayers)
		if gameState.GlobalTurnsMade == numberOfTurns {
			break
		}
	}
	fmt.Println("Finish.")
}
