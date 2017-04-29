package handler

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
)

type user struct {
	name  string
	age   int
	phone string
}

func (u *user) returnFormattedData() string {
	message := "Their name is "
	message += u.name
	message += ", they are: "
	message += strconv.Itoa(u.age)
	message += " years old, and can be reached at: "
	message += u.phone + "\n"
	return message
}

//fake user data
var users = map[string]user{
	"bjr": {"Brent", 20, "555-555-5555"},
	"klv": {"Kari", 20, "555-555-5554"},
	"bdr": {"Brad", 18, "555-555-5553"},
	"kjr": {"Katie", 15, "555-555-5552"},
	"lar": {"Leslie", 45, "555-555-5551"},
	"jbr": {"John", 53, "555-555-5550"},
	"lav": {"Lori", 55, "555-555-5549"},
	"jav": {"John", 55, "555-555-5548"},
}

//ProcessConnection takes the connection, reads the data (a username)
//and writes back formatted user data
func ProcessConnection(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil && reqLen != 0 {
		fmt.Println("Error reading:", err.Error())
	}
	// Builds the message.
	n := bytes.Index(buf, []byte{0})
	username := string(buf[:n-1])
	//TODO check for error here
	u := users[username]
	message := u.returnFormattedData()

	// Write the message in the connection channel.
	conn.Write([]byte(message))
	// Close the connection when you're done with it.
	conn.Close()

}
