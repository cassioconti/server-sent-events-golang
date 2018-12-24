package main

import (
	"log"
	"net/http"

	"github.com/cassioconti/server-sent-events-golang/server/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("client"))
	eventHandler := handlers.NewEventHandler()

	http.Handle("/", fs)
	http.HandleFunc("/v1/updates", eventHandler.Handler)

	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
