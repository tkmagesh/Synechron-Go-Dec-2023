package main

import "fmt"

// unused varible are allowed at package scope
// var app_name string
// var app_name string = "learning-app"
// var app_name = "learning-app"

// the following is not allowed
// app_name := "learning-app"

func main() {
	/*
		var name string
		name = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	// declaration & initialization
	/*
		var name string = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	// type inference
	/*
		var name = "Magesh"
		fmt.Printf("Hi %s, Have a nice day!\n", name)
	*/

	// Use :=
	name := "Magesh"
	fmt.Printf("Hi %s, Have a nice day!\n", name)

	// unused variables are NOT allowed in function scope
	/*
		var x int
		x += 100
	*/
	// fmt.Println(x)

	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y int = 100, 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Sum of %d and %d is %d\n"
			result int    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y   = 100, 200
			str    = "Sum of %d and %d is %d\n"
			result = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of %d and %d is %d\n"
			result    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	x, y, str := 100, 200, "Sum of %d and %d is %d\n"
	result := x + y
	fmt.Printf(str, x, y, result)
}
