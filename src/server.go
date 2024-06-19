package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Logs:")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()
	db := newDatabase()
	aofWriter := NewAOFWriter("../persist.aof")
	defer aofWriter.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handle_conn(conn, db, aofWriter)
	}

}
