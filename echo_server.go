package main

import (
	"io"
	"log"
	"net"
)

const listenAddr = "localhost:4000"

func main() {
	listener, err := net.Listen("tcp", listenAddr)

	if err != nil {
		log.Fatal(err)
	}

	for {
		connection, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go io.Copy(connection, connection)
	}
}