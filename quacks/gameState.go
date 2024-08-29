package quacks

import (
	"context"
	"fmt"
	"sort"

	"github.com/looplab/fsm"
)

type Input struct {
	Description string
	Options     []string
	Choice      int
	Choice2     []Chip // todo refactor for buying chips
	Player      int    // Player Position
	Code        int
}

type GameState struct {
	Players     []Player
	Round       int
	fortune     int
	winner      []int
	book        int
	bombLimit   int
	Awaiting    *Input
	debug       bool
	FSM         *fsm.FSM
	fortuneDeck []Fortune
	Stats       *Stats
	Id          string
	Status      string // Status of the game for client consumption
}

func (gs *GameState) GameIsOver() bool {
	return len(gs.winner) > 0
}

func (gs *GameState) enterState(e *fsm.Event) {
	// Add your state transition logic here
	if gs.debug {
		fmt.Println("State Transition -> -> -> ")
		fmt.Printf("Entering %s from %s\n\n", e.Dst, e.Src)
	}

	// reset buying flag
	// Reset Cherry Bomb count
	if e.Dst == FortuneState.String() {
		if gs.debug {
			fmt.Println("Resetting buy flag for all players")
		}
		for _, player := range gs.Players {
			player.isDoneDrawing = false
			player.Board.CherryBombValue = 0
			player.hasCompletedTheFortune = false
			player.hasSpentRubies = false
		}

		gs.fortune = -1
	}
}

func (gs GameState) PrintRound() {
	fmt.Printf("===============\nRound %d\n---------------\n", gs.Round)
	scores := ""
	for _, player := range gs.Players {
		scores = scores + fmt.Sprintf("%s - %d\n", player.Name, player.score)
	}
	fmt.Printf("Scores:\n%s\n===============\n", scores)

}

// ByScore implements the sort.Interface for sorting players by score.
type ByScore []Player

func (p ByScore) Len() int           { return len(p) }
func (p ByScore) Less(i, j int) bool { return p[i].score > p[j].score }
func (p ByScore) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// GetTopPlayers returns a slice of players with the highest score.
func GetTopPlayers(players []Player) []int {
	sort.Sort(ByScore(players))
	highestScore := players[0].score
	var topPlayers []int

	for i, player := range players {
		if player.score == highestScore {
			topPlayers = append(topPlayers, i)
		} else {
			break // Stop when we encounter a lower score
		}
	}

	return topPlayers
}

func (gs *GameState) nextRound(e *fsm.Event) {
	if gs.Round == 9 {
		// Assign winners here
		gs.winner = GetTopPlayers(gs.Players)
		return
	}

	gs.Round = gs.Round + 1

	fmt.Printf("===============\nStarting Round %d\n---------------\n", gs.Round)
	scores := ""
	for _, player := range gs.Players {
		scores = scores + fmt.Sprintf("%s - %d\n", player.Name, player.score)
	}
	fmt.Printf("Scores:\n%s\n===============\n", scores)
}

func (gs *GameState) GetPlayersByScore() []string {
	sort.Slice(gs.Players, func(i, j int) bool {
		return gs.Players[i].score > gs.Players[j].score
	})

	playerNames := make([]string, len(gs.Players))
	for i, player := range gs.Players {
		playerNames[i] = player.Name
	}

	return playerNames
}

func (gs GameState) GetPlayerNames() []string {
	names := []string{}
	for _, player := range gs.Players {
		names = append(names, player.Name)
	}
	return names
}

