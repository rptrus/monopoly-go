package setup

import "github.com/rptrus/monopoly-go/game_objects"

func InitializeDrawCards() *game_objects.CardCollection {

	// CHANCE

	var ct game_objects.CardTransaction
	var card game_objects.DrawCard
	CardCollection := new(game_objects.CardCollection)

	spaceToMove := []int{24, 39, 11}
	const CC = 16 // offset for community chest

	ct = game_objects.CardTransaction{
		Sender:   nil,
		Receiver: nil,
		Amount:   200,
	}
	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'H',
		Content:        "Advance To Go. Collect $200",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   &ct,
	}
	CardCollection.AllDrawCards[0] = card
	//a := game_objects.CardCollection{AllDrawCards: [20]game_objects.DrawCard{}}

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'H',
		Content:        "Advance To Trafalgar Square. If you pass Go, collect $200",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    &spaceToMove[0],
	}
	CardCollection.AllDrawCards[1] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'H',
		Content:        "Advance To Mayfair",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    &spaceToMove[1],
	}
	CardCollection.AllDrawCards[2] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'H',
		Content:        "Advance To Pall Mall. If you pass Go, collect $200",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    &spaceToMove[2],
	}
	CardCollection.AllDrawCards[3] = card

	//Sq := new(game_objects.Square)
	//Sq.SquareType = game_objects.Station

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'H',
		Content:        "Advance To the nearest station. If owned, pay wonder twice the rental to which they are otherwise entitled",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
		NearestType:    &game_objects.Square{SquareType: game_objects.Station},
	}
	CardCollection.AllDrawCards[4] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'H',
		Content:        "Your building and loan matures. Collect $150.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
		NearestType:    nil,
	}
	CardCollection.AllDrawCards[5] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'H',
		Content:        "Make general repairs on all your property. For each house pay $25 and each hotel pay $100.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
		NearestType:    nil,
	}
	CardCollection.AllDrawCards[6] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'H',
		Content:        "Speeding fine $15.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
		NearestType:    nil,
	}
	CardCollection.AllDrawCards[7] = card

	back3Spaces := -3

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'H',
		Content:        "Go back 3 spaces.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
		NearestType:    nil,
		RelativeMove:   &back3Spaces,
	}
	CardCollection.AllDrawCards[8] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'H',
		Content:        "You have been elected chairman of the board. Pay each player $50.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
		NearestType:    nil,
	}
	CardCollection.AllDrawCards[9] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'H',
		Content:        "Take a trip to Kings Cross Station. If you pass Go, collect $200",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
		NearestType:    nil,
	}
	CardCollection.AllDrawCards[10] = card

	// appears twice in a deck
	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'H',
		Content:        "Advance To the nearest station. If owned, pay wonder twice the rental to which they are otherwise entitled",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
		NearestType:    &game_objects.Square{SquareType: game_objects.Station},
	}
	CardCollection.AllDrawCards[11] = card

	indirectJail := 30
	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'H',
		Content:        "Go to Jail. Go directly to Jail. Do no pass Go, Do not collect $200",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    &indirectJail,
		NearestType:    nil,
	}
	CardCollection.AllDrawCards[12] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'H',
		Content:        "Bank pays you dividend of $50",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
		NearestType:    nil,
	}
	CardCollection.AllDrawCards[13] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'H',
		Content:        "Advance to the nearest Utility. If unowned, you may buy it from the bank. If owned, throw dice and pay pwner 10 times the amount thrown.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
		NearestType:    &game_objects.Square{SquareType: game_objects.Utility},
	}
	CardCollection.AllDrawCards[14] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'H',
		Content:        "Get out of Jail free. This card may be kept until needed or traded.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
		NearestType:    &game_objects.Square{SquareType: game_objects.Utility},
	}
	CardCollection.AllDrawCards[15] = card

	// COMMUNITY CHEST

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'H',
		Content:        "Go to Jail. Go Directly to jail. Do not pass Go, do not collect $200",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
	}
	CardCollection.AllDrawCards[CC+0] = card
	//a := game_objects.CardCollection{AllDrawCards: [20]game_objects.DrawCard{}}

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'O',
		Content:        "Hospital Fees. Pay $100",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
	}
	CardCollection.AllDrawCards[CC+1] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'O',
		Content:        "Holiday fund matures. Receive $100.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
	}
	CardCollection.AllDrawCards[CC+2] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'O',
		Content:        "It's your birthday. Collect $10 from every player.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
	}
	CardCollection.AllDrawCards[CC+3] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'O',
		Content:        "From sale of stock you get $50.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
	}
	CardCollection.AllDrawCards[CC+4] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'O',
		Content:        "Life insurance matures. Collect $100.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
	}
	CardCollection.AllDrawCards[CC+5] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'O',
		Content:        "Doctor's fees. Pay $50.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
	}
	CardCollection.AllDrawCards[CC+6] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'O',
		Content:        "You have won second prize in a beauty contest. Collect $10.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
	}
	CardCollection.AllDrawCards[CC+7] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'O',
		Content:        "You are assessed for street repairs. Pay $40 per house and $115 per hotel you own.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
	}
	CardCollection.AllDrawCards[CC+8] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'O',
		Content:        "Receieve $25 consultancy fee.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
	}
	CardCollection.AllDrawCards[CC+9] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'O',
		Content:        "Income tax refund. Collect $20.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
	}
	CardCollection.AllDrawCards[CC+10] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'O',
		Content:        "Get out of Jail free. This card may be kept until needed or traded.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
	}
	CardCollection.AllDrawCards[CC+11] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'O',
		Content:        "Bank error in your favour. Collect $200.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
	}
	CardCollection.AllDrawCards[CC+12] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'O',
		Content:        "You inherit $100.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
	}
	CardCollection.AllDrawCards[CC+13] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'O',
		Content:        "School fees. Pay $50.",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
	}
	CardCollection.AllDrawCards[CC+14] = card

	card = game_objects.DrawCard{
		Id:             0,
		Designator:     'O',
		Content:        "Advance to Go. (Collect $200)",
		PlayerToPlayer: nil,
		PlayerToBank:   nil,
		BankToPlayer:   nil,
		MoveToSpace:    nil,
	}
	CardCollection.AllDrawCards[CC+15] = card

	CardCollection.ShuffleOrder = game_objects.GenerateOrderForChanceCommunityChestCards()
	CardCollection.CurrentCard = 0

	return CardCollection
}
