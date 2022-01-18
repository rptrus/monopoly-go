package main

import (
	"fmt"
	"github.com/rptrus/monopoly-go/game_objects"
	"github.com/rptrus/monopoly-go/setup"
	"os"
	"strconv"
	"strings"
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
	allPlayers := setup.InitializePlayers(numberOfPlayers)
	firstUp, score := game_objects.RollToSeeWhoGoesFirst(allPlayers)
	println(firstUp.Name, " is going first with score", score, "...")
	gameState := game_objects.GameState{
		CurrentPlayer:   firstUp,
		GlobalTurnsMade: 1,
		AllPlayers:      allPlayers,
		AllProperties:   propertyCardCollection,
		Others:          otherCardCollection,
	}
	game_objects.BankGameState = &gameState
	for {
		fmt.Println("\n==============================================================================="+
			"\nTurn:", gameState.GlobalTurnsMade, "Current Player", gameState.CurrentPlayer.Name, gameState.CurrentPlayer.PlayerNumber, "currently on", game_objects.GetTheCurrentCardName(gameState.CurrentPlayer.PositionOnBoard, &gameState), "rolled a", gameState.CurrentDiceRoll,
			"\n===============================================================================")
		deedsOwned := game_objects.ShowPropertyDeedsOfPlayer(gameState.CurrentPlayer.PlayerNumber, &gameState)
		gameState.CurrentPlayer.CheckToUnmortgage(gameState.CurrentPlayer, deedsOwned)
		gameState.DoDeals(gameState.AllProperties)
		gameState.CurrentPlayer.PutUpHouses(&gameState)
		gameState.RollDice()
		prePosition := gameState.CurrentPlayer.PositionOnBoard // place before we advance to our roll
		passGoPayment := gameState.CurrentPlayer.AdvancePlayer(gameState.CurrentDiceRoll, drawCards)
		if passGoPayment > 0 {
			fmt.Println("BANK PAYS PLAYER $", passGoPayment)
		}
		thePropertyName, theDeed := game_objects.GetTheCurrentCard(gameState.CurrentPlayer.PositionOnBoard, &gameState)
		if theDeed != nil {
			preName, _ := game_objects.GetTheCurrentCard(prePosition, &gameState)
			movedToStr := "===> Moved from space " + strconv.Itoa(prePosition) + " " + preName + " and Landed on space " + strconv.Itoa(gameState.CurrentPlayer.PositionOnBoard) + " " + string(thePropertyName) + " owned by "
			if theDeed.Owner == 'u' {
				fmt.Println(movedToStr+"Bank", theDeed.Owner, "<===")
				_, err := gameState.CurrentPlayer.BuyProperty(theDeed)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println("Purchase $", theDeed.PurchaseCost, "by player", gameState.CurrentPlayer.Name, "who now has $", gameState.CurrentPlayer.CashAvailable)
			} else {
				fmt.Println(movedToStr+"Player", int(theDeed.Owner), "<===")
				rent, err := theDeed.PayRent(&allPlayers[gameState.CurrentPlayer.PlayerNumber], &allPlayers[int(theDeed.Owner)], board, gameState.AllProperties)
				if err != game_objects.ErrR2O { // suppress rent to ourself messages
					fmt.Println(allPlayers[gameState.CurrentPlayer.PlayerNumber].Name, gameState.CurrentPlayer.PlayerNumber, "paid $", rent, "rent to Player", allPlayers[int(theDeed.Owner)].Name, int(theDeed.Owner), "(", err, ")")
					fmt.Println(allPlayers[gameState.CurrentPlayer.PlayerNumber].Name, "now has $", allPlayers[gameState.CurrentPlayer.PlayerNumber].CashAvailable, "and", allPlayers[int(theDeed.Owner)].Name, "has $", allPlayers[int(theDeed.Owner)].CashAvailable)
				}
			}
			names, _ := game_objects.ShowPropertiesOfPlayer(gameState.CurrentPlayer.PlayerNumber, &gameState)
			if names != nil && len(names) == 0 {
				names = append(names, "none")
			}
			game_objects.LogPropertiesByPlayer(&gameState)
			fmt.Println("\n"+allPlayers[gameState.CurrentPlayer.PlayerNumber].Name, "owns the following properties:", "[ \""+strings.Join(names, "\",\"")+"\" ]\n")
		} else {
			sqType := board.MonopolySpace[gameState.CurrentPlayer.PositionOnBoard].SquareType
			fmt.Println("Landed on a non property square!", gameState.CurrentPlayer.PositionOnBoard, game_objects.GetPropertyType(sqType))
			gameState.ProcessNonPropertySquare(gameState.CurrentPlayer, sqType, tax, drawCards)
		}
		gameState.UnownedProperties(gameState.AllProperties) // needs to set AllPropsSold when applicable
		gameWon := gameState.NextPlayer()
		if gameWon == true || gameState.GlobalTurnsMade == numberOfTurns {
			break
		}
	}
	fmt.Println("Finish.", gameState.GlobalTurnsMade, "turns.")
}
