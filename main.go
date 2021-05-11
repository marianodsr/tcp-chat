package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	s := newServer()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error starting server")
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection, %s", err.Error())
			continue
		}

		go s.newClient(conn)

	}
}
