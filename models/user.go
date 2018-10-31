package models

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// WebsocketConn *websocket.Conn
	// Connexions    []*websocket.Conn
}

func (u *User) Register() (err error) {

	// Register function check user's data validity,
	if u.Name == "" || u.Email == "" || u.Password == "" {
		err = errors.New("invalid user registration's data")
		return err
	}

	// hash user's password and store user's data on database
	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.Password = string(cryptedPassword)
	err = db.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING user_id", u.Name, u.Email, u.Password).Scan(&u.ID)
	if err != nil {
		return err
	}
	return
}

func (u *User) Login() (err error) {
	var userFromDB User
	// check for null values
	if u.Email == "" || u.Password == "" {
		err = errors.New("invalid user login's data")
		return
	}
	// try to find user in database
	row := db.QueryRow("SELECT user_id, name, email, password FROM users WHERE email=$1", u.Email)

	err = row.Scan(&userFromDB.ID, &userFromDB.Name, &userFromDB.Email, &userFromDB.Password)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(u.Password))
	if err != nil {
		return
	}
	*u = userFromDB
	return
}

// set empty user's fields from user id
func (u *User) GetDataFromDB(ID int) (err error) {
	err = db.QueryRow("SELECT * FROM users WHERE user_id=$1", ID).Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return
	}
	return
}

func (u *User) GetChatMessages(receiver *User) (msg []Message, err error) {
	rows, err := db.Query("SELECT sender_id ,message, createdat FROM messages WHERE (sender_id=$1 AND receiver_id=$2) OR (sender_id=$3 AND receiver_id=$4)", u.ID, receiver.ID, receiver.ID, u.ID)
	if err != nil {
		return
	}
	for rows.Next() {
		var m Message
		err := rows.Scan(&m.SenderID, &m.Body, &m.Date)
		if err != nil {
			log.Println(err)
			continue
		}
		msg = append(msg, m)
	}
	return
}

func GetAllUsers(exclude int) (users []User, err error) {
	rows, err := db.Query("SELECT user_id, name, email FROM users where user_id!=$1", exclude)
	if err != nil {
		return
	}
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, u)
	}
	return
}
