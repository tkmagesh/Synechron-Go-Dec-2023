/* Using WaitGroup for goroutine synchronization */
/* Avoiding data races using sync.Mutex */

/* Encapsulating the concurrent safe state manipulation in custom types (Counter) */
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var counter Counter

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
	fmt.Printf("%d goroutines completed\n", counter.count)
}

func fn(id int, wg *sync.WaitGroup) {
	counter.Increment()
	wg.Done() // decrement the wg counter by 1
}
