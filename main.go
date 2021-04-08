package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raphaelrbr/subscriber-monitor/websocket"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Youtube Monitor!")
}

// Expose YouTube stats via a websocket connection
func stats(w http.ResponseWriter, r *http.Request) {

	// upgrade HTTP connection to a webscocket one
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
		return
	}

	// Call Writter function to this websocket connection
	go websocket.Writer(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", stats)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Youtube Subscriber Monitor")
	setupRoutes()
}
