package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	version = "1.0.0"
	cssVersion = "1"
)

func main()  {
	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorlog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	
	err := godotenv.Load()
	if err != nil {
		errorlog.Fatal(err)
	}
	
	var cfg config
	
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment {development|production}")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL to API")
	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	
}