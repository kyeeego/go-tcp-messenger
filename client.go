package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// Client ...
type Client struct {
	conn      net.Conn
	reader    bufio.Reader
	hub       *Hub
	outcoming chan string
}

// NewClient ...
func NewClient(conn net.Conn, hub *Hub) *Client {

	cl := &Client{
		conn:      conn,
		reader:    *bufio.NewReader(os.Stdin),
		hub:       hub,
		outcoming: make(chan string),
	}

	hub.clients = append(hub.clients, cl)

	fmt.Printf("Connected to %s", hub.name)

	cl.Work()

	return cl
}

// Work ..
func (cl *Client) Work() {
	go cl.Read()
	go cl.Write()
}

func (cl *Client) Read() {
	for {
		tmp := make([]byte, 1024)
		_, err := cl.conn.Read(tmp)
		if err != nil {
			log.Fatal(err)
		}

		cl.outcoming <- string(tmp)
	}
}

func (cl *Client) Write() {
	for msg := range cl.outcoming {
		msg = fmt.Sprintf("Message: %sFrom: %s\n\n", msg, cl.conn.RemoteAddr().String())

		cl.hub.broadcast(msg)
	}
}
