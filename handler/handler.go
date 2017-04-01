package handler

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
)

// Default returns the size of the message and repeats
// the message back to the sender
func Default(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Builds the message.
	message := "Hi, I received your message! It was "
	message += strconv.Itoa(reqLen)
	message += " bytes long and here's what it said: \""
	n := bytes.Index(buf, []byte{0})
	message += string(buf[:n-1])
	message += "\"!\n"

	// Write the message in the connection channel.
	conn.Write([]byte(message))
	// Close the connection when you're done with it.
	conn.Close()
}
