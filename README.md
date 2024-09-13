## ipusher
Package ipusher will help to push any message to multiple clients through socket connection.

### Install
```
go get -u github.com/imrancluster/ipusher
```

### Examples
Let's start registering a default routes for pusher:
```go
func main() {

	// Instantiate the IPusher
	ipusher := ipusher.IPusher{}

	// Used build-in two routes from IPusher
	router := ipusher.Use()

	// Add IPusher HandleMessage Func using Go Routine
	go ipusher.HandleMessages()

	router.Run(":8088")
}
```

### Clints
The client will connect using the following API
```
https//example.com:8088/api/v1/ws/[CHANNEL_NAME]?user_id=[USER_ID]
```

### Testing
```bash
curl -X POST http://example.com:8088/api/v1/broadcast \
-H "Content-Type: application/json" \
-d '{"type": "notification", "message": "Test message for user 100", "user_id": 100, "channel": "channelOne"}'
```