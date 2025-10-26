package main

import (
	"fmt"
	"hyperttps/internal/httpenums"
	"hyperttps/internal/parser"
	"net"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func main() {

	fmt.Println("starting")

	var server, err = net.Listen("tcp", ":8080")
	check(err)
	fmt.Println("Server started")
	for {
		var connection, err = server.Accept()
		check(err)
		fmt.Printf("connection from %v\n", connection.RemoteAddr().String())

		go handlec(connection)

	}
}
func handlec(client net.Conn) {
	var rrequest, err = parser.Parserequest(client)
	fmt.Println(rrequest)
	check(err)
	fmt.Println(rrequest)
	if rrequest.Method == string(httpenums.GET){
		
	}
}
