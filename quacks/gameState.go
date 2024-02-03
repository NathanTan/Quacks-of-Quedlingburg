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
	Player      int
	Code        int
}

type GameState struct {
	players     []Player
	round       int
	fortune     int
	winner      int
	book        int
	bombLimit   int
	awaiting    *Input
	debug       bool
	FSM         *fsm.FSM
	fortuneDeck []Fortune
}

func (gs *GameState) GameIsOver() bool {
	if gs.winner > 0 {
		return true
	}
	return false
}

func (gs *GameState) enterState(e *fsm.Event) {
	// Add your state transition logic here
	if gs.debug {
		fmt.Println("State Transition -> -> -> ")
		fmt.Printf("Entering %s from %s\n\n", e.Dst, e.Src)
	}
}

func (gs *GameState) nextRound(e *fsm.Event) {
	if gs.round == 9 {
		gs.winner = 100
		// Assign winner here
		return
	}

	gs.round = gs.round + 1

	fmt.Printf("===============\nStarting Round %d\n===============\n", gs.round)
}

func (gs *GameState) GetPlayersByScore() []string {
	sort.Slice(gs.players, func(i, j int) bool {
		return gs.players[i].score > gs.players[j].score
	})

	playerNames := make([]string, len(gs.players))
	for i, player := range gs.players {
		playerNames[i] = player.name
	}

	return playerNames
}

func CreateGameState(playerNames []string, debug bool) *GameState {
	players := setUpPlayers(playerNames)

	gs := &GameState{
		players,
		0,
		1,
		0,
		1,
		7,
		nil,
		debug,
		nil,
		nil,
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
			{Name: EnterBuying.String(), Src: []string{ScoringState.String()}, Dst: BuyingState.String()},
			{Name: HandleBuying.String(), Src: []string{BuyingState.String()}, Dst: BuyingInputState.String()},
			{Name: LeaveBuying.String(), Src: []string{BuyingState.String()}, Dst: ScoringState.String()},
			{Name: HandleScoringInput.String(), Src: []string{ScoringInputState.String()}, Dst: ScoringState.String()},
			{Name: EnterNextRound.String(), Src: []string{ScoringState.String()}, Dst: FortuneState.String()},
			{Name: End.String(), Src: []string{ScoringState.String()}, Dst: End.String()},
		},
		fsm.Callbacks{
			"enter_state":   func(_ context.Context, e *fsm.Event) { gs.enterState(e) },
			"leave_scoring": func(_ context.Context, e *fsm.Event) { gs.nextRound(e) },
		},
	)

	return gs
}

func GameIsOver(gs GameState) bool {
	if gs.winner > 0 {
		return true
	}
	return false
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
	case EnterNextRound:
		return "enter_next_round"
	case End:
		return "end"
	default:
		return "unknown"
	}
}
