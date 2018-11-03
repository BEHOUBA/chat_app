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

var db *sql.DB
var err error

func init() {
	connStr := "postgres://txezbgusdiyzas:9dfe85b639ce49a0cce3e6bf6ee1f39de7d38090f3d3dfcc84364d4372511cf8@ec2-54-246-86-167.eu-west-1.compute.amazonaws.com:5432/dcut3u0egv6k8e"

	dbConfig := session.ManagerConfig{CookieName: "gosessionid", Gclifetime: 3600, ProviderConfig: connStr}

	GlobalSessions, err = session.NewManager("postgresql", &dbConfig)
	if err != nil {
		fmt.Println(err)
	}
	go GlobalSessions.GC()

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

}
