package parser

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func ParseRESP(reader *bufio.Reader) (RESPType, error) {
	prefix, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}
	switch prefix {
	case '+':
		return parseSimpleString(reader)
	case '-':
		return parseError(reader)
	case ':':
		return parseInteger(reader)
	case '$':
		return parseBulkString(reader)
	case '*':
		return parseArray(reader)
	default:
		return nil, errors.New("unknown prefix")
	}

}

func parseSimpleString(reader *bufio.Reader) (SimpleString, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error parsing simple string: ", err.Error())
		return SimpleString{}, err
	}
	str := strings.TrimSuffix(line, "\r\n")
	return SimpleString{Value: str}, nil

}

func parseError(reader *bufio.Reader) (Error, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error parsing error: ", err.Error())
		return Error{}, err
	}
	str := strings.TrimSuffix(line, "\r\n")

	return Error{Value: str}, nil
}

func parseInteger(reader *bufio.Reader) (Integer, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error parsing Int: ", err.Error())
		return Integer{}, err
	}
	str := strings.TrimSuffix(line, "\r\n")
	val, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting to int: ", err.Error())
		return Integer{}, err
	}
	return Integer{Value: int64(val)}, nil

}

func parseBulkString(reader *bufio.Reader) (BulkString, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error parsing BulkString(line 52): ", err.Error())
		return BulkString{}, err
	}
	length, err := strconv.Atoi(strings.TrimSuffix(line, "\r\n"))
	if err != nil {
		fmt.Println("Error parsing BulkString(line 56): ", err.Error())
		return BulkString{}, err
	}
	if length == -1 {
		return BulkString{Nil: true}, nil
	}
	data := make([]byte, length+2)
	_, err = io.ReadFull(reader, data)
	if err != nil {
		fmt.Println("Error parsing BulkString(line 63): ", err.Error())
		return BulkString{}, err
	}
	return BulkString{Value: string(data[:length])}, nil
}

func parseArray(reader *bufio.Reader) (Array, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return Array{}, err
	}
	length, err := strconv.Atoi(strings.TrimSuffix(line, "\r\n"))
	if err != nil {
		return Array{}, err
	}
	if length == -1 {
		return Array{Values: nil}, nil
	}
	values := make([]RESPType, length)

	for i := 0; i < length; i++ {
		value, err := ParseRESP(reader)
		if err != nil {
			return Array{}, err
		}
		values[i] = value
	}
	return Array{Values: values}, nil
}
