package routers

import (
	"github.com/astaxie/beego"
	"github.com/behouba/chat_app/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
