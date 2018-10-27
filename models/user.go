package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register function check if user's data validity,
// hash user's password and store user's data on database
func (u *User) Register() (err error) {
	if u.Name == "" || u.Email == "" || u.Password == "" {
		err = errors.New("invalid user registration's data")
		return err
	}

	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.Password = string(cryptedPassword)
	_, err = db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}
	return
}

func (u *User) Login() (err error) {
	var userFromDB User
	if u.Email == "" || u.Password == "" {
		err = errors.New("invalid user login's data")
		return
	}
	row := db.QueryRow("SELECT name, email, password FROM users WHERE email=$1", u.Email)

	err = row.Scan(&userFromDB.Name, &userFromDB.Email, &userFromDB.Password)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(u.Password))
	if err != nil {
		return
	}
	return
}
