package setup

import (
	"fmt"
	"github.com/rptrus/monopoly-go/game_objects"
)

func InitializeBoard() *game_objects.Board {
	fmt.Println("Calling initialize")
	brd := game_objects.Board{}
	// side 1
	brd.MonopolySpace[0].SquareType = game_objects.Payment
	brd.MonopolySpace[1].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[2].SquareType = game_objects.CommunityChest
	brd.MonopolySpace[3].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[4].SquareType = game_objects.Tax
	brd.MonopolySpace[5].SquareType = game_objects.Station
	brd.MonopolySpace[6].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[7].SquareType = game_objects.Chance
	brd.MonopolySpace[8].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[9].SquareType = game_objects.BuildableProperty
	// side 2
	brd.MonopolySpace[10].SquareType = game_objects.NoAction
	brd.MonopolySpace[11].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[12].SquareType = game_objects.Utility
	brd.MonopolySpace[13].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[14].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[15].SquareType = game_objects.Station
	brd.MonopolySpace[16].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[17].SquareType = game_objects.CommunityChest
	brd.MonopolySpace[18].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[19].SquareType = game_objects.BuildableProperty
	// side 3
	brd.MonopolySpace[20].SquareType = game_objects.NoAction
	brd.MonopolySpace[21].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[22].SquareType = game_objects.Chance
	brd.MonopolySpace[23].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[24].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[25].SquareType = game_objects.Station
	brd.MonopolySpace[26].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[27].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[28].SquareType = game_objects.Utility
	brd.MonopolySpace[29].SquareType = game_objects.BuildableProperty
	// side 4
	brd.MonopolySpace[30].SquareType = game_objects.Jail
	brd.MonopolySpace[31].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[32].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[33].SquareType = game_objects.CommunityChest
	brd.MonopolySpace[34].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[35].SquareType = game_objects.Station
	brd.MonopolySpace[36].SquareType = game_objects.Chance
	brd.MonopolySpace[37].SquareType = game_objects.BuildableProperty
	brd.MonopolySpace[38].SquareType = game_objects.Tax
	brd.MonopolySpace[39].SquareType = game_objects.BuildableProperty
	return &brd
}

func InitializeBank() *game_objects.Bank {
	bk := new(game_objects.Bank)
	bk.CashReservesInDollars = 20580
	bk.TotalHouses = 32
	bk.TotalHotels = 12
	return bk
}

func InitializePlayers(numberOfPlayers int) []game_objects.Player {

	var AllPlayers []game_objects.Player

	for i := 0; i < numberOfPlayers; i++ {
		p := game_objects.Player{
			PlayerNumber:    i,
			CashAvailable:   1500,
			PositionOnBoard: 0,
			Active:          true,
		}
		// using new is probably not idiomatic Go, but is still available to use. Must deref though.
		q := new(game_objects.Player)
		q.PlayerNumber = 1
		// something to note: q gives pointer, p gives the variable
		AllPlayers = append(AllPlayers, p)
	}
	// give some names to the players, make it less boring
	AllPlayers[0].Name = "Fred"
	AllPlayers[0].Token = "Wheelbarrow"
	AllPlayers[1].Name = "Mary"
	AllPlayers[1].Token = "Racing car"
	AllPlayers[2].Name = "Jason"
	AllPlayers[2].Token = "Top Hat"
	AllPlayers[3].Name = "Sally"
	AllPlayers[3].Token = "Cat"
	AllPlayers[4].Name = "Bradley"
	AllPlayers[4].Token = "Boot"
	AllPlayers[5].Name = "Indigo"
	AllPlayers[5].Token = "Thimble"

	for a, b := range AllPlayers {
		fmt.Println("Player", a, ":", b.Name, b.Token, "$", b.CashAvailable)
	}
	return AllPlayers
}
