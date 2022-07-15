package main

import (
	"authentication/data"
	"database/sql"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct {
	db     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service...")

	app := Config{}

	server := &http.Server{
		Addr:    ":" + webPort,
		Handler: app.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
