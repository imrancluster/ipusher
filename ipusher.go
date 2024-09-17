package ipusher

import (
	"crypto/rand"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

type IPusher struct{}

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// Run function will return the gin.Enging
// with some routes for this package
func (ip IPusher) Use() *gin.Engine {

	ip.broadcastMessageForInactiveClient()

	return ip.setupRouter()
}

// broadcastMessageForInactiveClient for testing
// In future, we can use database or any others system that will make sure thate
// the inactive message should be tried to broadcast more than 3 times
// If the client not found then the message should be stored in cold storage
func (ip IPusher) broadcastMessageForInactiveClient() {
	// Set up cron job to fetch and broadcast messages periodically
	c := cron.New()
	c.AddFunc("@every 1m", func() { // Adjust the interval as needed
		log.Println("Running scheduled task: Broadcasting ")
		for key, inactiveMesage := range inActiveMessages {
			broadcast <- inactiveMesage

			log.Println("Key: ", key, " Msg: ", inactiveMesage)
			delete(inActiveMessages, key)
		}
	})
	c.Start()

	// Graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		log.Println("Shutting down...")
		c.Stop()
		os.Exit(0)
	}()
}

// Handle broadcasting messages
func (ip IPusher) HandleMessages() {

	for {
		msg := <-broadcast
		mutex.Lock()

		client := ip.findClilentByUserId(msg)
		if client.userID != 0 {
			err := client.conn.WriteJSON(msg)
			if err != nil {
				log.Println("Error broadcasting message:", err)
				client.conn.Close()
				delete(clients, client)
			}
		}

		mutex.Unlock()

		// Actice Client count
		log.Println("Active Client Count: ", len(clients))
	}
}

// findClilentByUserId helper function
func (ip IPusher) findClilentByUserId(msg BroadcastMessage) *Client {
	var c Client
	for client := range clients {
		// Broadcast message only to clients on the specified channel and with matching userID
		if client.channel == msg.Channel && client.userID == msg.UserID {
			c = *client
		}
	}

	if c.userID == 0 {
		c = Client{}
		// TODO: have to keep the message for future broadcasting | Database with try count
		inActiveMessages[ip.randomString(15)] = msg
		log.Println("Inactive Client: ", msg.UserID)
	}

	return &c
}

// RandomString returns a string of random characters of length n, using randomStringSource
// as the source for the string
func (ip IPusher) randomString(n int) string {
	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}
