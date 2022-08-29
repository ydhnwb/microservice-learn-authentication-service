package main

import "database/sql"

const WEB_PORT = "8091"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

}
