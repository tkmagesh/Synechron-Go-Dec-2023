/* Using WaitGroup for goroutine synchronization */

/* Modify the below to execute the f2 concurrently */
package main

import (
	"fmt"
	"time"
)

// modify to execute f1 asynchronously without changing the f1 function (main function can be modified)

func main() {
	fmt.Println("main started")
	f1()
	f2()
	fmt.Println("main completed")
}

// 3rd party API
func f1() {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}
