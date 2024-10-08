package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"quacks"
	"strconv"
	"sync"
	"time"
	"types"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const wsServerEndpoint = "ws://localhost:40000/ws"

type GameClient struct {
	clientId int
	username string
	conn     *websocket.Conn
}

var (
	gameStates   = make(map[string]*quacks.GameState)
	playerStates = make(map[string]*types.PlayerState)
	mutex        = &sync.Mutex{}
	player_mutex = &sync.Mutex{}
)

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

func newGameClient(conn *websocket.Conn, username string) *GameClient {
	return &GameClient{
		clientId: rand.Intn(math.MaxInt),
		username: username,
		conn:     conn,
	}
}

func getPlayerStateFromServer(c *GameClient, playerId int) error {
	fmt.Println("Sending request to get player state for player ", playerId)
	msg := types.WSMessage{
		Type: "PlayerState",
		Data: nil, // TODO spend player id in json payload
	}

	return sendMessageToGameServer(c, msg)
}

func createNewGameStateFromServer(c *GameClient, gameId string) error {
	fmt.Println("Sending request to create a new game state")

	b, err := json.Marshal(types.NewGameStateRequest{
		AuthToken:   "123",
		PlayerNames: []string{"Nathan", "Leah", "Raymond", "Hannah"},
	})

	if err != nil {
		return err
	}
	msg := types.WSMessage{
		Type: "NewGameStateRequest",
		Data: b,
	}

	return sendMessageToGameServer(c, msg)
}

func getGameStateFromServer(c *GameClient, gameId string) error {
	fmt.Println("Sending request to get game state")

	b, err := json.Marshal(types.GameStateRequest{
		AuthToken: "123",
		GameId:    gameId,
	})
	if err != nil {
		return err
	}
	msg := types.WSMessage{
		Type: "GameStateRequest",
		Data: b,
	}

	return sendMessageToGameServer(c, msg)
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
	}

	return sendMessageToGameServer(c, msg)
}

func sendMessageToGameServer(c *GameClient, msg types.WSMessage) error {

	jsonMsg, err := json.Marshal(msg)

	if err != nil {
		return err
	}

	fmt.Println("Send message")
	fmt.Println(msg)
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

	r.GET("/", func(gtinContext *gin.Context) {
		gtinContext.HTML(http.StatusOK, "index.html", nil) // Serve index.html file on accessing root route
	})

	r.GET("/requestPlayerState/:id", func(gtinContext *gin.Context) {
		// Extract the player ID from the route parameter
		playerID := gtinContext.Param("id")
		// Convert playerID to the appropriate type if necessary, e.g., to int
		playerIDInt, err := strconv.Atoi(playerID)

		if err != nil {
			// Handle error, maybe return an HTTP error response
			gtinContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid player ID"})
			return
		}

		fmt.Println("Received GET request for player state for player ", playerIDInt)

		getPlayerStateFromServer(c, playerIDInt) // TODO: Get player ID from request
		data := getPlayerState("game1")
		if gtinContext != nil {
			gtinContext.JSON(http.StatusOK, data)
		}
	})

	r.POST("/requestState", func(gtinContext *gin.Context) {
		// gtinContext.HTML(http.StatusOK, "index.html", nil) // Serve index.html file on accessing root route
		createNewGameStateFromServer(c, "123")
		// if gtinContext != nil {
		// 	gtinContext.JSON(http.StatusOK, data)
		// }
	})

	r.POST("/getState/:id", func(gtinContext *gin.Context) {
		gameId := gtinContext.Param("id")

		fmt.Println("Get State for Game: ", gameId)
		data := getGameState(gameId)
		fmt.Println("Gotten State Data: ", data)
		if gtinContext != nil {
			gtinContext.JSON(http.StatusOK, data)
		}
	})

	r.POST("/move", func(gtinContext *gin.Context) {
		fmt.Println("Move received")

		var playerMove types.PlayerMove
		if err := gtinContext.BindJSON(&playerMove); err != nil {
			// Handle error, maybe return a bad request status to the client
			gtinContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("Player Move: ", playerMove)

		playerMove.AuthToken = "123"
		playerMove.GameId = "game123"
		playerMove.PlayerId = 1

		sendGameMove(conn, playerMove)

	})

	port := ":3000"

	go readLoop(c) // TODO: Fix the connection so that it isn't the same for every client

	// Start the server on port 5000
	r.Run(port)

	fmt.Println("Server is running on port" + port)

}

func sendGameMove(conn *websocket.Conn, move types.PlayerMove) {
	fmt.Println("Sending move")

	b, err := json.Marshal(move)
	if err != nil {
		log.Fatal(err)
	}

	msg := types.WSMessage{
		Type: "PlayerMove",
		Data: b,
	}

	if err := conn.WriteJSON(msg); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Millisecond * 5)

}

func readLoop(c *GameClient) {
	var msg *types.WSMessage
	// i := 0
	fmt.Printf("Reading messages\n")
	for {
		_, payload, err := c.conn.ReadMessage()
		if err != nil {
			fmt.Println("read error", err)
			fmt.Println("Problem Message:")
			fmt.Println(msg)
			return
		}

		if payload == nil {
			fmt.Println("Payload is nil")
			return
		}

		var receivedMsg types.WSMessage

		err = json.Unmarshal(payload, &receivedMsg)
		if err != nil {
			fmt.Println("read error", err)
			fmt.Println("Problem Message:")
			fmt.Println(msg)
			return
		}
		fmt.Println(receivedMsg.Type)
		go handleMessage(&receivedMsg)
	}
}

func handleMessage(msg *types.WSMessage) {
	fmt.Printf("Handling message - type: %s\n", msg.Type)
	switch msg.Type {

	case "PlayerState":
		var ps types.PlayerState
		if err := json.Unmarshal(msg.Data, &ps); err != nil {
			panic(err)
		}
		fmt.Println("New PlayerState:")
		fmt.Println(ps)

		// TODO: Save this by player ID and make the sender send playerId also
		savePlayerState("game1", &ps)

	case "NewGameState":
		fmt.Println("Requesting New Game State")
		receiveGameState(msg)

	case "GameState":
		fmt.Println("Requesting Game State")
		receiveGameState(msg)

	}
}

func receiveGameState(msg *types.WSMessage) {
	fmt.Println("Recieved game state")
	var state quacks.GameState

	if err := json.Unmarshal(msg.Data, &state); err != nil {
		panic(err)
	}
	fmt.Println("GameState:")
	state.PrintGameStateForDebugging()

	// Get the awaiting input for the front end without a refactor
	if state.Awaiting != nil {
		state.FrontEndAwaiting = *state.Awaiting
	}

	saveGameState(state.Id, &state)
}
