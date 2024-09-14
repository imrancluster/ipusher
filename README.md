## IPusher
The `ipusher` package helps you push messages to multiple clients through a socket connection.

### Installation
To install the package, run the following command:
```bash
go get -u github.com/imrancluster/ipusher
```

### Example
Here’s how to register default routes for IPusher:
```go
func main() {
	// Instantiate the IPusher
	ipusher := ipusher.IPusher{}

	// Use built-in routes from IPusher
	router := ipusher.Use()

	// Start IPusher's HandleMessages function in a Go routine
	go ipusher.HandleMessages()

	// Run the router on port 8088
	router.Run(":8088")
}
```

### Clients
Clients can connect using the following API:
```
http://localhost:8088/api/v1/ws/[CHANNEL_NAME]?user_id=[USER_ID]
```

### Testing
To broadcast a message, you can use the following `curl` command:
```bash
curl -X POST http://localhost:8088/api/v1/broadcast \
-H "Content-Type: application/json" \
-d '{"type": "notification", "message": "Test message for user 100", "user_id": 100, "channel": "channelOne"}'
```

### VueJS Example
**Install Dependencies**
```bash
npm install vue-socket.io
```

```js
<template>
    <div>
      <h1>Real-time Messages for User ID: 100</h1>
      <ul>
        <li v-for="(msg, index) in messages" :key="index">{{ msg.message }}</li>
      </ul>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        messages: [],
        userId: 100, // Hardcoded for testing
        channelName: 'channelOne' // Change this value based on the page
      };
    },
    mounted() {
      this.setupWebSocket();
    },
    methods: {
      setupWebSocket() {
        const ws = new WebSocket(`ws://localhost:8088/api/v1/ws/${this.channelName}?user_id=${this.userId}`);
  
        ws.onopen = () => {
          console.log("WebSocket connection established on", this.channelName);
        };
  
        ws.onmessage = (event) => {
          const data = JSON.parse(event.data);
  
          // Check if the message is for the current user
          if (data.user_id === this.userId) {
            console.log("Data:", data)
            this.messages.push(data);
          }
        };
  
        ws.onerror = (error) => {
          console.error("WebSocket error:", error);
        };
  
        ws.onclose = (event) => {
          console.log("WebSocket connection closed", event);
        };
      }
    }
  };
  </script>
  
  <style>
  ul {
    list-style-type: none;
  }
  </style>
  
```

### Test with a Simple WebSocket Client
**Install wscat:**
```bash
npm install -g wscat
```

**Connect to your WebSocket server:**
```bash
wscat -c ws://localhost:8088/api/v1/ws/[CHANNEL_NAME]?user_id=[USER_ID]
```