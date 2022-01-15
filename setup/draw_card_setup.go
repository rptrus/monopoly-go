package setup

import "github.com/rptrus/monopoly-go/game_objects"

func InitializeDrawCards() *game_objects.CardCollection {

	// CHANCE

	var card game_objects.DrawCard
	CardCollection := new(game_objects.CardCollection)

	spaceToMove := []int{24, 39, 11, 0}
	const CC = 16 // offset for community chest

	station := game_objects.Station
	utility := game_objects.Utility

	ct1 := game_objects.Transaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   200,
	}
	card = game_objects.DrawCard{
		Id:           0,
		Designator:   'H',
		Content:      "Advance To Go. Collect $200",
		BankToPlayer: &ct1,
		MoveToSpace:  &spaceToMove[3],
	}
	CardCollection.AllDrawCards[0] = card
	//a := game_objects.CardCollection{AllDrawCards: [20]game_objects.DrawCard{}}

	card = game_objects.DrawCard{
		Id:          1,
		Designator:  'H',
		Content:     "Advance To Trafalgar Square. If you pass Go, collect $200",
		MoveToSpace: &spaceToMove[0],
	}
	CardCollection.AllDrawCards[1] = card

	card = game_objects.DrawCard{
		Id:          2,
		Designator:  'H',
		Content:     "Advance To Mayfair",
		MoveToSpace: &spaceToMove[1],
	}
	CardCollection.AllDrawCards[2] = card

	card = game_objects.DrawCard{
		Id:          3,
		Designator:  'H',
		Content:     "Advance To Pall Mall. If you pass Go, collect $200",
		MoveToSpace: &spaceToMove[2],
	}
	CardCollection.AllDrawCards[3] = card

	card = game_objects.DrawCard{
		Id:          4,
		Designator:  'H',
		Content:     "Advance To the nearest station. If owned, pay wonder twice the rental to which they are otherwise entitled",
		NearestType: &station,
	}
	CardCollection.AllDrawCards[4] = card

	ct2 := game_objects.Transaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   150,
	}
	card = game_objects.DrawCard{
		Id:           5,
		Designator:   'H',
		Content:      "Your building and loan matures. Collect $150.",
		BankToPlayer: &ct2,
	}
	CardCollection.AllDrawCards[5] = card

	card = game_objects.DrawCard{
		Id:         6,
		Designator: 'H',
		Content:    "Make general repairs on all your property. For each house pay $25 and each hotel pay $100.",
	}
	CardCollection.AllDrawCards[6] = card

	ct3 := game_objects.Transaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   15,
	}
	card = game_objects.DrawCard{
		Id:           7,
		Designator:   'H',
		Content:      "Speeding fine $15.",
		PlayerToBank: &ct3,
	}
	CardCollection.AllDrawCards[7] = card

	back3Spaces := -3
	card = game_objects.DrawCard{
		Id:           8,
		Designator:   'H',
		Content:      "Go back 3 spaces.",
		RelativeMove: &back3Spaces,
	}
	CardCollection.AllDrawCards[8] = card

	ct4 := game_objects.Transaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   50,
	}
	card = game_objects.DrawCard{
		Id:            9,
		Designator:    'H',
		Content:       "You have been elected chairman of the board. Pay each player $50.",
		PlayerPaysAll: &ct4,
	}
	CardCollection.AllDrawCards[9] = card

	kingsCrossStation := 5
	card = game_objects.DrawCard{
		Id:          10,
		Designator:  'H',
		Content:     "Take a trip to Kings Cross Station. If you pass Go, collect $200",
		MoveToSpace: &kingsCrossStation,
	}
	CardCollection.AllDrawCards[10] = card

	// appears twice in a deck
	card = game_objects.DrawCard{
		Id:          11,
		Designator:  'H',
		Content:     "Advance To the nearest station. If owned, pay wonder twice the rental to which they are otherwise entitled",
		NearestType: &station,
	}
	CardCollection.AllDrawCards[11] = card

	DirectJail := 10
	card = game_objects.DrawCard{
		Id:          12,
		Designator:  'H',
		Content:     "Go to Jail. Go directly to Jail. Do no pass Go, Do not collect $200",
		MoveToSpace: &DirectJail,
	}
	CardCollection.AllDrawCards[12] = card

	ct5 := game_objects.Transaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   50,
	}
	card = game_objects.DrawCard{
		Id:           13,
		Designator:   'H',
		Content:      "Bank pays you dividend of $50",
		BankToPlayer: &ct5,
	}
	CardCollection.AllDrawCards[13] = card

	card = game_objects.DrawCard{
		Id:          14,
		Designator:  'H',
		Content:     "Advance to the nearest Utility. If unowned, you may buy it from the bank. If owned, throw dice and pay pwner 10 times the amount thrown.",
		NearestType: &utility,
	}
	CardCollection.AllDrawCards[14] = card

	card = game_objects.DrawCard{
		Id:         15,
		Designator: 'H',
		Content:    "Get out of Jail free. This card may be kept until needed or traded.",
	}
	CardCollection.AllDrawCards[15] = card

	// COMMUNITY CHEST

	card = game_objects.DrawCard{
		Id:          0,
		Designator:  'H',
		Content:     "Go to Jail. Go Directly to jail. Do not pass Go, do not collect $200",
		MoveToSpace: &DirectJail,
	}
	CardCollection.AllDrawCards[CC+0] = card

	ct6 := game_objects.Transaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   100,
	}
	card = game_objects.DrawCard{
		Id:           1,
		Designator:   'O',
		Content:      "Hospital Fees. Pay $100",
		PlayerToBank: &ct6,
	}
	CardCollection.AllDrawCards[CC+1] = card

	ct7 := game_objects.Transaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   100,
	}
	card = game_objects.DrawCard{
		Id:           2,
		Designator:   'O',
		Content:      "Holiday fund matures. Receive $100.",
		BankToPlayer: &ct7,
	}
	CardCollection.AllDrawCards[CC+2] = card

	ct8 := game_objects.Transaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   10,
	}
	card = game_objects.DrawCard{
		Id:            3,
		Designator:    'O',
		Content:       "It's your birthday. Collect $10 from every player.",
		AllPaysPlayer: &ct8,
	}
	CardCollection.AllDrawCards[CC+3] = card

	ct9 := game_objects.Transaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   50,
	}
	card = game_objects.DrawCard{
		Id:           4,
		Designator:   'O',
		Content:      "From sale of stock you get $50.",
		BankToPlayer: &ct9,
	}
	CardCollection.AllDrawCards[CC+4] = card

	ct10 := game_objects.Transaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   100,
	}
	card = game_objects.DrawCard{
		Id:           5,
		Designator:   'O',
		Content:      "Life insurance matures. Collect $100.",
		BankToPlayer: &ct10,
	}
	CardCollection.AllDrawCards[CC+5] = card

	ct11 := game_objects.Transaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   50,
	}
	card = game_objects.DrawCard{
		Id:           6,
		Designator:   'O',
		Content:      "Doctor's fees. Pay $50.",
		PlayerToBank: &ct11,
	}
	CardCollection.AllDrawCards[CC+6] = card

	ct12 := game_objects.Transaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   10,
	}
	card = game_objects.DrawCard{
		Id:           7,
		Designator:   'O',
		Content:      "You have won second prize in a beauty contest. Collect $10.",
		BankToPlayer: &ct12,
	}
	CardCollection.AllDrawCards[CC+7] = card

	card = game_objects.DrawCard{
		Id:         8,
		Designator: 'O',
		Content:    "You are assessed for street repairs. Pay $40 per house and $115 per hotel you own.",
	}
	CardCollection.AllDrawCards[CC+8] = card

	ct13 := game_objects.Transaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   25,
	}
	card = game_objects.DrawCard{
		Id:           9,
		Designator:   'O',
		Content:      "Receieve $25 consultancy fee.",
		BankToPlayer: &ct13,
	}
	CardCollection.AllDrawCards[CC+9] = card

	ct14 := game_objects.Transaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   20,
	}
	card = game_objects.DrawCard{
		Id:           10,
		Designator:   'O',
		Content:      "Income tax refund. Collect $20.",
		BankToPlayer: &ct14,
	}
	CardCollection.AllDrawCards[CC+10] = card

	card = game_objects.DrawCard{
		Id:         11,
		Designator: 'O',
		Content:    "Get out of Jail free. This card may be kept until needed or traded.",
	}
	CardCollection.AllDrawCards[CC+11] = card

	ct15 := game_objects.Transaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   200,
	}
	card = game_objects.DrawCard{
		Id:           12,
		Designator:   'O',
		Content:      "Bank error in your favour. Collect $200.",
		BankToPlayer: &ct15,
	}
	CardCollection.AllDrawCards[CC+12] = card

	ct16 := game_objects.Transaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   100,
	}
	card = game_objects.DrawCard{
		Id:           13,
		Designator:   'O',
		Content:      "You inherit $100.",
		BankToPlayer: &ct16,
	}
	CardCollection.AllDrawCards[CC+13] = card

	ct17 := game_objects.Transaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   50,
	}
	card = game_objects.DrawCard{
		Id:           14,
		Designator:   'O',
		Content:      "School fees. Pay $50.",
		PlayerToBank: &ct17,
	}
	CardCollection.AllDrawCards[CC+14] = card

	card = game_objects.DrawCard{
		Id:          15,
		Designator:  'O',
		Content:     "Advance to Go. (Collect $200)",
		MoveToSpace: &spaceToMove[3],
	}
	CardCollection.AllDrawCards[CC+15] = card

	CardCollection.ShuffleOrderH = game_objects.GenerateOrderForChanceCommunityChestCards()
	CardCollection.ShuffleOrderO = game_objects.GenerateOrderForChanceCommunityChestCards()

	return CardCollection
}
