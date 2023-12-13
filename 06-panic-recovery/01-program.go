package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero error = errors.New("divide by zero error")

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[deferred-main] Recovered from panic - 1, err :", err)
		}
		if err := recover(); err != nil {
			fmt.Println("[deferred-main] Recovered from panic - 2, err :", err)
		}
		fmt.Println("Thank you!")
	}()
	multiplier := 100
	divisor := 0
	q, r := divide(multiplier, divisor)
	fmt.Printf("[main] - Dividing %d by %d, quotient = %d and remainder = %d\n", multiplier, divisor, q, r)
}

func divide(x, y int) (quotient int, remainder int) {
	if y == 0 {
		panic(ErrDivideByZero)
	}
	fmt.Println("[divide] - calculating quotient")
	quotient = x / y
	fmt.Println("[divide] - calculating remainder")
	remainder = x % y
	return
}
