package main

import (
	"fmt"
	"net/http"
	"time"
)

type eventHandler struct {
	numberOfClients int
}

func NewEventHandler() *eventHandler {
	return &eventHandler{
		numberOfClients: 0,
	}
}

func (eventHandler *eventHandler) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	eventHandler.numberOfClients++
	fmt.Printf("New connection. There are %+v clients connected now\n", eventHandler.numberOfClients)
	go eventHandler.removeClient(w.(http.CloseNotifier).CloseNotify())

	i := 0
	for {
		i++
		fmt.Fprintf(w, "event: myEvent\nid: %d\ndata: this is an example event sent to browser\n\n", i)
		flusher.Flush()
		time.Sleep(time.Second * 10)
	}
}

func (eventHandler *eventHandler) removeClient(closed <-chan bool) {
	<-closed
	eventHandler.numberOfClients--
	fmt.Printf("Closed connection. There are %+v clients connected now\n", eventHandler.numberOfClients)
}
