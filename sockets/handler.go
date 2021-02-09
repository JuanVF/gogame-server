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
// if an error occurs it will delete the user from the connection
func (s *socketHandler) SendToAll(msg Message) error {
	for user, _ := range s.connections {
		err := user.WriteJSON(msg)

		if err != nil {
			delete(s.connections, user)

			return err
		}
	}

	return nil
}

// SendTo will send a message to a certain user
// if an error occurs it will delete the user from the connection
func (s *socketHandler) SendTo(msg Message, ws *websocket.Conn) error {
	if s.connections[ws] == nil {
		return fmt.Errorf("Socket Handler: %v does not exists...\n", ws)
	}

	err := ws.WriteJSON(msg)

	if err != nil {
		delete(s.connections, ws)
	}

	return err
}

// Send will send a message to various users
func (s *socketHandler) Send(msg Message, users []*websocket.Conn) (err error) {
	for _, user := range users {
		// We will send the message to all the users that exists
		// but this will throw an error anyways
		if s.connections[user] == nil {
			err = fmt.Errorf("Socket Handler: %v does not exists...\n", user)

			continue
		}

		err = user.WriteJSON(msg)

		if err != nil {
			delete(s.connections, user)

			return err
		}
	}

	return
}

// AddAction allow you to add a function to be executed by
// the user given a certain key
func (s *socketHandler) AddAction(key int, action Void) error {
	if s.actions[key] != nil {
		return fmt.Errorf("Socket Handler: Action already exists...\n")
	}

	s.actions[key] = action

	return nil
}

// AddActions allow you to add several actions to the handler
func (s *socketHandler) AddActions(actions map[int]Void) error {
	for key, action := range actions {
		err := s.AddAction(key, action)

		if err != nil {
			return err
		}
	}

	return nil
}

// RemoveAction allow you to quit an action
func (s *socketHandler) RemoveAction(key int) (action Void, err error) {
	if s.actions[key] == nil {
		return nil, fmt.Errorf("Socket Handler: Action doesnt exists...\n")
	}

	action = s.actions[key]

	delete(s.actions, key)

	return
}

// RemoveActions allow you to quit several actions
func (s *socketHandler) RemoveActions(keys []int) (actions []Void, err error) {
	actions = make([]Void, 0)

	for key, _ := range keys {
		action, err := s.RemoveAction(key)

		if err != nil {
			return nil, err
		}

		actions = append(actions, action)
	}

	return
}

// SetAction allow you to modify a certain action
func (s *socketHandler) SetAction(key int, action Void) error {
	if s.actions[key] == nil {
		return fmt.Errorf("Socket Handler: Action doesnt exists...\n")
	}

	s.actions[key] = action

	return nil
}
