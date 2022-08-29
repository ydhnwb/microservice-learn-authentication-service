package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const WEB_PORT = "8091"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service...")

	//Connect to DB
	//...

	//Setup
	app := Config{}
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", WEB_PORT),
		Handler: app.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Authentication service is started")

}
