package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
	"types"

	"github.com/gorilla/websocket"
)

const wsServerEndpoint = "ws://localhost:40000/ws"

// func (c *GameClient) Connect() error {
// 	c.conn
// }

//=======================

type GameClient struct {
	clientId int
	username string
	conn     *websocket.Conn
}

func newGameClient(conn *websocket.Conn, username string) *GameClient {
	return &GameClient{
		clientId: rand.Intn(math.MaxInt),
		username: username,
		conn:     conn,
	}
}

func (c *GameClient) login() error {
	fmt.Println("Attempting login")
	b, err := json.Marshal(types.Login{
		ClientId: c.clientId,
		Username: c.username,
	})
	if err != nil {
		return err
	}
	msg := types.WSMessage{
		Type: "Login",
		Data: b,
	} // He said this is bad

	jsonMsg, err := json.Marshal(msg)

	if err != nil {
		return err
	}
	// message := []byte("Hello from the client!")
	// err = c.conn.WriteMessage(websocket.TextMessage, message)

	fmt.Println("Send message")
	fmt.Println(msg)
	fmt.Println(c)
	return c.conn.WriteMessage(websocket.TextMessage, jsonMsg)
}

func main() {

	dialer := websocket.Dialer{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, _, err := dialer.Dial(wsServerEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := newGameClient(conn, "James")
	if err := c.login(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./html/index.html")
	})

	port := ":5000"
	fmt.Println("Server is running on port" + port)
	// Serve api /hi
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/pos", func(w http.ResponseWriter, r *http.Request) {
		sendGameMove(conn)
	})
	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))

}

func sendGameMove(conn *websocket.Conn) {
	for i := 0; i < 5; i++ {
		state := types.PlayerState{
			HP:       100,
			Position: types.Position{X: 5, Y: i}}

		b, err := json.Marshal(state)
		if err != nil {
			log.Fatal(err)
		}

		msg := types.WSMessage{
			Type: "PlayerState",
			Data: b,
		}

		// fmt.Println("Sending state")
		if err := conn.WriteJSON(msg); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Millisecond * 60 * 10 * 4)
	}
}
