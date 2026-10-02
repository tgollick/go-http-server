package main

import (
	"errors"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("connection opened on port :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}

		log.Printf("connection accepted from %s", conn.RemoteAddr())
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)

		bytesRead, err := conn.Read(buffer)

		// EOF means a client hung up cleanly, which is a normal end, not an error
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			log.Print(err)
			break
		}

		log.Printf("read %d bytes from connection %s", bytesRead, conn.RemoteAddr())

		_, err = conn.Write(buffer[:bytesRead])
		if err != nil {
			log.Print(err)
			break
		}

		log.Printf("written %d bytes to connection %s", bytesRead, conn.RemoteAddr())
	}

	log.Printf("closing connection from address: %s", conn.RemoteAddr())
}
