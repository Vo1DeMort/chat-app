package main

import "github.com/jackc/pgx/v5"

type application struct {
	config config
	db     *pgx.Conn
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

func (app *application) runServer() error {
	//TODO this is where i left
}
