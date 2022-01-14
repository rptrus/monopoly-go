package game_objects

import (
	"fmt"
	"math/rand"
	"time"
)

type DrawCard struct {
	Id             int
	Designator     byte // 'c' = chance, 'o' = community chest
	Content        string
	PlayerToPlayer *Transaction
	PlayerToBank   *Transaction
	BankToPlayer   *Transaction
	MoveToSpace    *int // always check if they pass Go, except for Jail or Mayfair
	RelativeMove   *int // -3 = go back 3 spaces
	NearestType    *int
	PlayerPaysAll  *Transaction
	AllPaysPlayer  *Transaction
}

type CardCollection struct {
	AllDrawCards [32]DrawCard
	ShuffleOrder []int
	CurrentCard  int // TODO: seperate out so that each has it's own counter
}

func GenerateOrderForChanceCommunityChestCards() []int {
	cardsToDeal := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(cardsToDeal)
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cardsToDeal), func(i, j int) { cardsToDeal[i], cardsToDeal[j] = cardsToDeal[j], cardsToDeal[i] })
	fmt.Println(cardsToDeal)
	return cardsToDeal
}
