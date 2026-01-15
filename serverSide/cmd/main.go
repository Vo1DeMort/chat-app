package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	// TODO: recieve cmd parameters into cfg ,if provided
	// init air

	var cfg config
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	app := application{
		config: cfg,
		db:     conn,
	}

	err = app.runServer()
	if err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
