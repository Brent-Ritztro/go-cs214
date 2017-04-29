package server

import (
	"fmt"
	"net"
)

//RequestHandler is a type of function with handles a request given a connection
type RequestHandler func(conn net.Conn)

// StartListener takes a connection type (ex: tcp), a host ip (localhost),
//a port, and a handler to handle the connection
func StartListener(
	connType,
	connHost,
	connPort string,
	handler RequestHandler,
	concurrency bool) {

	// Listen for incoming connections. Set l to the listener and err to the error
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		//print the error and exit if there is one
		fmt.Println("Error listening:", err.Error())
	}

	// Close the listener when the listender stops.
	defer l.Close()

	// Give useful information on listening
	fmt.Println("Listening on " + connHost + ":" + connPort)

	// Start the unending loop that doesnt exit until the program does
	for {
		// Listen for an incoming connection and accept it.
		conn, err := l.Accept()

		//handle the error
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}

		//logs an incoming message to the console
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		// Handle connections in a new goroutine (concurrently).
		if concurrency == true {
			go handler(conn)
		} else {
			handler(conn)
		}
	}
}
