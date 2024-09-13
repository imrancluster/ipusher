package ipusher

import (
	"log"

	"github.com/gin-gonic/gin"
)

type IPusher struct{}

// Run function will return the gin.Enging
// with some routes for this package
func (ip IPusher) Use() *gin.Engine {
	return ip.setupRouter()
}

// Handle broadcasting messages
func (ip IPusher) HandleMessages() {
	for {
		msg := <-broadcast
		mutex.Lock()
		for client := range clients {
			// Broadcast message only to clients on the specified channel and with matching userID
			if client.channel == msg.Channel && client.userID == msg.UserID {
				err := client.conn.WriteJSON(msg)
				if err != nil {
					log.Println("Error broadcasting message:", err)
					client.conn.Close()
					delete(clients, client)
				}
			}
		}
		mutex.Unlock()
	}
}
