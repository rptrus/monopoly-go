package setup

import "fmt"

type board struct {
	monopoly_space [40]square
}

type square struct {
	squareType int
}

const (
	buildableProperty int = iota
	chanceChest
	tax
	station
	noAction
	utility
	action
)

func InitializeBoard() {
	fmt.Println("Calling initialize")
	b := board{}
	// side 1
	b.monopoly_space[0].squareType = 4
	b.monopoly_space[1].squareType = 0
	b.monopoly_space[2].squareType = 1
	b.monopoly_space[3].squareType = 0
	b.monopoly_space[4].squareType = 2
	b.monopoly_space[5].squareType = 3
	b.monopoly_space[6].squareType = 0
	b.monopoly_space[7].squareType = 1
	b.monopoly_space[8].squareType = 0
	b.monopoly_space[9].squareType = 0
	// side 2
	b.monopoly_space[10].squareType = 6
	b.monopoly_space[11].squareType = 0
	b.monopoly_space[12].squareType = 5
	b.monopoly_space[13].squareType = 0
	b.monopoly_space[14].squareType = 0
	b.monopoly_space[15].squareType = 3
	b.monopoly_space[16].squareType = 0
	b.monopoly_space[17].squareType = 1
	b.monopoly_space[18].squareType = 0
	b.monopoly_space[19].squareType = 0
	// side 3
	b.monopoly_space[20].squareType = 4
	b.monopoly_space[21].squareType = 0
	b.monopoly_space[22].squareType = 1
	b.monopoly_space[23].squareType = 0
	b.monopoly_space[24].squareType = 0
	b.monopoly_space[25].squareType = 3
	b.monopoly_space[26].squareType = 0
	b.monopoly_space[27].squareType = 0
	b.monopoly_space[28].squareType = 5
	b.monopoly_space[29].squareType = 0
	// side 4
	b.monopoly_space[30].squareType = 6
	b.monopoly_space[31].squareType = 0
	b.monopoly_space[32].squareType = 0
	b.monopoly_space[33].squareType = 1
	b.monopoly_space[34].squareType = 0
	b.monopoly_space[35].squareType = 3
	b.monopoly_space[36].squareType = 1
	b.monopoly_space[37].squareType = 0
	b.monopoly_space[38].squareType = 2
	b.monopoly_space[39].squareType = 0
}
