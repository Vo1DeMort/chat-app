package main

import (
	"os"
	"os/signal"
	"syscall"
)

func (app *application) serve() error {
	// to recieve error from graceful shutdown
	gsShutdownError := make(chan error)

	// goroutin to recieve graceful shutdown
	go func() {
		// buffered channel
		quit := make(chan os.Signal, 1)

		// catch the shutdown cmd
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		signal := <-quit
		app.logger.PrintInfo("caught signal", map[string]string{
			"signal": signal.String(),
		})

	}()

	return nil
}
