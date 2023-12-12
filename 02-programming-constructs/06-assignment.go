/*
Build an interactive calculator

Display the following menu
	1. Add
	2. Subtract
	3. Multiply
	4. Divide
	5. Exit

Accept the user choice

if the user choice == 5
	exit the application

if the user choice is 1 to 4
	accept 2 numbers
	perform the corresponding operation
	print the result
	IMPORTANT : display the menu again

if the user input is not (1-5)
	print "Invalid choice"
	IMPORTANT : display the menu again
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
