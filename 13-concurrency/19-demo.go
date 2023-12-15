/*
Write a goroutine that produces n values in the fibonocci series
where n = random(20)
print the generated values in the main() function as and when they are generated
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	go genFib(ch)
	for no := range ch {
		fmt.Println(no)
	}
}

func genFib(ch chan int) {
	for x, y, count, i := 0, 1, rand.Intn(20), 0; i < count; i++ {
		ch <- x
		time.Sleep(500 * time.Millisecond)
		x, y = y, x+y
	}
	close(ch)
}
