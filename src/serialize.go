package main

import (
	"fmt"
	"strconv"
)

func serializeBulkStream(out_str string) string {
	str_len := len(out_str)
	if str_len == 0 {
		return "$-1\r\n"
	}
	return "$" + strconv.Itoa(str_len) + "\r\n" + out_str + "\r\n"
}

func serializeSimpleString(out_str string) string {
	return "+" + out_str + "\r\n"
}

func serializeInteger(out_int int) string {
	return fmt.Sprintf(":%d\r\n", out_int)
}
