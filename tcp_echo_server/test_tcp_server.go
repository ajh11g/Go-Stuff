// based on https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go

package main

import (
	"fmt"
	"net"
	"os"
)

const (
	connHost = "localhost"
	connPort = "5555"
	connType = "tcp"
)

func main() {
	// listen for incoming connections
	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// close listener when app closes
	defer l.Close()
	fmt.Println("Listeninig on " + connHost + ":" + connPort)
	for {
		// listen for an incoming connection
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine
		go handleRequest(conn)
	}
}

// handles incoming requests
func handleRequest(conn net.Conn) {
	// make a bufer to hold data
	buf := make([]byte, 1024)
	// read the incoming connection into the buffer
	reqLen, err := conn.Read(buf)
	_ = reqLen
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// print the incoming message on the server
	data := buf[:reqLen]
	// Send a response and the message back to sender
	conn.Write([]byte("Message received.\n"))
	conn.Write(data)
	// close the connection when done
	conn.Close()
}
