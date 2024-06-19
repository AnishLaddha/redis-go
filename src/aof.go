package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/AnishLaddha/redis-go/src/parser"
)

type AOFWriter struct {
	logChannel chan parser.Array
	file       *os.File
	mux        sync.Mutex
}

func NewAOFWriter(filename string) *AOFWriter {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err) // or handle more gracefully
	}
	writer := &AOFWriter{
		logChannel: make(chan parser.Array, 1024), // Adjust the buffer size based on expected load
		file:       file,
	}
	go writer.processWrites()
	return writer
}

func (a *AOFWriter) processWrites() {
	for cmd := range a.logChannel {
		strings := parser.ParseArrayToStr(cmd)
		a.mux.Lock()
		for _, str := range strings {
			_, err := a.file.WriteString(str + " ")
			if err != nil {
				fmt.Println("Failed logging command to AOF")
			}
		}
		_, err := a.file.WriteString("\n")
		if err != nil {
			fmt.Println("Failed logging command to AOF")
		}
		a.mux.Unlock()
	}
}

func (a *AOFWriter) LogCommand(cmd parser.Array) {
	a.logChannel <- cmd
}

func (a *AOFWriter) Close() {
	close(a.logChannel)
	a.file.Close()
}
