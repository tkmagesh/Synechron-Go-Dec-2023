package calculator

import "fmt"

func init() {
	fmt.Println("calculator[multiply.go] - initialized")
}

func Multiply(x, y int) int {
	// opCount["multiply"]++
	incrementOperation("multiply")
	return x * y
}
