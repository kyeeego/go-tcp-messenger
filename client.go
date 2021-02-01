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
	server    *Server
	outcoming chan string
}

// NewClient ...
func NewClient(conn net.Conn, srv *Server) *Client {

	cl := &Client{
		conn:      conn,
		reader:    *bufio.NewReader(os.Stdin),
		server:    srv,
		outcoming: make(chan string),
	}

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
		tmp := make([]byte, 256)
		_, err := cl.conn.Read(tmp)
		if err != nil {
			log.Fatal(err)
		}

		cl.outcoming <- string(tmp)
	}
}

func (cl *Client) Write() {
	for msg := range cl.outcoming {
		msg = fmt.Sprintf("Message: %s\nFrom: %s\n", msg, cl.conn.RemoteAddr().String())
		for _, c := range cl.server.clients {
			c.conn.Write([]byte(msg))
		}
	}
}
