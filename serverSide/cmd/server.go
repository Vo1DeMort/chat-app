package main

import (
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)

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
	srv := &http.Server{
		Addr:         "",
		Handler:      app.routes(),
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 5,
	}

	// TODO: graceful shut down
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
