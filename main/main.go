//main.go

package main

import (
	"github.com/brentritzema/go-cs214/handler"
	"github.com/brentritzema/go-cs214/server"
)

const (
	connHost = "localhost"
	connPort = "3333"
	connType = "tcp"
)

func main() {
	server.StartListener(
		connType,
		connHost,
		connPort,
		handler.Default,
		true)
}
