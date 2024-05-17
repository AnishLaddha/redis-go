package main

import "strconv"

func serializeBulkStream(out_str string) string {
	str_len := len(out_str)
	return "$" + strconv.Itoa(str_len) + "\r\n" + out_str + "\r\n"
}

func serializeSimpleString(out_str string) string {
	return "+" + out_str + "\r\n"
}
