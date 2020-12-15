package main

import (
	"github.com/siddontang/go-mysql/server"
	"net"
)

func main() {
	l, _ := net.Listen("tcp", "127.0.0.1:4000")

	c, _ := l.Accept()

	// Create a connection with user root and an empty password.
	// You can use your own handler to handle command here.
	ttt := server.EmptyHandler{}
	conn, _ := server.NewConn(c, "root", "", ttt)
	for {

		conn.HandleCommand()
	}

}
