package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"projectx.io/drivethru/cmd/api"
	"projectx.io/drivethru/db"
)

// TODO :
// Logging
// Configuration

func main() {

	log.Println("PXDT - Menu service")

	db := connectDatabase()
	startServer(db)

}

func startServer(db *sql.DB) {
	server := api.NewAPIServer(":8081", db)
	errServ := server.Run()
	if errServ != nil {
		log.Fatal("Could not launch server ..", errServ)
	}
}

func connectDatabase() *sql.DB {

	log.Printf("Connecting to %s as %s \n", "127.0.0.1:3306", "root")

	db, _ := db.MySqlStorage(mysql.Config{
		User:                 "samin",
		Passwd:               "1",
		Addr:                 "127.0.0.1:3306",
		DBName:               "pxdt_menu",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	pingDb(db)

	return db
}

func pingDb(db *sql.DB) {

	err := db.Ping()
	if err != nil {
		log.Fatal("Falied to ping db")
	}
	log.Println("Database connected !")
}