func CreateGameState(playerNames []string, gameId string, debug bool) *GameState {
	players := setUpPlayers(playerNames)

	gs := &GameState{
		players,
		1,
		-1,
		[]int{},
		1,
		7,
		nil,
		debug,
		nil,
		createFortunes(),
		nil,
		gameId,
		"",
	}

	gs.FSM = fsm.NewFSM(
		ClosedState.String(),
		fsm.Events{
			{Name: Start.String(), Src: []string{ClosedState.String()}, Dst: FortuneState.String()},
			{Name: ReadFortune.String(), Src: []string{FortuneState.String()}, Dst: FortuneInputState.String()},
			{Name: HandleFortune.String(), Src: []string{FortuneInputState.String()}, Dst: FortuneState.String()},
			{Name: AssignRatTails.String(), Src: []string{FortuneState.String()}, Dst: RatTailsState.String()},
			{Name: BeginPreparation.String(), Src: []string{RatTailsState.String()}, Dst: PreparationState.String()},
			{Name: PreparationInput.String(), Src: []string{PreparationState.String()}, Dst: PreparationInputState.String()},
			{Name: HandlePreparationInput.String(), Src: []string{PreparationInputState.String()}, Dst: PreparationState.String()},
			{Name: EnterScoring.String(), Src: []string{PreparationState.String()}, Dst: ScoringState.String()},
			{Name: ScoringInput.String(), Src: []string{ScoringState.String()}, Dst: ScoringInputState.String()},
			{Name: EnterBuying.String(), Src: []string{ScoringState.String(), BuyingInputState.String()}, Dst: BuyingState.String()},
			{Name: HandleBuying.String(), Src: []string{BuyingState.String()}, Dst: BuyingInputState.String()},
			{Name: LeaveBuying.String(), Src: []string{BuyingState.String(), BuyingInputState.String()}, Dst: ScoringState.String()},
			{Name: HandleScoringInput.String(), Src: []string{ScoringInputState.String()}, Dst: ScoringState.String()},
			{Name: EnterRubySpending.String(), Src: []string{BuyingState.String(), RubySpendingInputState.String()}, Dst: RubySpendingState.String()},
			{Name: HandleRubySpending.String(), Src: []string{RubySpendingState.String()}, Dst: RubySpendingInputState.String()},
			{Name: EnterNextRound.String(), Src: []string{RubySpendingState.String(), RubySpendingInputState.String()}, Dst: FortuneState.String()},
			{Name: End.String(), Src: []string{ScoringState.String(), RubySpendingState.String()}, Dst: End.String()},
		},
		fsm.Callbacks{
			"enter_state":   func(_ context.Context, e *fsm.Event) { gs.enterState(e) },
			"leave_scoring": func(_ context.Context, e *fsm.Event) { gs.nextRound(e) },
		},
	)

	return gs
}

type State int

const (
	ClosedState State = iota
	FortuneState
	FortuneInputState
	RatTailsState
	PreparationState
	PreparationInputState
	ScoringState
	ScoringInputState
	BuyingState
	BuyingInputState
	RubySpendingState
	RubySpendingInputState
	EndState
)

func (s State) String() string {
	switch s {
	case ClosedState:
		return "closed"
	case FortuneState:
		return "fortune"
	case FortuneInputState:
		return "fortune_input"
	case RatTailsState:
		return "rat_tails"
	case PreparationState:
		return "preparation"
	case PreparationInputState:
		return "preparation_input"
	case ScoringState:
		return "scoring"
	case ScoringInputState:
		return "scoring_input"
	case BuyingState:
		return "buying_state"
	case BuyingInputState:
		return "buying_input_state"
	case RubySpendingState:
		return "ruby_spending_state"
	case RubySpendingInputState:
		return "ruby_spending_input_state"
	case EndState:
		return "end"
	default:
		return "unknown"
	}
}

type Transition int

const (
	Start Transition = iota
	ReadFortune
	HandleFortune
	AssignRatTails
	BeginPreparation
	PreparationInput
	HandlePreparationInput
	EnterScoring
	ScoringInput
	HandleScoringInput
	EnterBuying
	HandleBuying
	LeaveBuying
	EnterRubySpending
	HandleRubySpending
	EnterNextRound
	End
)

func (s Transition) String() string {
	switch s {
	case Start:
		return "start"
	case ReadFortune:
		return "read_fortune"
	case HandleFortune:
		return "handle_fortune"
	case AssignRatTails:
		return "assign_rat_tails"
	case BeginPreparation:
		return "begin_preparation"
	case PreparationInput:
		return "preparation_input"
	case HandlePreparationInput:
		return "handle_preparation_input"
	case EnterScoring:
		return "enter_scoring"
	case ScoringInput:
		return "scoring_input"
	case HandleScoringInput:
		return "handle_scoring_input"
	case EnterBuying:
		return "enter_buying"
	case HandleBuying:
		return "handle_buying"
	case LeaveBuying:
		return "leave_buying"
	case EnterRubySpending:
		return "enter_ruby_spending"
	case HandleRubySpending:
		return "handle_ruby_spending"
	case EnterNextRound:
		return "enter_next_round"
	case End:
		return "end"
	default:
		return "unknown"
	}
}
