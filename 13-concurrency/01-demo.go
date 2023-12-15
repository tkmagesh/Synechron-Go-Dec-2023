package main

import (
	"fmt"
)

func main() {
	go f1() //scheduling the execution of this function through the scheduler
	f2()
	// blocking the execution of this function for 1 second so that the scheduler can look for other goroutines scheduled and execute them.

	// DO NOT DO THIS IN PRODUCTION APP
	// time.Sleep(1 * time.Second)
	fmt.Scanln()
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
