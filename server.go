package main

import (
	"log"
	"net"
)

// Server ...
type Server struct {
	clients []*Client
}

// NewServer creates a new Server instance
func NewServer() *Server {
	var cls []*Client
	return &Server{
		clients: cls,
	}
}

// Run runs a server
func (srv *Server) Run() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		cl := NewClient(conn, srv)

		// Adding new connection to the "Hub"
		srv.clients = append(srv.clients, cl)
	}
}
