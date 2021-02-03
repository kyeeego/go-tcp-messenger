package main

// Hub ...
type Hub struct {
	name    string
	clients []*Client
}

func (hub *Hub) broadcast(msg string) {
	for _, cl := range hub.clients {
		cl.conn.Write([]byte(msg))
	}
}
