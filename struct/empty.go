package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s struct{} // empty struct consumes zero bytes,
	fmt.Println(unsafe.Sizeof(s))

	var arr [1000]struct{}
	fmt.Println(unsafe.Sizeof(arr))
}

// N returns a slice of n 0-sized elements
func N(n int) []struct{} {
	return make([]struct{}, n)
}
