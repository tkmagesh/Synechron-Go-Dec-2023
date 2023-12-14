package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func main() {

	/*
		var emp Employee
		emp.Id = 100
		emp.Name = "Magesh"
		emp.Salary = 10000

		fmt.Println(emp)
	*/

	/*
		var emp Employee = Employee{
			Id:     100,
			Name:   "Magesh",
			Salary: 10000,
		}
		fmt.Println(emp)
	*/

	// type inference

	/*
		var emp = Employee{
			Id:     100,
			Name:   "Magesh",
			Salary: 10000,
		}
		fmt.Println(emp)
	*/

	// using :=

	emp := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}

	// emp := Employee{100, "Magesh", 10000} // NOT advisable
	fmt.Println(Format(emp))

}

func Format(emp Employee) string {
	return fmt.Sprintf("Id = %d, Name = %q, Salary = %0.2f\n", emp.Id, emp.Name, emp.Salary)
}
