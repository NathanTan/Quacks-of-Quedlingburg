package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"quacks"
	"sync"
	"types"

	"github.com/anthdm/hollywood/actor"
	"github.com/gorilla/websocket"
)

type PlayerState struct {
	Position types.Position
	HP       int
}

type PlayerSession struct {
	sessionId int
	clientId  int
	username  string
	inLobby   bool
	conn      *websocket.Conn
}

type GameServer struct {
	ctx      *actor.Context
	sessions map[*actor.PID]struct{}
}

var (
	gameStates   = make(map[string]*quacks.GameState)
	playerStates = make(map[string]*types.PlayerState)
	mutex        = &sync.Mutex{}
	player_mutex = &sync.Mutex{}
)

func newPlayerSession(sid int, conn *websocket.Conn) actor.Producer {
	fmt.Printf("Creating new player session for %d\n", sid)
	return func() actor.Receiver {
		return &PlayerSession{
			conn:      conn,
			sessionId: sid,
			clientId:  3,
			username:  "username",
			inLobby:   true,
		}
	}
}

func saveGameState(gameID string, gameState *quacks.GameState) {
	mutex.Lock()
	gameStates[gameID] = gameState
	mutex.Unlock()

	// Save the game state to the database
	// This is a placeholder. Replace with your actual database saving code.
	// db.Save(gameID, gameState)
}

func savePlayerState(gameID string, playerState *types.PlayerState) {
	player_mutex.Lock()
	playerStates[gameID] = playerState
	player_mutex.Unlock()
}

func getGameState(gameID string) *quacks.GameState {
	mutex.Lock()
	gameState := gameStates[gameID]
	mutex.Unlock()

	return gameState
}

func getPlayerState(gameID string) *types.PlayerState {
	mutex.Lock()
	gameState := playerStates[gameID]
	mutex.Unlock()

	return gameState
}

func (s *PlayerSession) Receive(c *actor.Context) {
	fmt.Println("PlayerSession Receiving messages")
	fmt.Println(c.Message())
	switch c.Message().(type) {
	case actor.Started:
		s.readLoop()
		// s.ctx = c
		// _ = msg
	}
}

func (s *PlayerSession) readLoop() {
	var msg *types.WSMessage
	for {
		_, p, err := s.conn.ReadMessage()
		if err != nil {
			fmt.Println("read error", err)
			fmt.Println("Problem Message:")
			fmt.Println(msg)
			return
		}
		// fmt.Println("Handle Message")
		// fmt.Println(msg)
		// fmt.Println(messageType)
		// fmt.Println(p)
		var receivedMsg types.WSMessage

		err = json.Unmarshal(p, &receivedMsg)
		if err != nil {
			fmt.Println("read error", err)
			fmt.Println("Problem Message:")
			fmt.Println(msg)
			return
		}
		// fmt.Println("Unmarshaled payload")
		// fmt.Println(receivedMsg)
		go s.handleMessage(&receivedMsg)

	}
}

