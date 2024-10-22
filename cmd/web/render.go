package main

type templateData struct {
	stringMap map[string]string
	intMap    map[string]int
	floatMap  map[string]float32
	data      map[string]any
	csrfToken string
}