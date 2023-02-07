package main

import (
	"fmt"
)

type Number interface {
	int | int32 | int64 | float64
}

// SumNumbers sums the values of slice m. Its supports both integers
// and floats values.
func SumNumbers[V Number](m []V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// SumIntsOrFloats sums the values of slice m. Its supports both
// int64 and float64 values.
func SumIntsOrFloats[V int64 | float64](m []V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

// SumInts adds together the values of m.
func SumInts(m []int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m []float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	// Initialize a slice for the integer values
	ints := []int64{34, 12}

	// Initialize a slice for the float values
	floats := []float64{35.98, 26.99}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[int64](ints),
		SumIntsOrFloats[float64](floats))

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))

}
