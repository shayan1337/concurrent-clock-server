// TCP server that periodically writes the time

package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handleConnection(connection)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().String())
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}


