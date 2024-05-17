package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/app/parser"
)

func handle_conn(c net.Conn, db *Database) {
	buf := make([]byte, 128)

	for {
		_, err := c.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("Error reading connection: ", err.Error())
			os.Exit(1)
		}
		reader := bufio.NewReader(strings.NewReader(string(buf)))
		result, err := parser.ParseRESP(reader)
		if err != nil {
			fmt.Println("Error parsing RESP: ", err.Error())
			os.Exit(1)
		}
		db.mux.Lock()

		output := ""

		switch casted_result := result.(type) {
		case parser.Array:
			switch first_element := casted_result.Values[0].(type) {
			case parser.BulkString:
				command_str := strings.ToLower(first_element.Value)
				if command_str == "echo" {
					second_element := casted_result.Values[1].(parser.BulkString)
					output = serializeBulkStream(second_element.Value)
				} else if command_str == "ping" {
					output = serializeSimpleString("PONG")
				} else if command_str == "set" {
					key := casted_result.Values[1].(parser.BulkString).Value
					val := casted_result.Values[2].(parser.BulkString).Value
					db.data[key] = val
					output = serializeSimpleString("OK")
				} else if command_str == "get" {
					key := casted_result.Values[1].(parser.BulkString).Value
					value, exists := db.data[key]
					if !exists {
						output = serializeBulkStream("")
					} else {
						output = serializeBulkStream(value)
					}
				}
			}
		}
		db.mux.Unlock()

		_, err = c.Write([]byte(output))
		if err != nil {
			fmt.Println("Error writing connection: ", err.Error())
			os.Exit(1)
		}
	}
}
