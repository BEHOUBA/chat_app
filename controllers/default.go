package controllers

import (
	"encoding/json"
	"log"
	"regexp"

	"github.com/astaxie/beego"
	"github.com/behouba/chat_app/models"
)

var (
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["isAuth"] = c.GetSession("session")
	log.Println(c.GetSession("session"))
	c.TplName = "index.html"
}

type Login struct {
	beego.Controller
}

func (c *Login) Get() {
	if c.GetSession("session") != nil {
		c.Redirect("/", 303)
		return
	}
	c.TplName = "login.html"
}

func (c *Login) Post() {
	var user models.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		// log.Println(err)
		c.Abort("500")
	}

	if err := user.Login(); err != nil {
		// log.Println(err)
		switch err.Error() {
		case "sql: no rows in result set":
			c.Abort("401")
			break
		case "crypto/bcrypt: hashedPassword is not the hash of the given password":
			c.Abort("400")
			break
		default:
			c.Abort("500")
			break
		}
	}
	c.SetSession("session", user.ID)
	c.ServeJSON()
}

type Register struct {
	beego.Controller
}

func (c *Register) Get() {
	if c.GetSession("session") != nil {
		c.Redirect("/", 303)
		return
	}
	c.TplName = "register.html"
}

func (c *Register) Post() {

	// check if user already have an active session
	if c.GetSession("session") != nil {
		c.CustomAbort(403, "error")
	}
	var user models.User

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		log.Println(err)
	}
	// validate user email
	if !emailRegexp.MatchString(user.Email) {
		c.Abort("400")
	}

	if err := user.Register(); err != nil {
		log.Println(err)
		switch err.Error() {
		case "invalid user registration's data":
			c.Abort("400")
			break
		case `pq: duplicate key value violates unique constraint "users_email_key"`:
			c.Abort("409")
			break
		default:
			c.Abort("500")
		}
	}
	c.SetSession("session", user.ID)
	c.ServeJSON()
}

type Logout struct {
	beego.Controller
}

func (c *Logout) Get() {
	c.DelSession("session")
	c.Redirect("/", 303)
}
