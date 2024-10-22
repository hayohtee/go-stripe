package main

import (
	"flag"
	"html/template"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	version    = "1.0.0"
	cssVersion = "1"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		errorLog.Fatal(err)
	}

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment {development|production}")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL to API")
	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	app := application{
		cfg:           cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: make(map[string]*template.Template),
		version:       version,
	}

	err = app.serve()
	if err != nil {
		app.errorLog.Fatal(err)
	}

}
