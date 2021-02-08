package sockets

import "github.com/gorilla/websocket"

// Sockets is a structure to save all user connection
// related info
type Sockets struct {
	user *websocket.Conn
}

// NewSocket is the constructor for Sockets
func NewSocket(user *websocket.Conn) *Sockets {
	return &Sockets{
		user: user,
	}
}
