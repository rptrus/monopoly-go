package main

import (
	"fmt"
	"github.com/rptrus/monopoly-go/game_objects"
	"github.com/rptrus/monopoly-go/setup"
	"os"
	"strconv"
)

const (
	numberOfTurns = 800
	tax           = 100
)

var numberOfPlayers int = 6

func init() {
	fmt.Println("Starting Monopoly Go SIM")
	if len(os.Args) == 2 {
		numberOfPlayers, _ = strconv.Atoi(os.Args[1])
	}
}

func main() {
	board := setup.InitializeBoard()
	game_objects.TheBank = setup.InitializeBank()
	propertyCardCollection := setup.InitializePropertyCards()
	otherCardCollection := setup.InitializeNonPropertyCards()
	drawCards := setup.InitializeDrawCards()
	globalToken := "Canon"
	allPlayers := setup.InitializePlayers(numberOfPlayers, globalToken)
	firstUp, _ := game_objects.RollToSeeWhoGoesFirst(allPlayers, globalToken)
	globState := setup.InitGlobalState(globalToken)
	gameState := setup.InitGameState(firstUp, globState, allPlayers, propertyCardCollection, otherCardCollection)
	game_objects.BankGameState = &gameState
	for {
		fmt.Println("\n================================================================================================"+
			"\n[ COUNTER:", gameState.Globals.GlobalTurnsMade, "] Turn:", gameState.CurrentPlayer.Turns, "for Current Player", gameState.CurrentPlayer.PlayerNumber, "(", gameState.CurrentPlayer.Name, ") currently on", game_objects.GetTheCurrentCardName(gameState.CurrentPlayer.PositionOnBoard, &gameState),
			"\n================================================================================================")
		deedsOwned := game_objects.ShowPropertyDeedsOfPlayer(gameState.CurrentPlayer.PlayerNumber, &gameState)
		gameState.CurrentPlayer.CheckToUnmortgage(gameState.CurrentPlayer, deedsOwned)
		gameState.DoDeals(gameState.AllProperties)
		gameState.CurrentPlayer.PutUpHouses(&gameState)
		gameState.RollDice()
		prePosition := gameState.CurrentPlayer.PositionOnBoard // place before we advance to our roll
		gameState.CurrentPlayer.AdvancePlayer(&gameState, gameState.CurrentDiceRoll, drawCards)
		thePropertyName, theDeed := game_objects.GetTheCurrentCard(gameState.CurrentPlayer.PositionOnBoard, &gameState)
		if theDeed != nil {
			preName, _ := game_objects.GetTheCurrentCard(prePosition, &gameState)
			movedToStr := "===> Moved from space " + strconv.Itoa(prePosition) + " " + preName + " and Landed on space " + strconv.Itoa(gameState.CurrentPlayer.PositionOnBoard) + " " + string(thePropertyName) + " owned by "
			if theDeed.Owner == 'u' {
				fmt.Println(movedToStr + "Bank <===")
				_, err := gameState.CurrentPlayer.BuyProperty(theDeed)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Purchase $", theDeed.PurchaseCost, "by player", gameState.CurrentPlayer.Name, "who now has $", gameState.CurrentPlayer.CashAvailable)
			} else {
				fmt.Println(movedToStr+"Player", int(theDeed.Owner), "(", allPlayers[theDeed.Owner].Name, ") <===")
				rent, err := theDeed.PayRent(&allPlayers[gameState.CurrentPlayer.PlayerNumber], &allPlayers[int(theDeed.Owner)], board, gameState.AllProperties)
				if err != game_objects.ErrR2O { // suppress rent to ourself messages
					fmt.Println(allPlayers[gameState.CurrentPlayer.PlayerNumber].Name, "(player", gameState.CurrentPlayer.PlayerNumber, ")", "paid $", rent, "rent to Player", allPlayers[int(theDeed.Owner)].Name, "(player", int(theDeed.Owner), ")", "(", err, ")")
					fmt.Println(allPlayers[gameState.CurrentPlayer.PlayerNumber].Name, "now has $", allPlayers[gameState.CurrentPlayer.PlayerNumber].CashAvailable, "and", allPlayers[int(theDeed.Owner)].Name, "has $", allPlayers[int(theDeed.Owner)].CashAvailable)
				}
			}
			game_objects.LogPropertiesByPlayer(&gameState)
		} else {
			sqType := board.MonopolySpace[gameState.CurrentPlayer.PositionOnBoard].SquareType
			fmt.Println("Landed on a non property square!", gameState.CurrentPlayer.PositionOnBoard, game_objects.GetPropertyType(sqType))
			gameState.ProcessNonPropertySquare(gameState.CurrentPlayer, sqType, tax, drawCards)
		}
		gameState.UnownedProperties(gameState.AllProperties) // needs to set AllPropsSold when applicable
		gameWon := gameState.NextPlayer()
		if gameWon == true || gameState.Globals.GlobalTurnsMade == numberOfTurns {
			break
		}
	}
}
