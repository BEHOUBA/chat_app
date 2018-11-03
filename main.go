package main

import (
	"os"

	"github.com/astaxie/beego"
	_ "github.com/behouba/chat_app/models"
	_ "github.com/behouba/chat_app/routers"
)

func main() {
	beego.Run(getPort())
}

func getPort() string {
	if os.Getenv("PORT") != "" {
		return ":" + os.Getenv("PORT")
	}
	return ""
}
