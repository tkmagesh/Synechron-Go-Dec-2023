package main

import (
	"fmt"
	"sync"
)

func main() {
	var ch chan int     // declaration
	ch = make(chan int) //initialization
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg, ch)
	wg.Wait()
	result := <-ch //receive the data from channel
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	result := x + y
	ch <- result // send the data to the channel
}
