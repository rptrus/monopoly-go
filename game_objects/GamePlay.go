package game_objects

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	maxdicelow         int = 1
	maxdicehigh        int = 6
	placesonboard      int = 40
	minThresholdHouses int = 500
	formattingHashes   int = 28
)

var (
	TheBank             *Bank
	TotalPlayersPlaying int = 6
)

type GameState struct {
	CurrentPlayer         *Player
	CurrentPropertyOfTurn *Property
	CurrentDiceRoll       int
	GlobalTurnsMade       int
	allPropsSold          bool
	AllPlayers            []Player
	AllProperties         *PropertyCollection
	Others                *OtherPropertyCollection
}

func RollToSeeWhoGoesFirst(AllPlayers []Player) (*Player, int) {
	var (
		highestSoFarPlayer int = 0
		highestSoFarScore  int = 0
		rollOffPlayer      []int
	)

	for i, _ := range AllPlayers {
		total := rollDice()
		if total > highestSoFarScore {
			rollOffPlayer = rollOffPlayer[:0]
			highestSoFarScore = total
			highestSoFarPlayer = i
		}
		if total == highestSoFarScore {
			rollOffPlayer = append(rollOffPlayer, i)
		}
	}

	if len(rollOffPlayer) > 1 {
		// TIE! - roll off time. If again a tie, then the leftmost player will prevail
		highestSoFarPlayer = 0
		highestSoFarScore = 0
		for i, _ := range rollOffPlayer {
			total := rollDice()
			if total > highestSoFarScore {
				highestSoFarScore = total
				highestSoFarPlayer = rollOffPlayer[i]
			}
		}
	}
	fmt.Println(AllPlayers[highestSoFarPlayer].Name, AllPlayers[highestSoFarPlayer].Token, highestSoFarPlayer, " wins the toss")
	return &AllPlayers[highestSoFarPlayer], highestSoFarScore
}

