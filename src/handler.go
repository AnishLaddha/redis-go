package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

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

		output := ""

		switch casted_result := result.(type) {
		case parser.Array:
			switch first_element := casted_result.Values[0].(type) {
			case parser.BulkString:
				switch command_str := strings.ToLower(first_element.Value); command_str {
				case "echo":
					output = serializeBulkStream("")
					if len(casted_result.Values) >= 2 {
						second_element := casted_result.Values[1].(parser.BulkString)
						output = serializeBulkStream(second_element.Value)
					}
				case "ping":
					output = serializeSimpleString("PONG")
				case "set":
					output = handleSet(db, casted_result)
				case "get":
					output = handleGet(db, casted_result)
				}
			}
		}

		_, err = c.Write([]byte(output))
		if err != nil {
			fmt.Println("Error writing connection: ", err.Error())
			os.Exit(1)
		}
	}
}

func handleSet(db *Database, result parser.Array) string {
	result_length := len(result.Values)
	key := result.Values[1].(parser.BulkString).Value
	expiry_time := time.Time{}
	val := result.Values[2].(parser.BulkString).Value
	if result_length == 5 {
		precision := result.Values[3].(parser.BulkString).Value
		expiry_str := result.Values[4].(parser.BulkString).Value
		expiry_ms, err := strconv.Atoi(expiry_str)
		if err != nil {
			fmt.Println("Error converting the expiry bulk string to int (handleSet): ", err.Error())
			os.Exit(1)
		}
		if strings.ToLower(precision) == "ex" {
			expiry_ms = expiry_ms * 1000
		}
		expiry_time = time.Now().Add(time.Millisecond * time.Duration(expiry_ms))
	}
	db.Set(key, val, expiry_time)
	return serializeSimpleString("OK")
}

func handleGet(db *Database, result parser.Array) string {
	key := result.Values[1].(parser.BulkString).Value
	value, exists := db.Get(key)
	output := serializeBulkStream("")
	if exists {
		output = serializeBulkStream(value.val)
	}
	return output
}
