package types

import "github.com/gorilla/websocket"

type Login struct {
	ClientId int    `json:"clientID"`
	Username string `json:"userName"`
}

type NewGameStateRequest struct {
	AuthToken   string   `json:"authToken"`
	PlayerNames []string `json:"playerNames"`
}

type GameStateRequest struct {
	AuthToken string `json:"authToken"`
	GameId    string `json:"gameId"`
}

type PlayerMove struct {
	AuthToken string `json:"authToken"`
	GameId    string `json:"gameId"`
	PlayerId  int    `json:"playerId"`
	Move      string `json:"move"`
}

type WSMessage struct {
	Type string `json:"type"`
	Data []byte `json:"data"`
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

type PlayerStateRequest struct {
	PlayerId int `json: "playerId"`
}
