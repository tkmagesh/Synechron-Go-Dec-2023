package main

import "fmt"

func main() {
	var x int8 = 100
	var f float64
	f = float64(x) //use the typename like a function for type conversion
	fmt.Println(f)
}
