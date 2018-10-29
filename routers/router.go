package routers

import (
	"github.com/astaxie/beego"
	"github.com/behouba/chat_app/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.Login{})
	beego.Router("/register", &controllers.Register{})
	beego.Router("/logout", &controllers.Logout{})
	beego.Router("/room", &controllers.ChatRoom{})
	beego.Router("/ws/chat", &controllers.ChatWebSocket{})
}
