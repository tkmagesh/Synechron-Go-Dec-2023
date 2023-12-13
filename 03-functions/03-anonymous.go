package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hi")
	}()

	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	// convert the fullNameGreet, getFullNameGreet & divide into anonymous functions

	(func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	})("Magesh", "Kuppan")

	/* with 2 arguments & 1 return result */
	msg := func(firstName, lastName string) string {
		return fmt.Sprintf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}("Suresh", "Kannan")
	fmt.Print(msg)

	q, r := func(x, y int) (quotient int, remainder int) {
		quotient, remainder = x/y, x%y
		return
	}(100, 7)
	fmt.Println(q, r)

	sumResult := func(nos ...int) int {
		var result int
		for i := 0; i < len(nos); i++ {
			result += nos[i]
		}
		return result
	}(10, 20, 30, 40, 50)
	fmt.Println(sumResult)
}
