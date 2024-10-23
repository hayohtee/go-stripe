package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type application struct {
	cfg           config
	infoLog       *log.Logger
	errorLog      *log.Logger
	version       string
}

func (app *application) serve() error {
	srv := http.Server{
		Addr:              fmt.Sprintf(":%d", app.cfg.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Printf("Starting backend server in %s mode on port %d\n", app.cfg.env, app.cfg.port)

	return srv.ListenAndServe()
}