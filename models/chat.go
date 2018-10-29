package models

import (
	"log"
	"time"
)

type Message struct {
	SenderID   int    `json:"sender_id"`
	ReceiverID int    `json:"receiver_id"`
	Body       string `json:"body"`
	Date       time.Time
}

func (m *Message) StoreAndSend(users []User) (err error) {
	err = db.QueryRow("INSERT INTO messages (sender_id, message , receiver_id ) VALUES ($1, $2,$3)  RETURNING createdat", m.SenderID, m.Body, m.ReceiverID).Scan(&m.Date)
	if err != nil {
		return
	}
	for _, u := range users {
		// break the loop when two user have received the messages
		if u.ID == m.SenderID || u.ID == m.ReceiverID {
			if err := u.WebsocketConn.WriteJSON(&m); err != nil {
				log.Println(err)
			}
		}
	}
	return
}
