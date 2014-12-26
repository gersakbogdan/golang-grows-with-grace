package main

import (
	"fmt"
	"log"
	"net/http"
	"code.google.com/p/go.net/websocket"
)

const listenAddr = "localhost:4000"

func handler(conn *websocket.Conn) {
	var s string

	fmt.Fscan(conn, &s)
	fmt.Println("Receive: ", s)
	fmt.Fprint(conn, "How are you?")
}

func main() {
	http.Handle("/", websocket.Handler(handler))
	
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}