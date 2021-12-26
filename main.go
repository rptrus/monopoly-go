package main

import (
	"alphastar/setup"
	"fmt"
)

type player struct {
	playerNumber  int
	cashAvailable int
}

func main() {
	fmt.Println("Starting Monopoly Go SIM")
	setup.InitializeBoard()
}
