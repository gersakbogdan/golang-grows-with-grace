package main

import (
	"fmt"
	"net/http"
	"log"
)

const listenAddr = "localhost:4000"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Web!")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, YOU!")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/name", handler2)

	err := http.ListenAndServe(listenAddr, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server is listening on port 4000")
}