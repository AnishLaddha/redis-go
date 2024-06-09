package parser

import (
	"bufio"
	"strings"
	"testing"
)

func TestParseRESP(t *testing.T) {
	tests := []struct {
		input    string
		expected RESPType
	}{
		{"+OK\r\n", SimpleString{Value: "OK"}},
		{"-Error message\r\n", Error{Value: "Error message"}},
		{":123\r\n", Integer{Value: 123}},
		{"$6\r\nfoobar\r\n", BulkString{Value: "foobar"}},
		{"$-1\r\n", BulkString{Nil: true}},
		{"*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n", Array{Values: []RESPType{
			BulkString{Value: "foo"},
			BulkString{Value: "bar"},
		}}},
		{"*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n", Array{Values: []RESPType{
			BulkString{Value: "ECHO"},
			BulkString{Value: "hey"}}}},
	}

	for _, test := range tests {
		reader := bufio.NewReader(strings.NewReader(test.input))
		result, err := ParseRESP(reader)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !compareRESP(result, test.expected) {
			t.Errorf("expected %v, got %v", test.expected, result)
		}
	}
}

func compareRESP(a, b RESPType) bool {
	switch a := a.(type) {
	case SimpleString:
		b, ok := b.(SimpleString)
		return ok && a.Value == b.Value
	case Error:
		b, ok := b.(Error)
		return ok && a.Value == b.Value
	case Integer:
		b, ok := b.(Integer)
		return ok && a.Value == b.Value
	case BulkString:
		b, ok := b.(BulkString)
		return ok && a.Value == b.Value && a.Nil == b.Nil
	case Array:
		b, ok := b.(Array)
		if !ok || len(a.Values) != len(b.Values) {
			return false
		}
		for i := range a.Values {
			if !compareRESP(a.Values[i], b.Values[i]) {
				return false
			}
		}
		return true
	default:
		return false
	}
}
