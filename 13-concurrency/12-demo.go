/* Using WaitGroup for goroutine synchronization */

/* Modify the below to execute the f2 concurrently */
package main

import (
	"fmt"
	"sync"
	"time"
)

// modify to execute f1 asynchronously without changing the f1 function (main function can be modified)

func main() {
	fmt.Println("main started")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		f1()
		wg.Done()
	}()
	f2()
	wg.Wait()
	fmt.Println("main completed")
}

// 3rd party API
func f1() {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f2 completed")
}
