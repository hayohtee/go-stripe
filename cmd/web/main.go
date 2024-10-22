package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
)

const (
	version = "1.0.0"
	cssVersion = "1"
)

func main()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	
	var cfg config
	
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment {development|production}")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL to API")
	flag.Parse()


}