package server

import (
	"fmt"
	"log"
	"net"
)

func RunTCPServer(s Server) {
	port := s.Port()

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("0.0.0.0:%v", port))
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening for TCP Connections on port %v", port)

	ctx := s.Setup()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go s.Handle(ctx, conn)
	}
}
