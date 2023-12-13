package main

import "fmt"

func main() {
	sayHi()
	greet("Magesh")
	fullNameGreet("Magesh", "Kuppan")
	msg := getFullNameGreet("Suresh", "Kannan")
	fmt.Print(msg)

	fmt.Println(divide(100, 7))

	/*
		q, r := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
	*/

	// Using ONLY the quotient
	// Use _ to receive the result from a function and ignore the result
	/*
		q, _ := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d\n", q)
	*/

}

/* with no arguments & no return results */
func sayHi() {
	fmt.Println("Hi")
}

/* with 1 argument & no return results */
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

/* with 2 arguments & no return results */
/*
func fullNameGreet(firstName string, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}
*/

func fullNameGreet(firstName, lastName string) {
	fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

/* with 2 arguments & 1 return result */
func getFullNameGreet(firstName, lastName string) string {
	return fmt.Sprintf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}

/* with 2 arguments & 2 return results */
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

// using named results
func divide(x, y int) (quotient int, remainder int) {
	/*
		quotient = x / y
		remainder = x % y
	*/
	quotient, remainder = x/y, x%y
	return
}
