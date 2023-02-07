package main

import "fmt"

type Employee struct {
	firstName, lastName string
	location            string
}

func (e Employee) Name() string {
	return e.firstName + e.lastName
}

func (e Employee) PrintDetails() {
	fmt.Printf("%+v", e)
}

type Developer struct {
	Employee //type embedding for composition
	Skills   []string
}

func main() {
	jim := Developer{
		Employee: Employee{
			firstName: "Jim",
			lastName:  "Co",
			location:  "United States",
		},
		Skills: []string{"Go", "Docker", "Kubernetes"},
	}
	jim.Name()
	jim.PrintDetails()
}
