package setup

// Key things learned here are range over for loop
// new provides a pointer and struct literal returns the actual data structure (use & if necessary to get address)
// slices don't provide a length, but are otherwise defined like arrays

import (
	"fmt"
	"github.com/rptrus/monopoly-go/game_objects"
)

func InitializeBoard() *game_objects.Board {
	fmt.Println("Calling initialize")
	brd := game_objects.Board{}
	// side 1
	brd.MonopolySpace[0].SquareType = 4
	brd.MonopolySpace[1].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[2].SquareType = 1
	brd.MonopolySpace[3].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[4].SquareType = 2
	brd.MonopolySpace[5].SquareType = game_objects.Station
	brd.MonopolySpace[6].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[7].SquareType = 1
	brd.MonopolySpace[8].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[9].SquareType = game_objects.BuildableProperty
	// side 2
	brd.MonopolySpace[10].SquareType = 6
	brd.MonopolySpace[11].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[12].SquareType = 5
	brd.MonopolySpace[13].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[14].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[15].SquareType = game_objects.Station
	brd.MonopolySpace[16].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[17].SquareType = 1
	brd.MonopolySpace[18].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[19].SquareType = game_objects.BuildableProperty
	// side 3
	brd.MonopolySpace[20].SquareType = 4
	brd.MonopolySpace[21].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[22].SquareType = 1
	brd.MonopolySpace[23].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[24].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[25].SquareType = game_objects.Station
	brd.MonopolySpace[26].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[27].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[28].SquareType = 5
	brd.MonopolySpace[29].SquareType = game_objects.BuildableProperty
	// side 4
	brd.MonopolySpace[30].SquareType = 6
	brd.MonopolySpace[31].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[32].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[33].SquareType = 1
	brd.MonopolySpace[34].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[35].SquareType = game_objects.Station
	brd.MonopolySpace[36].SquareType = 1
	brd.MonopolySpace[37].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[38].SquareType = 2
	brd.MonopolySpace[39].SquareType = game_objects.BuildableProperty
	return &brd
}

func InitializeBank() *game_objects.Bank {
	bk := new(game_objects.Bank)
	bk.CashReservesInDollars = 20580
	return bk
}

//// functions to work on all players ////

func InitializePlayers(numberOfPlayers int) []game_objects.Player {

	var AllPlayers []game_objects.Player // slice

	for i := 0; i < numberOfPlayers; i++ {
		p := game_objects.Player{
			PlayerNumber:    i,
			CashAvailable:   1500,
			PositionOnBoard: 0,
		}
		// using new is probably not idiomatic Go, but is still available to use. Must deref though.
		q := new(game_objects.Player)
		q.PlayerNumber = 1
		// something to note: p gives pointer, q gives the variable
		AllPlayers = append(AllPlayers, p)
		fmt.Println("Initialized players:\n", p)
	}

	for a, b := range AllPlayers {
		fmt.Println("Player: \n", a, b)
	}
	return AllPlayers
}
