package quacks

import (
	"context"
	"fmt"

	"github.com/looplab/fsm"
)

type Input struct {
	Description string
	Choice      int
	Player      int
}

type GameState struct {
	players  []Player
	turn     int
	fortune  int
	winner   int
	book     int
	awaiting *Input
	debug    bool
	FSM      *fsm.FSM
}

func (gs *GameState) enterState(e *fsm.Event) {
	// Add your state transition logic here
	if gs.debug {
		fmt.Println("State Transition -> -> -> ")
		fmt.Printf("Entering %s from %s\n\n", e.Dst, e.Src)
	}
}

func CreateGameState(playerNames []string, debug bool) *GameState {
	players := setUpPlayers(playerNames)

	gs := &GameState{
		players,
		0,
		0,
		0,
		1,
		nil,
		debug,
		nil,
	}

	gs.FSM = fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: Start.String(), Src: []string{"closed"}, Dst: "fortune"},
			{Name: ReadFortune.String(), Src: []string{"fortune"}, Dst: "fortune_input"},
			{Name: HandleFortune.String(), Src: []string{"fortune_input"}, Dst: "fortune"},
			{Name: AssignRatTails.String(), Src: []string{"fortune"}, Dst: "rat_tails"},
			{Name: BeginPreparation.String(), Src: []string{"rat_tails"}, Dst: "preparation"},
			{Name: PreparationInput.String(), Src: []string{"preparation"}, Dst: "preparation_input"},
			{Name: HandlePreparationInput.String(), Src: []string{"preparation_input"}, Dst: "preparation"},
			{Name: EnterScoring.String(), Src: []string{"preparation"}, Dst: "scoring"},
			{Name: ScoringInput.String(), Src: []string{"scoring"}, Dst: "scoring_input"},
			{Name: HandleScoringInput.String(), Src: []string{"scoring_input"}, Dst: "scoring"},
			{Name: EnterNextRound.String(), Src: []string{"scoring"}, Dst: "fortune"},
			{Name: End.String(), Src: []string{"scoring"}, Dst: "end"},
		},
		fsm.Callbacks{
			"enter_state": func(_ context.Context, e *fsm.Event) { gs.enterState(e) },
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

func GetHandleFortuneString() string {
	return "handle_fortune"
}

type State int

const (
	Start State = iota
	ReadFortune
	HandleFortune
	AssignRatTails
	BeginPreparation
	PreparationInput
	HandlePreparationInput
	EnterScoring
	ScoringInput
	HandleScoringInput
	EnterNextRound
	End
)

func (s State) String() string {
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
	case EnterNextRound:
		return "enter_next_round"
	case End:
		return "end"
	default:
		return "unknown"
	}
}
