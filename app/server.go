package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

type Database struct {
	data map[string]string
	mux  sync.Mutex
}

func main() {
	fmt.Println("Logs:")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()
	db := Database{
		data: make(map[string]string),
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handle_conn(conn, &db)
	}

}
