package main

type commandID int

const (
	CMD_NICK commandID =iota
	CMD_JOIN 
	CMD_ROOM
	CMD_MSG
	CMD_QUIT
)

type command struct {
	id commandID
	client *client
	args []string
}