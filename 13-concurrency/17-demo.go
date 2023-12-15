// streaming data through channels

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go genData(ch)
	for {
		if data, isOpen := <-ch; isOpen {
			fmt.Println(data)
			time.Sleep(500 * time.Millisecond)
			continue
		}
		break
	}
}

func genData(ch chan int) {
	count := rand.Intn(20)
	for i := 1; i <= count; i++ {
		ch <- i * 10
	}
	close(ch)
	fmt.Println("[producer] - all data produced & channel closed")
}
