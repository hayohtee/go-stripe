package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

func main()  {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		errorLog.Fatal(err)
	}

	var cfg config

	flag.IntVar(&cfg.port, "port", 4001, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment {development|production|maintenance}")
	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	app := application{
		cfg:           cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		version:       version,
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Fatal(err)
	}
}