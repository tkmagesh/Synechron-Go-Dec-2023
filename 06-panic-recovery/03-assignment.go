/*
modify the below program in such a way that
  - it accepts the multiplier & divisor from the user and print the quotient and remainder
  - if a panic occurrs, inform the user and accept the next set of multiplier & divisor and calculate and print the quotient and remainder
    else exit the application
*/
package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	multiplier := 100
	divisor := 0
	q, r, err := divideClient(multiplier, divisor)
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

// wrapper to convert the panic into an error
func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

// 3rd party library
func divide(x, y int) (quotient, remainder int) {
	if y == 0 {
		panic(ErrDivideByZero)
	}
	quotient = x / y
	remainder = x % y
	return
}
