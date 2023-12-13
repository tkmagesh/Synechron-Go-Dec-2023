/*
Refactor into functions
Make sure the functions have ONLY ONE responsibility (SRP)
*/

package main

import "fmt"

func main() {

	var userChoice int
	for {
		userChoice = getUserChoice()
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		if userChoice == 5 {
			break
		}
		doOperation(userChoice)
	}
}

func getUserChoice() (userChoice int) {
	fmt.Println("1.Add\n2.Subtract\n3.Multply\n4.Divide\n5.Exit")
	fmt.Println("Enter your choice :")
	fmt.Scanln(&userChoice)
	return
}

func getOperands() (n1, n2 int) {
	fmt.Println("Enter the operands :")
	fmt.Scanln(&n1, &n2)
	return
}

func add(x, y int) int {
	return x + y
}

func subntract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func doOperation(userChoice int) {
	var result int
	n1, n2 := getOperands()
	switch userChoice {
	case 1:
		result = add(n1, n2)
	case 2:
		result = subntract(n1, n2)
	case 3:
		result = multiply(n1, n2)
	case 4:
		result = divide(n1, n2)
	}
	fmt.Println("Result :", result)
}
