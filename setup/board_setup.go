package setup

// Key things learned here are range over for loop
// new provides a pointer and struct literal returns the actual data structure (use & if necessary to get address)
// slices don't provide a length, but are otherwise defined like arrays

import (
	"fmt"
	"github.com/rptrus/monopoly-go/game_objects"
)

func InitializeBoard() {
	fmt.Println("Calling initialize")
	brd := game_objects.Board{}
	// side 1
	brd.MonopolySpace[0].SquareType = 4
	brd.MonopolySpace[1].SquareType = 0
	brd.MonopolySpace[2].SquareType = 1
	brd.MonopolySpace[3].SquareType = 0
	brd.MonopolySpace[4].SquareType = 2
	brd.MonopolySpace[5].SquareType = 3
	brd.MonopolySpace[6].SquareType = 0
	brd.MonopolySpace[7].SquareType = 1
	brd.MonopolySpace[8].SquareType = 0
	brd.MonopolySpace[9].SquareType = 0
	// side 2
	brd.MonopolySpace[10].SquareType = 6
	brd.MonopolySpace[11].SquareType = 0
	brd.MonopolySpace[12].SquareType = 5
	brd.MonopolySpace[13].SquareType = 0
	brd.MonopolySpace[14].SquareType = 0
	brd.MonopolySpace[15].SquareType = 3
	brd.MonopolySpace[16].SquareType = 0
	brd.MonopolySpace[17].SquareType = 1
	brd.MonopolySpace[18].SquareType = 0
	brd.MonopolySpace[19].SquareType = 0
	// side 3
	brd.MonopolySpace[20].SquareType = 4
	brd.MonopolySpace[21].SquareType = 0
	brd.MonopolySpace[22].SquareType = 1
	brd.MonopolySpace[23].SquareType = 0
	brd.MonopolySpace[24].SquareType = 0
	brd.MonopolySpace[25].SquareType = 3
	brd.MonopolySpace[26].SquareType = 0
	brd.MonopolySpace[27].SquareType = 0
	brd.MonopolySpace[28].SquareType = 5
	brd.MonopolySpace[29].SquareType = 0
	// side 4
	brd.MonopolySpace[30].SquareType = 6
	brd.MonopolySpace[31].SquareType = 0
	brd.MonopolySpace[32].SquareType = 0
	brd.MonopolySpace[33].SquareType = 1
	brd.MonopolySpace[34].SquareType = 0
	brd.MonopolySpace[35].SquareType = 3
	brd.MonopolySpace[36].SquareType = 1
	brd.MonopolySpace[37].SquareType = 0
	brd.MonopolySpace[38].SquareType = 2
	brd.MonopolySpace[39].SquareType = 0
}

func InitializeBank() {
	bk := new(game_objects.Bank)
	bk.CashReservesInDollars = 20580
}

//// functions to work on all players ////

func InitializePlayers(numberOfPlayers int) {
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
		game_objects.AllPlayers = append(game_objects.AllPlayers, p)
		fmt.Println("Initialized players:\n", p)
	}

	for a, b := range game_objects.AllPlayers {
		fmt.Println("Player: \n", a, b)
	}
}
