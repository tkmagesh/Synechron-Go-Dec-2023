package main

import (
	"errors"
	"fmt"
)

func main() {
	multiplier := 100
	divisor := 7
	q, r, err := divide(multiplier, divisor)
	if err != nil {
		fmt.Println(err)
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
		err = errors.New("divide by zero error")
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
