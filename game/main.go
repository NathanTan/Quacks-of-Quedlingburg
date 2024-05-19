package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
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

func (s *PlayerSession) handleMessage(msg *types.WSMessage) {
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

	case "PlayerState":
		var ps types.PlayerState

		if err := json.Unmarshal(msg.Data, &ps); err != nil {
			panic(err)
		}
		fmt.Println("PlayerState:")
		fmt.Println(ps)
	}

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
