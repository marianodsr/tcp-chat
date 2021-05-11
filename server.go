package main

import "net"

type server struct {
	rooms map[string]*room
}

func newServer() *server {
	return &server{
		rooms: make(map[string]*room),
	}
}

func (s *server) newClient(conn net.Conn) {
	client := &client{
		conn: conn,
		nick: "guest",
	}

	client.readMessages()
}
