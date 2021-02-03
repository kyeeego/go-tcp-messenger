package main

import (
	"log"
	"net"
)

// Server ...
type Server struct {
	hubs []*Hub
}

// NewServer creates a new Server instance
func NewServer() *Server {
	var hubs []*Hub
	return &Server{
		hubs: hubs,
	}
}

// Run runs a server
func (srv *Server) Run() {
	ln, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}

	var cls []*Client
	mainHub := &Hub{
		name:    "main",
		clients: cls,
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		_ = NewClient(conn, mainHub)
	}
}