func (s *PlayerSession) handleMessage(msg *types.WSMessage) error {
	fmt.Printf("Handling message - type: %s\n", msg.Type)
	switch msg.Type {
	case "Login":
		var loginMsg types.Login
		if err := json.Unmarshal(msg.Data, &loginMsg); err != nil {
			panic(err) // Panic since this is a single player session, not in the server
		}
		s.clientId = loginMsg.ClientId
		s.username = loginMsg.Username
		fmt.Printf("Conducted login for id: %d, name: %s\n", s.clientId, s.username)

	case "PlayerMove":
		var pm types.PlayerMove

		if err := json.Unmarshal(msg.Data, &pm); err != nil {
			panic(err)
		}

		fmt.Println("Recieved player move: ", pm)

		state := getGameState(pm.GameId)
		if pm.Move == "StartGame" {
			fmt.Println("Starting game")
			state.StartGame()
		} else {
			fmt.Println("Move not yet implemented")
		}
		state.Status = state.FSM.Current()

		fmt.Println("New state: ", state.FSM.Current())
		saveGameState("game123", state)
		state.PrintGameStateForDebugging()

		b, err := json.Marshal(state)
		if err != nil {
			log.Fatal(err)
		}
		sendGameState(s, "GameState", b)

	case "PlayerState":
		// TODO: Recieve the player id here and send for a specific player
		var ps types.PlayerStateRequest

		if err := json.Unmarshal(msg.Data, &ps); err != nil {
			panic(err)
		}
		fmt.Println("PlayerState:")
		fmt.Println(ps)

		b, err := json.Marshal(ps)
		if err != nil {
			log.Fatal(err)
		}
		sendGameState(s, "PlayerState", b)

	case "NewGameStateRequest":

		var gs types.NewGameStateRequest

		if err := json.Unmarshal(msg.Data, &gs); err != nil {
			panic(err)
		}
		playerNames := gs.PlayerNames

		state := quacks.CreateGameState(playerNames, "game123", true)
		state.Status = state.FSM.Current()

		saveGameState("game123", state)

		fmt.Println("Created new game - state: ", state.FSM.Current())
		fmt.Println("New GameState: ", state)
		b, err := json.Marshal(state)
		if err != nil {
			log.Fatal(err)
		}
		sendGameState(s, "NewGameState", b)

	case "GameStateRequest":

		var gs types.GameStateRequest

		if err := json.Unmarshal(msg.Data, &gs); err != nil {
			panic(err)
		}

		state := getGameState(gs.GameId)
		state.Status = state.FSM.Current()

		fmt.Println("Retriving GameState: ", state)
		b, err := json.Marshal(state)
		if err != nil {
			log.Fatal(err)
		}
		sendGameState(s, "GameState", b)

	case "StartGame":
		var gs types.GameStateRequest
		if err := json.Unmarshal(msg.Data, &gs); err != nil {
			panic(err)
		}
		state := getGameState(gs.GameId)
		state.StartGame()
		state.Status = state.FSM.Current()

		fmt.Println("Starting game ", gs.GameId, " state: ", state.FSM.Current())

		state.PrintGameStateForDebugging()

		b, err := json.Marshal(state)
		if err != nil {
			log.Fatal(err)
		}
		sendGameState(s, "GameState", b)

	}

	return nil
}

func sendGameState(s *PlayerSession, messageType string, payload []byte) error {
	msg := types.WSMessage{
		Type: messageType,
		Data: payload,
	}

	jsonMsg, err := json.Marshal(msg)

	fmt.Println("Sending Message: " + string(jsonMsg))

	// Send the response message
	err = s.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		// handle error
		fmt.Println("error while sending gamestate:", err)
		return err
	}

	return nil
}

func newGameServer() actor.Receiver {
	return &GameServer{
		sessions: make(map[*actor.PID]struct{}),
	}
}

func (s *GameServer) Receive(c *actor.Context) {
	fmt.Printf("GameServer Receiving messages\n")
	switch msg := c.Message().(type) {
	case actor.Started:
		s.startHTTP()
		s.ctx = c
		_ = msg
	}
}

func (s *GameServer) startHTTP() {
	fmt.Println("starting HTTP server on port 40000")
	go func() {
		http.HandleFunc("/ws", s.handleWS)
		http.ListenAndServe(":40000", nil)
	}()
}

func newGameClient(conn *websocket.Conn, username string) *types.GameClient {
	return &types.GameClient{
		ClientId: rand.Intn(math.MaxInt),
		Username: username,
		Conn:     conn,
	}
}

// handles the updates of the websocket
func (s *GameServer) handleWS(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}
	fmt.Print("New client trying to connect\n")
	// fmt.Print(conn)

	// ps := newPlayerSessions()
	sid := rand.Intn(math.MaxInt)
	pid := s.ctx.SpawnChild(newPlayerSession(sid, conn), fmt.Sprintf("session_%d", sid))
	s.sessions[pid] = struct{}{}
	fmt.Printf("client with sid %d and pid %s just connected\n", sid, pid)
}

func main() {

	e, _ := actor.NewEngine(actor.EngineConfig{})
	pid := e.Spawn(newGameServer, "server")
	fmt.Printf("PID: id %s, addy %s\n", pid.ID, pid.Address)
	select {}
}

// func login(conn *websocket.Conn, data types.Login) error {
// 	return conn.WriteJSON(data)
// }
