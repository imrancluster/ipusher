package ipusher

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
var broadcast = make(chan BroadcastMessage)
var clients = make(map[*Client]bool)
var mutex sync.Mutex
var inActiveMessages = make(map[string]BroadcastMessage)

type Client struct {
	conn    *websocket.Conn
	channel string
	userID  int // Now dynamic userID
}

func (ip IPusher) setupRouter() *gin.Engine {
	r := gin.Default()

	h := handler{}

	ipusherGroup := r.Group("/api/v1")
	{
		// Endpoint to broadcast message to a specific channel
		ipusherGroup.POST("/broadcast", h.broadcastMessage)
		// WebSocket connection with dynamic userID
		ipusherGroup.GET("/ws/:channelName", h.websocketForClients)
	}

	return r
}

// Handle WebSocket connections with dynamic userID
func (ip IPusher) handleConnections(w http.ResponseWriter, r *http.Request, channelName string, userID int) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade WebSocket:", err)
		return
	}
	defer ws.Close()

	client := &Client{
		conn:    ws,
		channel: channelName,
		userID:  userID, // Now dynamically assigned
	}

	mutex.Lock()
	clients[client] = true
	mutex.Unlock()

	// Keep connection open
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}

	// Remove client on disconnect
	mutex.Lock()
	delete(clients, client)
	mutex.Unlock()
	log.Println("Client disconnected")
}
