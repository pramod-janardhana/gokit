package main

import "fmt"

func NewSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	next := NewSeq()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}
