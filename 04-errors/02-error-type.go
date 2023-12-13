package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	multiplier := 100
	divisor := 0
	q, r, err := divide(multiplier, divisor)
	if err == ErrDivideByZero {
		fmt.Println("Do not attempt to divide by zero")
		return
	}
	if err != nil {
		fmt.Println("something went wrong:", err)
		return
	}
	fmt.Printf("Dividing %d by %d, quotient = %d and remainder = %d\n", multiplier, divisor, q, r)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("divide by zero error")
		return 0, 0, err
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = ErrDivideByZero
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
