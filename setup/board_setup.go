package setup

import (
	"fmt"
	game_objects "github.com/rptrus/monopoly-go/game_objects"
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
