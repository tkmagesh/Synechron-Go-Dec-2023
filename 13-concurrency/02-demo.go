package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //scheduling the execution of this function through the scheduler
	f2()
	// blocking the execution of this function for 1 second so that the scheduler can look for other goroutines scheduled and execute them.

	// DO NOT DO THIS IN PRODUCTION APP
	time.Sleep(4 * time.Second)

}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
