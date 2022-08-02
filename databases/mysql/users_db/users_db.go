package users_db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB

	username = os.Getenv("DBUSER")
	password = os.Getenv("DBPASS")
	host     = os.Getenv("DBADDR")
	schema   = os.Getenv("DBSCHEMA")
)

func init() {
	config := mysql.Config{
		User:                 username,
		Passwd:               password,
		Addr:                 host,
		DBName:               schema,
		Net:                  "tcp",
		AllowNativePasswords: true,
	}
	var err error

	Client, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		panic(err)
	}
	pingErr := Client.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	// mysql.SetLogger(Client)
	log.Println("Connected to db!")
}
