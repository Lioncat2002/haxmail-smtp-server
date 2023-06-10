package main

import (
	"fmt"
	"net"
	"smtpserver/src/services"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:2525")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("SMTP server listening on port 2525")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed of accept connection", err)
			continue
		}
		go services.HandleClient(conn)
	}
}
