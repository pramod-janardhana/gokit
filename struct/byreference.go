// Person struct with methods of pointer receiver

package main

import (
	"fmt"
)

type Person struct {
	firstName, lastName string
	location            string
}

// A person method with pointer receiver
func (p *Person) Name() string {
	return p.firstName + p.lastName
}

// A person method with pointer receiver
func (p *Person) Location() string {
	return p.location
}

// A person method with pointer receiver
func (p *Person) SetLocation(newLocation string) {
	p.location = newLocation
}
func main() {
	p := &Person{
		"Pramod",
		"J",
		"Bangalore",
	}
	p.SetLocation("Mysore")
	fmt.Println(p.Name())
	fmt.Println(p.Location())
}
