package controllers

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

// var upgrader websocket.Upgrader

type ChatRoom struct {
	beego.Controller
}

func (this *ChatRoom) Get() {
	// Safe check.
	// uname := this.GetString("uname")
	// if len(uname) == 0 {
	// 	this.Redirect("/", 302)
	// 	return
	// }

	this.TplName = "chat-room.html"
	this.Data["IsWebSocket"] = true
}

func (c *ChatRoom) Join() {
	// uname := c.GetString("uname")
	// if len(uname) == 0 {
	// 	c.Redirect("/", 302)
	// 	return
	// }

	// Upgrade from http request to WebSocket.
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	// Join chat room.
	// Join(uname, ws)
	// defer Leave(uname)

	// Message receive loop.
	for {
		msgT, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		ws.WriteMessage(msgT, p)
		// publish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
	}
}
