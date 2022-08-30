package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const WEB_PORT = "8091"

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service...")

	//Connect to DB
	conn := connectDatabase()
	if conn == nil {
		log.Panic("Cannot connect to database")
	}

	//Setup
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}
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

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectDatabase() *sql.DB {
	dsn := os.Getenv("DSN")
	for {
		conn, err := openDB(dsn)
		if err != nil {
			log.Println("Database is not ready yet")
			counts++
		} else {
			log.Println("Database connected!")
			return conn
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Waiting DB to be ready in  2 secs")
		time.Sleep(2 * time.Second)
		continue
	}
}
