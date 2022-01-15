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
	drawCards := setup.InitializeDrawCards()
	allPlayers := setup.InitializePlayers(numberOfPlayers)
	firstUp, score := game_objects.RollToSeeWhoGoesFirst(allPlayers)
	println(firstUp.Name, " is going first with score", score, "...")
	gameState := game_objects.GameState{
		CurrentPlayer:   firstUp,
		CurrentDiceRoll: 0,
		GlobalTurnsMade: 0,
		AllPlayers:      allPlayers,
		AllProperties:   propertyCardCollection,
	}
	game_objects.BankGameState = &gameState
	for {
		deedsOwned := game_objects.ShowPropertyDeedsOfPlayer(gameState.CurrentPlayer.PlayerNumber, gameState.AllProperties)
		gameState.CurrentPlayer.CheckToUnmortgage(gameState.CurrentPlayer, deedsOwned)
		gameState.DoDeals(gameState.AllProperties)
		gameState.CurrentPlayer.PutUpHouses(gameState.AllProperties)
		gameState.RollDice()
		prePosition := gameState.CurrentPlayer.PositionOnBoard // place before we advance to our roll
		passGoPayment := gameState.CurrentPlayer.AdvancePlayer(gameState.CurrentDiceRoll, drawCards)
		// do some monopoly stuff here
		fmt.Println("==============================================================================="+
			"\nTurn:", gameState.GlobalTurnsMade, "Current Player", gameState.CurrentPlayer.Name, gameState.CurrentPlayer.PlayerNumber, "rolled a", gameState.CurrentDiceRoll)
		if passGoPayment > 0 {
			fmt.Println("BANK PAYS PLAYER $", passGoPayment)
		}
		thePropertyName, theDeed := game_objects.GetTheCurrentCard(gameState.CurrentPlayer.PositionOnBoard, gameState.AllProperties)
		if theDeed != nil {
			preName, _ := game_objects.GetTheCurrentCard(prePosition, gameState.AllProperties)
			movedToStr := "Moved from space " + strconv.Itoa(prePosition) + " " + preName + " and Landed on space " + strconv.Itoa(gameState.CurrentPlayer.PositionOnBoard) + " " + string(thePropertyName) + " owned by "
			if theDeed.Owner == 'u' {
				fmt.Println(movedToStr+"Bank", theDeed.Owner)
				_, err := gameState.CurrentPlayer.BuyProperty(theDeed)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Purchase $", theDeed.PurchaseCost, "by player", gameState.CurrentPlayer.Name, "who now has $", gameState.CurrentPlayer.CashAvailable)
			} else {
				fmt.Println(movedToStr+"Player", int(theDeed.Owner))
				rent, err := theDeed.PayRent(&allPlayers[gameState.CurrentPlayer.PlayerNumber], &allPlayers[int(theDeed.Owner)], board, gameState.AllProperties)
				if err != game_objects.ErrR2O { // suppress rent to ourself messages
					fmt.Println(allPlayers[gameState.CurrentPlayer.PlayerNumber].Name, gameState.CurrentPlayer.PlayerNumber, "paid $", rent, "rent to Player", allPlayers[int(theDeed.Owner)].Name, int(theDeed.Owner), "(", err, ")")
					fmt.Println(allPlayers[gameState.CurrentPlayer.PlayerNumber].Name, "now has $", allPlayers[gameState.CurrentPlayer.PlayerNumber].CashAvailable, "and", allPlayers[int(theDeed.Owner)].Name, "has $", allPlayers[int(theDeed.Owner)].CashAvailable)
				}
			}
			names, _ := game_objects.ShowPropertiesOfPlayer(gameState.CurrentPlayer.PlayerNumber, gameState.AllProperties)
			fmt.Println("\n"+allPlayers[gameState.CurrentPlayer.PlayerNumber].Name, "owns the following properties:", "[ \""+strings.Join(names, "\",\"")+"\" ]\n")
		} else {
			sqType := board.MonopolySpace[gameState.CurrentPlayer.PositionOnBoard].SquareType
			fmt.Println("Landed on a non property square!", gameState.CurrentPlayer.PositionOnBoard, game_objects.GetPropertyType(sqType))
			gameState.ProcessNonPropertySquare(gameState.CurrentPlayer, sqType, tax, drawCards)
		}
		gameWon := gameState.NextPlayer()
		if gameWon == true || gameState.GlobalTurnsMade == numberOfTurns {
			break
		}
	}
	fmt.Println("Finish.", gameState.GlobalTurnsMade, "turns.")
}
