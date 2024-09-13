package ipusher

// BroadcastMessage model to mange broadcast request body
type BroadcastMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	UserID  int    `json:"user_id"`
	Channel string `json:"channel"`
}
