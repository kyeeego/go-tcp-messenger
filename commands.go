package main

import (
	"log"
	"strings"
)

// IsCommand checks input if it's
func (cl *Client) IsCommand(cmd string) bool {
	if cmd[0] == '!' {
		return true
	}
	return false
}

// ExecCommand executes command
func (cl *Client) ExecCommand(cmd string) {
	cmd = strings.Split(cmd, " ")[0]
	log.Println(cmd)
	switch cmd {
	case "!test":
		log.Println(strings.Split(cmd, " ")[0])

		comm := testHandler(cmd)
		log.Println(comm)
		cl.outcoming <- comm
	}
}

func testHandler(msg string) string {
	return "Got !test command"
}
