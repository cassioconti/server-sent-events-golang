package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("client"))
	eventHandler := NewEventHandler()

	http.Handle("/", fs)
	http.HandleFunc("/v1/updates", eventHandler.Handler)

	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
