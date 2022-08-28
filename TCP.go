package main

import (
	"log"
	"net"
	"os"
)

func handleIncomingRequest(conn net.Conn, callback func(net.Conn, []byte)) {
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	callback(conn, buffer)

	conn.Close()
}

func TCP(HOST string, PORT string, msgHandler func(net.Conn, []byte)) {
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
		go handleIncomingRequest(conn, msgHandler)
	}
}
