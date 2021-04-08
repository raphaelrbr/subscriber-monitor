package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/raphaelrbr/subscriber-monitor/youtube"
)

// Set up Read and Write buffers
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Receive an incoming request and upgrade the request into
// a websocket connection
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	// Alow other origin hosts connect to our websocket server
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}

	return ws, nil
}

// websocket.go
func Writer(conn *websocket.Conn) {

	for {
		ticker := time.NewTicker((10 * time.Second))

		for t := range ticker.C {
			fmt.Printf("Update: %+v\n", t)
			items, err := youtube.GetSubscribers()
			if err != nil {
				fmt.Println(err)
			}
			//marshal our response into a JSON string
			jsonString, err := json.Marshal(items)
			if err != nil {
				fmt.Println(err)
			}
			// write this JSON string to our WebSocket
			// connection and record any errors
			err = conn.WriteMessage(websocket.TextMessage, []byte(jsonString))
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}
}
