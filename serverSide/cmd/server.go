package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
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

// might need refactoring
var wg sync.WaitGroup

func (app *application) runServer() error {
	srv := &http.Server{
		Addr:         "",
		Handler:      app.routes(),
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 5,
	}

	//to recieve error from graceful shutdown if any
	gsdError := make(chan error)

	go func() {
		// recieves os signal for graceful shutdown
		osSignal := make(chan os.Signal, 1)

		// NOTE:
		// listen for SIGINT and SIGTERM
		// SIGINT : signal interrupt, caused by ctrl c
		// SIGTERM : signal termination , triggered by something like kill cmd
		signal.Notify(osSignal, syscall.SIGINT, syscall.SIGTERM)

		// NOTE:
		// until the value from osSignal is released , this goroutine will be blocked
		sig := <-osSignal

		log.Println("System signal received:", sig.String())

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			gsdError <- err
		}

		log.Println("Completing Background Tasks")
		wg.Wait()
		gsdError <- nil

	}()

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	//need to learn the connection between 2 goroutin
	err = <-gsdError
	if err != nil {
		return err
	}

	fmt.Println("Server has stopped")

	return nil
}
