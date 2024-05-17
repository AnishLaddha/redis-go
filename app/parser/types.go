package parser

type RESPType interface{}

type SimpleString struct {
	Value string
}

type Error struct {
	Value string
}

type Integer struct {
	Value int64
}

type BulkString struct {
	Value string
	Nil   bool
}

type Array struct {
	Values []RESPType
}
