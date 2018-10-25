package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/astaxie/beego/session/postgres"

	"github.com/astaxie/beego/session"
	_ "github.com/lib/pq"
)

var GlobalSessions *session.Manager

var Db *sql.DB
var err error

func init() {
	connStr := "postgresql://localhost/chat_app?user=postgres"

	dbConfig := session.ManagerConfig{CookieName: "gosessionid", Gclifetime: 3600, ProviderConfig: connStr}

	GlobalSessions, err = session.NewManager("postgresql", &dbConfig)
	if err != nil {
		fmt.Println(err)
	}
	go GlobalSessions.GC()

	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

}
