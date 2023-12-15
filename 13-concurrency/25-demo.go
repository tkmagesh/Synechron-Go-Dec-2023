/* Keep generating the fib series until timeout occurs */
package main

import (
	"fmt"
	"time"
)

func main() {
	fibCh := genFib()
	for no := range fibCh {
		fmt.Println(no)
	}
}

func genFib() <-chan int {
	ch := make(chan int)
	// timeoutCh := timeOut(5 * time.Second)
	timeoutCh := time.After(5 * time.Second)
	go func() {
	LOOP:
		for x, y := 0, 1; ; {
			select {
			case <-timeoutCh:
				fmt.Println("timeout occurred")
				break LOOP
			default:
				ch <- x
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}

func timeOut(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}
