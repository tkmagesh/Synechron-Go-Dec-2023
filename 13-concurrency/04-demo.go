/* Using WaitGroup for goroutine synchronization */

/* Modify the below to execute the f2 concurrently */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1) // increment the wg counter by 1
	go f1(wg)

	wg.Add(1) // increment the wg counter by 1
	go f2(wg)

	wg.Wait() // block until the wg counter becomes 0
}

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the wg counter by 1
}

func f2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("f2 invoked")
}
