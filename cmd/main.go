package main

import (
	"database/sql"
	"log"

	"github.com/Splucheviy/TiagoEcomm/cmd/api"
	"github.com/Splucheviy/TiagoEcomm/config"
	"github.com/Splucheviy/TiagoEcomm/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPass,
		Addr:                 config.Envs.DBAddr,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

// InitStorage...
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Sucessfully connected")
}
