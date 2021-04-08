package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Youtube Monitor!")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Youtube Subscriber Monitor")
	setupRoutes()
}
