package main

import (
	"net"
)

func main() {
	println("Starting TCP server...")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		println("Error starting TCP server:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)
		length, err := conn.Read(buffer)
		if err != nil {
			println("Error reading:", err)
			return
		}
		message := string(buffer[:length])
		println("Received message:", message)
	}
}
