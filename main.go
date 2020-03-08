package main

import (
	"io"
	"net"

	"log"
)

func echo(conn net.Conn) {
	defer conn.Close()

	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalf("Unable to read/write data: %s", err.Error())
	}
}

func main() {
	port := ":20080"
	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Unable to bind to port %s\n", port)
	}
	log.Println("Listening on 0.0.0.0" + port)

	for {
		conn, err := listener.Accept()
		log.Println("Received connection")

		if err != nil {
			log.Fatalf("Unable to accept connection: %s\n", err.Error())
		}
		go echo(conn)
	}
}
