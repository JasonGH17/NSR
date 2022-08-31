package main

import (
	"log"
	"net"
	"os"
)

func checkerr(err error) bool {
	if err != nil {
		log.Println(err)
		return true
	}
	return false
}

func handleIncomingRequest(conn net.Conn, callbacks *map[string]Controller) {
	dbnbuff := make([]byte, 64)
	dbnBytes, err := conn.Read(dbnbuff)
	if !checkerr(err) {
		dbn := string(dbnbuff)[:dbnBytes]
		buffer := make([]byte, 1024)
		_, err = conn.Read(buffer)
		_ = checkerr(err)

		(*callbacks)[dbn](conn, buffer)
	}

	conn.Close()
}

func TCP(HOST string, PORT string, msgHandlers *map[string]Controller) {
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
		go handleIncomingRequest(conn, msgHandlers)
	}
}
