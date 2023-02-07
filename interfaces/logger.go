package main

import "fmt"

type FileWriter struct{}

// FileWriter writes logs into file
func (fw FileWriter) Write(message string) {
	fmt.Printf("Writing '%s' to file\n", message)
}

type Logger struct {
	fw FileWriter
}

func NewLogger(fw FileWriter) Logger {
	return Logger{fw: fw}
}

func (l Logger) Log(message string) {
	l.fw.Write(message)
}
