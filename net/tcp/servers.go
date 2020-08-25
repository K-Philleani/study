package tcp

import (
	"fmt"
	"net"
)

func doServerStuff(conn net.Conn) {
	buf := make([]byte, 512)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading", err.Error())
		return
	}
	fmt.Printf("Received data: %v", string(buf[:]))
}

func Server() {
	fmt.Println("Starting the server")
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Print("Error listening", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}
