/* Using WaitGroup for goroutine synchronization */
/* Avoiding data races using sync.Mutex */

/* Using sync/atomic package apis for concurrent safe manipulation of intrinsic data */
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var completedCount int64

func main() {
	var count int = 200
	wg := &sync.WaitGroup{}
	fmt.Printf("Starting %d goroutines... Hit ENTER to start!\n", count)
	fmt.Scanln()
	for i := 1; i <= count; i++ {
		wg.Add(1) // increment the wg counter by 1
		go fn(i, wg)
	}
	wg.Wait() // block until the wg counter becomes 0
	fmt.Printf("%d goroutines completed\n", completedCount)
}

func fn(id int, wg *sync.WaitGroup) {
	atomic.AddInt64(&completedCount, 1)
	wg.Done() // decrement the wg counter by 1
}
