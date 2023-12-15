package main

import (
	"fmt"
)

func main() {
	var ch chan int     // declaration
	ch = make(chan int) //initialization
	go add(100, 200, ch)
	result := <-ch //receive the data from channel
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result // send the data to the channel
}
