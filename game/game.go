package game

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

type PlayerSession struct {
	sessionId int
	username  string
	clientId  string
	inLobby   bool
	conn      *websocket.Conn
}

type GameServer struct {
	ctx      *actor.Context
	sessions map[*actor.PID]struct{}
}

func newPlayerSession(sid int, conn *websocket.Conn) actor.Producer {
	return func() actor.Receiver {
		return &PlayerSession{
			sessionId: sid,
			clientId:  clientId,
			username:  username,
			inLobby:   true,
			conn:      conn,
		}
	}
}

func (s *PlayerSession) Recieve(c *actor.Context) {
	switch c.Message().(type) {
	case actor.Started:
		s.readLoop()
		// s.ctx = c
		// _ = msg
	}

}

func (s *PlayerSession) readLoop() {
	var msg types.WSMessage
	for {
		if err := s.conn.ReadJSON(msg); err != nil {
			fmt.Println("read error", err)
			return
		}
		go s.handleMessage(msg)
	}
}

func (s *PlayerSession) handleMessage(msg types.WSMessage) {
	switch msg.Type {
	case "Login":
		var loginMsg types.Login
		if err := json.Unmarshal(msg.Data, &loginMsg); err != nil {
			panic(err) // Panic since this is a single player session, not in the server
		}
		s.clientId = loginMsg.ClientId
		s.username = loginMsg.Username
		fmt.Println(loginMsg)
	}

}

func newGameServer() actor.Receiver {
	return &GameServer{
		sessions: make(map[*actor.PID]struct{}),
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

func (s *GameServer) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		s.startHTTP()
		s.ctx = c
		_ = msg
	}
}

// handles the updates of the websocket
func (s *GameServer) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}
	fmt.Print("New client trying to connect")
	fmt.Print(conn)

	// ps := newPlayerSessions()
	sid := rand.Intn(math.MaxInt)
	pid := s.ctx.SpawnChild(newPlayerSession(sid, conn), fmt.Sprintf("session_%d", sid))
	s.sessions[pid] = struct{}{}
	fmt.Printf("client with sid %d and pid %s just connected\n", sid, pid)
}

func RunGameServer() {

	e, _ := actor.NewEngine(actor.EngineConfig{})
	pid := e.Spawn(newGameServer, "server")
	fmt.Printf("PID: id %s, addy %s\n", pid.ID, pid.Address)
	select {}
}

func login(conn *websocket.Conn, data types.Login) error {
	return conn.WriteJSON(data)
}
