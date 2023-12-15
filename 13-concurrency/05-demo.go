/* Using WaitGroup for goroutine synchronization */

/* Modify the below to execute the f2 concurrently */
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int
	wg := &sync.WaitGroup{}
	flag.IntVar(&count, "count", 0, "Number of goroutines to execute")
	flag.Parse()
	fmt.Printf("Starting %d goroutines... Hit ENTER to start!\n", count)
	fmt.Scanln()
	for i := 1; i <= count; i++ {
		wg.Add(1) // increment the wg counter by 1
		go fn(i, wg)
	}
	wg.Wait() // block until the wg counter becomes 0
	fmt.Println("All goroutines completed")
}

func fn(id int, wg *sync.WaitGroup) {
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
	wg.Done() // decrement the wg counter by 1
}
