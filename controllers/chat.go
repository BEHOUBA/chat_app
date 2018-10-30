package controllers

import (
	"log"
	"strconv"

	"github.com/behouba/chat_app/models"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var activeUsers []models.User

type ChatRoom struct {
	beego.Controller
}

func (c *ChatRoom) Get() {
	userID := c.GetSession("session").(int)
	if userID == 0 {
		c.Redirect("/", 303)
		return
	}

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
	var sender, receiver models.User
	sender.ID = c.GetSession("session").(int)
	receiver.ID, _ = strconv.Atoi(c.Ctx.Input.Param(":id"))

	messages, err := sender.GetChatMessages(&receiver)
	if err != nil {
		log.Println(err)
	}
	receiver.GetDataFromDB(receiver.ID)
	c.Data["username"] = receiver.Name
	c.Data["messages"] = messages
	c.TplName = "channel.html"
}

type ChatWebSocket struct {
	beego.Controller
}

func (c *ChatWebSocket) Get() {

	// Upgrade from http request to WebSocket.
	ws, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	var user models.User

	// connect user to the websocket connection
	user.ID = c.GetSession("session").(int)
	err = user.Connect(ws, user.ID)
	if err != nil {
		log.Println(err)
		return
	}

	// add user to active user list
	activeUsers = append(activeUsers, user)
	log.Println("there are ", len(activeUsers), " online")

	// Message receive loop.
	for {
		var msg models.Message
		msg.SenderID = user.ID
		if err := ws.ReadJSON(&msg); err != nil {
			delUserFromActiveList(ws, user)
			log.Println(err)
			return
		}
		log.Println("received msg", msg)
		if err := msg.StoreAndSend(activeUsers); err != nil {
			log.Println(err)
			return
		}
		// ws.WriteMessage(msgT, p)
		// publish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
	}
}

// remove user from active user list
func delUserFromActiveList(ws *websocket.Conn, user models.User) {
	for i, u := range activeUsers {
		if u.WebsocketConn == ws {
			activeUsers = append(activeUsers[:i], activeUsers[i+1:]...)
			log.Println("user removed from active user's list")
			return
		}
	}
}
