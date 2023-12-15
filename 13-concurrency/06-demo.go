/* Using WaitGroup for goroutine synchronization */
/* Avoiding data races using sync.Mutex */

package main

import (
	"fmt"
	"sync"
)

var completedCount int
var mutex sync.Mutex

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
	mutex.Lock()
	{
		completedCount++
	}
	mutex.Unlock()
	wg.Done() // decrement the wg counter by 1
}
