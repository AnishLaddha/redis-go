package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/codecrafters-io/redis-starter-go/app/parser"
)

func serializeBulkStream(out_str string) string {
	str_len := len(out_str)
	return "$" + strconv.Itoa(str_len) + "\r\n" + out_str + "\r\n"
}

func handle_conn(c net.Conn) {
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

		output_str := ""
		output := ""

		switch casted_result := result.(type) {
		case parser.Array:
			switch first_element := casted_result.Values[0].(type) {
			case parser.BulkString:
				if strings.ToLower(first_element.Value) == "echo" {
					second_element := casted_result.Values[1].(parser.BulkString)
					output_str = second_element.Value
					output = serializeBulkStream(output_str)
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

func main() {
	fmt.Println("Logs:")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handle_conn(conn)
	}

}
