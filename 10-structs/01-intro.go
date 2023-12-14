package main

import "fmt"

func main() {
	/*
		var emp struct {
			Id     int
			Name   string
			Salary float32
		}
		emp.Id = 100
		emp.Name = "Magesh"
		emp.Salary = 10000
	*/

	/*
		var emp struct {
			Id     int
			Name   string
			Salary float32
		} = struct {
			Id     int
			Name   string
			Salary float32
		}{
			Id:     100,
			Name:   "Magesh",
			Salary: 10000,
		}
	*/

	// type inference
	/*
		var emp = struct {
			Id     int
			Name   string
			Salary float32
		}{
			Id:     100,
			Name:   "Magesh",
			Salary: 10000,
		}
	*/

	// using :=
	emp := struct {
		Id     int
		Name   string
		Salary float32
	}{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	fmt.Println(Format(emp))
}

func Format(emp struct {
	Id     int
	Name   string
	Salary float32
}) string {
	return fmt.Sprintf("Id = %d, Name = %q, Salary = %0.2f\n", emp.Id, emp.Name, emp.Salary)
}
