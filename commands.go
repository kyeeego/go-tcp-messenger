package main

import (
	"fmt"
)

// IsCommand checks input if it's
func (cl *Client) IsCommand(cmd string) bool {
	if cmd[0] == '!' {
		return true
	}
	return false
}

// ExecCommand executes command
func (cl *Client) ExecCommand(cmd string, args []string) {
	// cmdName := strings.Split(cmd, " ")[0]
	// log.Println(cmd)
	switch cmd {
	case "!name":
		cl.Name(cmd, args)
	case "!exit":
		cl.Exit()
	}
}

// Name sets new name for user
func (cl *Client) Name(cmd string, args []string) {
	if len(args) == 1 {
		return
	}
	prev := cl.name
	cl.name = args[1]
	cl.outcoming <- fmt.Sprintf("%s changed his name to %s", prev, cl.name)
}

// Enter sends user to a new Hub
func (cl *Client) Enter() {

}

// Quit deletes user from the hub and pushes him to "main"
func (cl *Client) Quit() {
	msg := fmt.Sprintf("%s has left the chat", cl.name)
	cl.outcoming <- msg
}

// Exit  will delete user from his hub and close connection
func (cl *Client) Exit() {
	cl.Quit()
	cl.conn.Close()
}
