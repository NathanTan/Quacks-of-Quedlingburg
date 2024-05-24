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

	"github.com/gin-gonic/gin"
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

	// Serve React app
	// r.StaticFS("/", http.Dir("/src/public/index.html"))
	// Set the router as the default one shipped with Gin
	// router := gin.Default()

	// Serve frontend static files
	// router.Use(static.Serve("/", static.LocalFile("./src/public", true)))
	r := gin.Default()
	// Serve React app
	// r.StaticFS("/static/public", http.Dir("../src/public")) // Serve static files under the /static route
	r.StaticFS("/static", http.Dir("../dist")) // Serve static files from the dist directory under the /static route

	// Serve static files from the dist/public directory
	r.Static("/public", "./dist/public")

	// r.StaticFS("/static/public", http.Dir("../src/public")) // Serve static files from the dist directory under the /static route
	// r.Static("/static", "../src/public") // Serve static files from the dist directory under the /static route
	r.LoadHTMLGlob("../dist/index.html") // Load HTML files

	// // Create a file server for serving static files from the dist directory
	// distFileServer := http.FileServer(http.Dir("../dist"))

	// // Create a file server for serving static files from the src/public directory
	// publicFileServer := http.FileServer(http.Dir("../src/public"))

	// // Serve static files from the dist directory under the /static/dist route
	// r.GET("/static/dist/*filepath", gin.WrapH(distFileServer))

	// // Serve static files from the src/public directory under the /static/public route
	// r.GET("/static/public/*filepath", gin.WrapH(publicFileServer))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil) // Serve index.html file on accessing root route
	})

	r.POST("/move", func(c *gin.Context) {
		fmt.Println("Move received")
		sendGameMove(conn)
	})

	port := ":3000"

	// Start the server on port 5000
	r.Run(port)

	fmt.Println("Server is running on port" + port)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./html/index.html")
	// })

	// Serve api /hi
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})
}

func sendGameMove(conn *websocket.Conn) {
	fmt.Println("Sending move")
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
