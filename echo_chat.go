package main

import (
	"io"
	"log"
	"net"
	"fmt"
)

const listenAddr = "localhost:4000"

var partner = make(chan io.ReadWriteCloser)

func match(c io.ReadWriteCloser) {
	fmt.Fprintln(c, "Waiting for a partner ...")

	select {
	case partner <- c:
		// handled by the other goroutine
	case p := <-partner:
		chat(p, c)
	}
}

func chat(a, b io.ReadWriteCloser) {
	fmt.Fprintln(a, "Found one! Say hi.")
	fmt.Fprintln(b, "Found one! Say hi.")

	errc := make(chan error, 1)

	go cp(a, b, errc)
	go cp(b, a, errc)

	if err := <-errc; err != nil {
		log.Println(err)
	}

	a.Close()
	b.Close()

	go io.Copy(b, a)
	io.Copy(a, b)
}

func cp(w io.Writer, r io.Reader, errc chan<- error) {
	_, err := io.Copy(w, r)
	errc <- err
}

func main() {
	listener, err := net.Listen("tcp", listenAddr)

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go match(conn)
	}
}