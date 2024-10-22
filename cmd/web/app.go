package main

import (
	"html/template"
	"log"
)

type application struct {
	cfg           config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
}
