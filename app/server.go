package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

type data struct {
	val    string
	expiry time.Time
}

type Database struct {
	keyval map[string]data
	mux    sync.Mutex
}

func newDatabase() *Database {
	return &Database{
		keyval: make(map[string]data),
	}
}

func (db *Database) Set(key string, item string, exp time.Time) {
	db.mux.Lock()
	defer db.mux.Unlock()
	db.keyval[key] = data{val: item, expiry: exp}
}

func (db *Database) Get(key string) (data, bool) {
	db.mux.Lock()
	defer db.mux.Unlock()
	d, exists := db.keyval[key]
	if exists && !d.expiry.IsZero() && time.Now().After(d.expiry) {
		delete(db.keyval, key)
		exists = false
	}
	return d, exists
}

func main() {
	fmt.Println("Logs:")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()
	db := newDatabase()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handle_conn(conn, db)
	}

}
