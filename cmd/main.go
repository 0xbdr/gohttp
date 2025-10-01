package main

import (
	"fmt"
	"gohttp/internal/parser"
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

	var server, err = net.Listen("tcp", ":9099")
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
	var trequest, err = parser.Parserequest(client)
	fmt.Println(trequest)
	check(err)
	fmt.Println(trequest)

}
