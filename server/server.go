package server

import (
	"fmt"
	"net"
	"time"
)

//RequestHandler is a type of function with handles a request given a connection
type RequestHandler func(conn net.Conn)

// StartListener takes a connection type (ex: tcp), a host ip (localhost),
//a port, and a handler to handle the connection
func StartListener(
	connHost,
	connPort,
	connType string,
	handler RequestHandler,
	concurrency bool,
	stop chan bool) {

	//-----------Setup listener-----------

	// Listen for incoming connections. Set l to the listener and err to the error
	lAddr, tcpErr := net.ResolveTCPAddr(connType, connHost+":"+connPort)
	if tcpErr != nil {
		fmt.Println("Error processing " + connHost + ":" + connPort)
	}

	l, err := net.ListenTCP(connType, lAddr)
	if err != nil {
		//print the error and exit if there is one
		fmt.Println("Error listening:", err.Error())
	}

	// Give useful information on listening
	fmt.Println("Listening on " + connHost + ":" + connPort)
	//--------------------------

	defer l.Close()

	// Start the unending loop that doesnt exit until the program does
	for {

		//set a deadline so that l.Accept can't have an infinite lock if no connection is recieved
		l.SetDeadline(time.Now().Add(time.Second))

		// Listen for an incoming connection and accept it.
		conn, err := l.Accept()

		// if a timeout error and quit = true
		if err, ok := err.(*net.OpError); ok && err.Timeout() {
			//use select, not if, or else it will block until something comes through the channel
			select {
			case <-stop:
				fmt.Println("Stopping Listerner")
				//quit if something is sent throught the stop channel
				return
			default:
				//do nothing, just needed for no blocking
			}
		} else if err != nil {
			//handle the error if it's real
			fmt.Println("Error accepting: ", err.Error())
		} else {
			//only get here if there was a connection accepted and no errors

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
}
