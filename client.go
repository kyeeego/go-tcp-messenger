package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// Client ...
type Client struct {
	name      string
	conn      net.Conn
	reader    bufio.Reader
	hub       *Hub
	outcoming chan string
}

// NewClient ...
func NewClient(conn net.Conn, hub *Hub) *Client {

	cl := &Client{
		name:      "annonymous",
		conn:      conn,
		reader:    *bufio.NewReader(conn),
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
		msg, err := cl.reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			cl.conn.Close()
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		isCmd := cl.IsCommand(cmd)
		if !isCmd {
			cl.outcoming <- msg
			continue
		}
		cl.ExecCommand(cmd, args)
	}
}

func (cl *Client) Write() {
	for msg := range cl.outcoming {
		msg = fmt.Sprintf("%s: %s\n\n", cl.name, msg)
		cl.hub.broadcast(msg)
	}
}
