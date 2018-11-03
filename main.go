package main

import (
	"os"

	"github.com/astaxie/beego"
	_ "github.com/behouba/chat_app/models"
	_ "github.com/behouba/chat_app/routers"
)

func main() {
	// port, err := strconv.Atoi(os.Getenv("PORT"))
	// if err == nil {
	// 	beego.HttpPort = port
	// }
	beego.Run(":" + os.Getenv("PORT"))
}
