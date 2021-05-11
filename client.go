package main

import (
	"fmt"
	"net"
)

type client struct {
	conn net.Conn
	room *room
	nick string
}

func (c *client) readMessages() {
	for {
		message := make([]byte, 4096)
		n, err := c.conn.Read(message)
		if err != nil {
			c.conn.Close()
			break
		}
		if n > 0 {
			fmt.Printf("\n%s -> %s", c.nick, string(message))
			c.send(fmt.Sprintf("ECHOING MESSAGE: %s", string(message)))
		}
	}

}

func (c *client) send(msg string) {
	c.conn.Write([]byte(msg))
}
