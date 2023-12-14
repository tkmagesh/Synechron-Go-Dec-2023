package main

import "fmt"

type Organization struct {
	Name string
	City string
}

type Employee struct {
	Id     int
	Name   string
	Salary float32
	Org    *Organization
}

func main() {
	synechron := Organization{
		Name: "Synechron",
		City: "Bengaluru",
	}

	e1 := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
		Org:    &synechron,
	}

	e2 := Employee{
		Id:     200,
		Name:   "Suresh",
		Salary: 20000,
		Org:    &synechron,
	}

	// changing Org city to "Pune"
	synechron.City = "Pune"

	fmt.Println("e1.Org.City : ", e1.Org.City)
	fmt.Println("e2.Org.City : ", e2.Org.City)
}
