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

}
