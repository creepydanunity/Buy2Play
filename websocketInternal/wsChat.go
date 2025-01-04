package websocketInternal

import (
	"github.com/gorilla/websocket"
	"sync"
)

var (
	mu      sync.Mutex
	clients = make(map[uint]map[*websocket.Conn]bool) // conversationID -> clients
)

// AddClient adds a WebSocket connection for a user in a conversation
func AddClient(conversationID uint, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()

	if clients[conversationID] == nil {
		clients[conversationID] = make(map[*websocket.Conn]bool)
	}

	clients[conversationID][conn] = true
}

// RemoveClient removes a WebSocket connection for a user in a conversation
func RemoveClient(conversationID uint, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()

	if _, ok := clients[conversationID]; ok {
		delete(clients[conversationID], conn)
	}
}

// BroadcastMessage sends a message to all clients in a specific conversation
func BroadcastMessage(conversationID uint, message []byte) {
	mu.Lock()
	defer mu.Unlock()

	for conn := range clients[conversationID] {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			conn.Close()
			delete(clients[conversationID], conn)
		}
	}
}
