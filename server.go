package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader = bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message", err)
			break
		}
		var grid [][]int
		err = json.Unmarshal([]byte(message), &grid)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			break
		}

	}
}

func main() {
	ln, err = net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error connecting to server")
		os.Exit(1)
	}
	defer ln.Close()
	// accept connections
	for {
		conn, err = ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			continue
		}
		go handleConnection(conn)
	}
}
