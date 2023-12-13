/*
Refactor into functions
Make sure the functions have ONLY ONE responsibility (SRP)
*/

package main

import "fmt"

func main() {

	var n1, n2, userChoice, result int
	for {
		fmt.Println("1.Add\n2.Subtract\n3.Multply\n4.Divide\n5.Exit")
		fmt.Println("Enter your choice :")
		fmt.Scanln(&userChoice)
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		if userChoice == 5 {
			break
		}
		fmt.Println("Enter the operands :")
		fmt.Scanln(&n1, &n2)
		switch userChoice {
		case 1:
			result = n1 + n2
		case 2:
			result = n1 - n2
		case 3:
			result = n1 * n2
		case 4:
			result = n1 / n2
		}
		fmt.Println("Result :", result)
	}
}
