package main

import (
	"net"
	"os"
	"log"
)

func handleIncomingRequest(conn net.Conn) {
    buffer := make([]byte, 1024)
    _, err := conn.Read(buffer)
    if err != nil {
        log.Fatal(err)
    }
    conn.Write([]byte("TCP Server\n"))

    conn.Close()
}

func TCP(HOST string, PORT string) {
	listen, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	
	defer listen.Close()

	for {
        conn, err := listen.Accept()
        if err != nil {
            log.Fatal(err)
            os.Exit(1)
        }
        go handleIncomingRequest(conn)
    }
}