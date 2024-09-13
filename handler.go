package ipusher

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// handler struct to orgazine the all handlers function
type handler struct{}

// broadcastMessage func is a REST API to braodcast any message using BroadcastMessage model
func (h handler) broadcastMessage(c *gin.Context) {
	var msg BroadcastMessage
	if err := c.BindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid message format"})
		return
	}
	log.Println("Broadcasting message:", msg)
	broadcast <- msg
	c.JSON(http.StatusOK, gin.H{"status": "Message broadcasted"})
}

// websocketForClients func is a REST API to connect with clients
func (h handler) websocketForClients(c *gin.Context) {
	ip := IPusher{}

	channelName := c.Param("channelName")
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	ip.handleConnections(c.Writer, c.Request, channelName, userID)
}
