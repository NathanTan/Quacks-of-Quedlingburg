package types

import "golang.org/x/net/websocket"

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
