package types

import "github.com/gorilla/websocket"

type Login struct {
	ClientId int    `json: "clientID"`
	Username string `json: "userName"`
}

type WSMessage struct {
	Type string `json: "type"`
	Data []byte `json: "data"`
}

type GameClient struct {
	ClientId int
	Username string
	Conn     *websocket.Conn
}
type Position struct {
	X int `json: "x"`
	Y int `json: "y"`
}

// Player state
type PlayerState struct {
	HP       int `json: "HP"`
	Position Position
}
