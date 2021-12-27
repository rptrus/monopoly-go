package main

import (
	"fmt"
	"github.com/rptrus/monopoly-go/setup"
)

func main() {
	fmt.Println("Starting Monopoly Go SIM")
	setup.InitializeBoard()
	setup.InitializeBank()
	setup.SetupPropertyCards()
}
