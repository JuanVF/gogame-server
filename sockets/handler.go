package sockets

import (
	"fmt"
	"github.com/gorilla/websocket"
)

// Singleton object for socketHandler
var handler *socketHandler

// Void is a type to send void functions as arguments
type Void func()

// socketHandler will handle all the connections from
// users through websockets, also will contain all the actions
// can be performed by what the user sends
type socketHandler struct {
	connections map[*websocket.Conn]*Sockets
	actions     map[int]Void
}

// GetSocketHandlerInstance will return the instance of the socketHandler
// this will be a Singleton object
func GetInstance() *socketHandler {
	if handler == nil {
		handler = &socketHandler{
			connections: make(map[*websocket.Conn]*Sockets),
			actions:     make(map[int]Void),
		}
	}

	return handler
}

// AddConn will add a user to the handler
func (s *socketHandler) AddConn(ws *websocket.Conn) error {
	if s.connections[ws] != nil {
		return fmt.Errorf("User already exists...\n")
	}

	s.connections[ws] = NewSocket(ws)

	return nil
}

// RemoveConn will remove a user from the handler
func (s *socketHandler) RemoveConn(ws *websocket.Conn) error {
	if s.connections[ws] == nil {
		return fmt.Errorf("User does not exists...\n")
	}

	delete(s.connections, ws)

	return nil
}

// SendToAll will send a message to all the users
func (s *socketHandler) SendToAll() error {
	return fmt.Errorf("Not implemented yet\n")
}

// SendTo will send a message to a certain user
func (s *socketHandler) SendTo() error {
	return fmt.Errorf("Not implemented yet\n")
}

// Send will send a message to various users
func (s *socketHandler) Send() error {
	return fmt.Errorf("Not implemented yet\n")
}

// AddAction allow you to add a function to be executed by
// the user given a certain key
func (s *socketHandler) AddAction(key int, action Void) error {
	return fmt.Errorf("Not implemented yet\n")
}

// RemoveAction allow you to quit an action
func (s *socketHandler) RemoveAction(key int) (Void, error) {
	return nil, fmt.Errorf("Not implemented yet\n")
}

// SetAction allow you to modify a certain action
func (s *socketHandler) SetAction(key int, action Void) error {
	return fmt.Errorf("Not implemented yet\n")
}
