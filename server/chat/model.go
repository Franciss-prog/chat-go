package chat

import "time"

type Message struct {
	ID        string    `json:"id"`
	Room      string    `json:"room"`
	SenderID  string    `json:"sender_id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type WSMessage struct {
	Type    string `json:"type"`
	Room    string `json:"room"`
	Content string `json:"content,omitempty"`
}

type WSResponse struct {
	Type      string `json:"type"`
	SenderID  string `json:"sender_id,omitempty"`
	Username  string `json:"username,omitempty"`
	Content   string `json:"content,omitempty"`
	Room      string `json:"room,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}
