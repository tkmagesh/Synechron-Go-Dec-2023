package main

import "fmt"

// modify the below so that the receive operation is done in a goroutine
func main() {
	ch := make(chan int)
	go func() {
		ch <- 100
	}()
	data := <-ch
	fmt.Println(data)
}
