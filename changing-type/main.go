// Add methods to primitive types
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num1 int = 100
	num2 := float64(num1)
	fmt.Printf("%v is type casted to %v\n", reflect.TypeOf(num1), reflect.TypeOf(num2))

	var num3 interface{} = 100
	num4 := num3.(int)
	fmt.Printf("%v is type asserted to %v\n", reflect.TypeOf(num3), reflect.TypeOf(num4))
}
