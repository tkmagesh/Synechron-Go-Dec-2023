// deadlock

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(10)
	go f1(wg)
	wg.Wait() // counter will never become zero and the application will indefinitely wait and there by resulting in a deadlock
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("f1 invoked")
}
