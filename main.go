package main

import (
	"github.com/astaxie/beego"
	_ "github.com/behouba/chat_app/models"
	_ "github.com/behouba/chat_app/routers"
)

func main() {
	beego.Run()
}
