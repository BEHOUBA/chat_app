package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/behouba/chat_app/models"
	"github.com/gorilla/websocket"

	"github.com/astaxie/beego"
)

type ChatRoom struct {
	beego.Controller
}

func (c *ChatRoom) Get() {
	if c.GetSession("session") == nil {
		c.Redirect("/", 303)
		return
	}
	userID := c.GetSession("session").(int)

	users, err := models.GetAllUsers(userID)
	if err != nil {
		log.Println(err)
	}

	c.Data["users"] = users
	c.Data["IsWebSocket"] = true
	c.TplName = "chat-room.html"
}

type Channel struct {
	beego.Controller
}

func (c *Channel) Get() {
	if c.GetSession("session") == nil {
		c.Redirect("/", 303)
		return
	}
	var sender, receiver models.User
	sender.ID = c.GetSession("session").(int)
	receiver.ID, _ = strconv.Atoi(c.Ctx.Input.Param(":id"))

	messages, err := sender.GetChatMessages(&receiver)
	if err != nil {
		log.Println(err)
	}
	receiver.GetDataFromDB(receiver.ID)
	c.Data["receiver"] = receiver
	c.Data["messages"] = messages
	c.TplName = "channel.html"
}

type ChatWebSocket struct {
	beego.Controller
}

func (c *ChatWebSocket) Get() {

	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	var user models.User

	// connect user to the websocket connection
	user.ID = c.GetSession("session").(int)

	// add user's connection to the active connection Hub
	activeConn.AddConn(user.ID, ws)

	log.Println("there are ", len(activeConn.ActiveUsersConn), " online")

	c.ServeJSON()

	// Message receive loop.
	for {
		var msg models.Message
		msg.SenderID = user.ID
		if err := ws.ReadJSON(&msg); err != nil {
			activeConn.DeleteConn(user.ID, ws)
			log.Println(err)
			log.Println("one user let the chat there are now ", len(activeConn.ActiveUsersConn), " user(s) online")
			return
		}
		log.Println("received msg", msg)
		if err := msg.Send(&activeConn); err != nil {
			log.Println(err)
			return
		}

	}
}