func rollDice() int {
	time.Sleep(1 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	first := maxdicelow + rand.Intn(maxdicehigh-maxdicelow+1)
	second := maxdicelow + rand.Intn(maxdicehigh-maxdicelow+1)
	total := first + second
	return total
}

// substantially similar code reuse (I know) but I'd rather keep both for convenience, mmkay?
func rollToGetOutOfJail() (int, int) {
	time.Sleep(1 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	first := maxdicelow + rand.Intn(maxdicehigh-maxdicelow+1)
	second := maxdicelow + rand.Intn(maxdicehigh-maxdicelow+1)
	return first, second
}

func (gs *GameState) RollDice() {
	gs.GlobalTurnsMade++
	gs.CurrentDiceRoll = rollDice()
	fmt.Println("4. Roll dice -> outcome: ", gs.CurrentDiceRoll)
}

// true: if game has been won by a player
func (gs *GameState) NextPlayer() bool {
	//gs.CurrentPlayer = &gs.AllPlayers[(gs.CurrentPlayer.PlayerNumber+1)%totalPlayersPlaying]
	var countActive int
	for _, j := range gs.AllPlayers {
		if j.Active == true {
			countActive++
		}
	}

	switch countActive {
	case 1:
		gs.CurrentPlayer = &gs.AllPlayers[(gs.CurrentPlayer.PlayerNumber+1)%TotalPlayersPlaying]
		if gs.CurrentPlayer.Active == true {
			break
		}
		fallthrough
	case 2:
		gs.CurrentPlayer = &gs.AllPlayers[(gs.CurrentPlayer.PlayerNumber+1)%TotalPlayersPlaying]
		if gs.CurrentPlayer.Active == true {
			break
		}
		fallthrough
	case 3:
		gs.CurrentPlayer = &gs.AllPlayers[(gs.CurrentPlayer.PlayerNumber+1)%TotalPlayersPlaying]
		if gs.CurrentPlayer.Active == true {
			break
		}
		fallthrough
	case 4:
		gs.CurrentPlayer = &gs.AllPlayers[(gs.CurrentPlayer.PlayerNumber+1)%TotalPlayersPlaying]
		if gs.CurrentPlayer.Active == true {
			break
		}
		fallthrough
	case 5:
		gs.CurrentPlayer = &gs.AllPlayers[(gs.CurrentPlayer.PlayerNumber+1)%TotalPlayersPlaying]
		if gs.CurrentPlayer.Active == true {
			break
		}
		fallthrough
	default:
		gs.CurrentPlayer = &gs.AllPlayers[(gs.CurrentPlayer.PlayerNumber+1)%TotalPlayersPlaying]
		if gs.CurrentPlayer.Active == true {
			break
		}
	}
	fmt.Println(gs.CurrentPlayer.Name, "is now up")
	fmt.Println("\n------------------------------------------------------------------------------------------------")

	if countActive == 1 {
		namelen := len(gs.CurrentPlayer.Name)
		for i := 0; i < namelen+formattingHashes; i++ {
			fmt.Print("#")
		}
		fmt.Println("\n🏆 Player", gs.CurrentPlayer.Name, "has won the game!")
		for i := 0; i < namelen+formattingHashes; i++ {
			fmt.Print("#")
		}
		fmt.Println()
		return true
	}
	return false
}

// board position to property[28]
func GetTheCurrentCard(board int, gs *GameState) (string, *PropertyDeed) {
	for _, card := range gs.AllProperties.AllProperty {
		// j is single entry map of name:Property
		aSingularCardMap := card.Card
		for n, v := range aSingularCardMap {
			if v.PositionOnBoard == board {
				return n, v
			}
		}
	}
	return GetTheOtherPropertyName(board, gs.Others), nil
}

func GetTheOtherPropertyName(board int, oc *OtherPropertyCollection) string {
	for _, card := range oc.AllProperty {
		aSingularCardMap := card.Card
		for n, v := range aSingularCardMap {
			if v.PositionOnBoard == board {
				return n
			}
		}
	}
	return "(unknown)"
}

// convenience if we just want the name, we can use directly in a fmt.println statement
func GetTheCurrentCardName(board int, gs *GameState) string {
	name, _ := GetTheCurrentCard(board, gs)
	return name
}

func (gs *GameState) ProcessNonPropertySquare(CurrentPlayer *Player, sqType int, tax int, cc *CardCollection) {
	taxCollection := 0
	switch sqType {
	case Tax:
		fmt.Println("Tax collection!")
		taxCollection += tax
		t := Transaction{
			//gs: gs,
			Sender:   CurrentPlayer,
			Receiver: nil,
			Amount:   tax,
		}
		// general tax need 200
		if CurrentPlayer.PositionOnBoard == 4 {
			t.Amount += tax
		}
		t.TransactWithBank()
		fmt.Println("Collected Tax: $", t.Amount)
	case Jail:
		if CurrentPlayer.PositionOnBoard == 30 {
			CurrentPlayer.PositionOnBoard = 10
			CurrentPlayer.JailTurns = 3
		}
	case JustVisiting, FreeParking:
		fmt.Println("Have a rest!")
	case Payment:
		fmt.Println("Landed on GO!")
	case Chance:
		ofs := 0
		gs.processDrawCard(ofs, cc)
	case CommunityChest:
		ofs := 16
		gs.processDrawCard(ofs, cc)
	default:
		fmt.Println("Unknown or To Be Implemented")
	}
}

func (gs *GameState) processDrawCard(offset int, cc *CardCollection) {
	var card DrawCard
	if offset == 0 {
		cc.CurrentCardH = (cc.CurrentCardH + 1) % len(cc.ShuffleOrderH)
		card = cc.AllDrawCards[cc.ShuffleOrderH[cc.CurrentCardH]]
	} else {
		cc.CurrentCardO = (cc.CurrentCardO + 1) % len(cc.ShuffleOrderO)
		card = cc.AllDrawCards[offset+cc.ShuffleOrderO[cc.CurrentCardO]]
	}
	fmt.Println(drawCardType(card.Designator), "Card", card.Id, "->", card.Content)
	processDrawCardInternal(&card, gs, cc)
}

func (gs *GameState) GoToSquare(space int, paymentCheck bool) {

	if space < 0 || space > 39 {
		panic("We are attempting to move to a board space out of range: " + strconv.Itoa(space))
	}
	prePosition := gs.CurrentPlayer.PositionOnBoard
	gs.CurrentPlayer.PositionOnBoard = space
	fmt.Println("Player has moved to space", space, GetTheCurrentCardName(space, gs))
	if gs.CurrentPlayer.PositionOnBoard < prePosition && paymentCheck == true {
		gs.CurrentPlayer.pay200Dollars()
	}
}

func (gs *GameState) RemoveToken(playerToRemove *Player) {
	fmt.Println("Removing token", playerToRemove.Token, "played by", playerToRemove.Name)
	//var playerToExterminate int
	for _, j := range gs.AllPlayers {
		if j.PlayerNumber == playerToRemove.PlayerNumber {
			//playerToExterminate = i
			j.Active = false
			break
		}
	}
}
