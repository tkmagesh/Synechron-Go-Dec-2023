package main

import "fmt"

func main() {
	exec(f1) // to execute f1
	exec(f2) // to execute f2
	// exec(f3)
	exec(func() {
		fmt.Println("anonymous fn invoked")
	})
}

func exec(fn func()) {
	fn()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
