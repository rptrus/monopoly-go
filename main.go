package main

import (
	"fmt"
	"github.com/rptrus/monopoly-go/setup"
)

const numberOfPlayers int = 6

func main() {
	fmt.Println("Starting Monopoly Go SIM")

	a := []string{"a", "b", "c"}
	b := []string{"x", "y", "z"}
	a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
	//... when calling a function does the opposite: if you have several arguments in a slice, it will unpack them and pass as separate arguments to a variadic function.

	setup.InitializeBoard()
	setup.InitializeBank()
	setup.InitializePropertyCards()
	setup.InitializePlayers(numberOfPlayers)
}
