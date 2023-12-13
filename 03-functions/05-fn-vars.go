package main

import "fmt"

func main() {
	var sayHi func()
	sayHi = func() {
		fmt.Println("Hi")
	}
	sayHi()

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	var fullNameGreet func(string, string)
	fullNameGreet = func(firstName, lastName string) {
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}
	fullNameGreet("Magesh", "Kuppan")

	var getFullNameGreet func(string, string) string
	getFullNameGreet = func(firstName, lastName string) string {
		return fmt.Sprintf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	}
	msg := getFullNameGreet("Suresh", "Kannan")
	fmt.Print(msg)
	/*


		// convert the fullNameGreet, getFullNameGreet & divide into anonymous functions






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
	*/
}
