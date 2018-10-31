package models

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// Hub struct keep track of all connected user
type Hub struct {
	// ActiveUsersConn map user's id with array of user's websocket connection
	ActiveUsersConn map[int][]*websocket.Conn
}

// add new connected user to Hub
func (h *Hub) AddConn(uID int, conn *websocket.Conn) {
	h.ActiveUsersConn[uID] = append(h.ActiveUsersConn[uID], conn)
}

// DeleteConn take user's id and websocket connection and delete what connection
func (h *Hub) DeleteConn(uID int, conn *websocket.Conn) {
	// delete use from map if he has only one websocket connection
	if len(h.ActiveUsersConn[uID]) == 1 {
		delete(h.ActiveUsersConn, uID)
		// close the connection
		conn.Close()
		return
	}
	// remote the current websocket connection from user's active connection's list
	for i, c := range h.ActiveUsersConn[uID] {
		if c == conn {
			h.ActiveUsersConn[uID] = append(h.ActiveUsersConn[uID][:i], h.ActiveUsersConn[uID][i+1:]...)

			// close the connection
			conn.Close()
			return
		}
	}
	return
}

type Message struct {
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Body       string `json:"body"`
	Date       time.Time
}

func (m *Message) Send(h *Hub) (err error) {
	err = db.QueryRow("INSERT INTO messages (sender_id, message , receiver_id ) VALUES ($1, $2,$3)  RETURNING createdat", m.SenderID, m.Body, m.ReceiverID).Scan(&m.Date)
	if err != nil {
		return
	}

	// should try to make these two bottom's for loop on two goroutines

	// send the message to receiver actives connections
	for _, c := range h.ActiveUsersConn[m.ReceiverID] {
		if err := c.WriteJSON(&m); err != nil {
			log.Println(err)
		}
	}
	// send the message to sender actives connections
	for _, c := range h.ActiveUsersConn[m.SenderID] {
		if err := c.WriteJSON(&m); err != nil {
			log.Println(err)
		}
	}
	return
}
