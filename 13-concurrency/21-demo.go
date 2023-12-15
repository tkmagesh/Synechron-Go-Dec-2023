/*
Write a goroutine that produces n values in the fibonocci series
where n = random(20)
print the generated values in the main() function as and when they are generated
*/

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// modify in such a way that the generated fib no & word are printed as and when they are generated

// consumer
func main() {
	fibCh := genFib()
	for no := range fibCh {
		fmt.Println(no)
	}

	wordCh := genWords()
	for word := range wordCh {
		fmt.Println(word)
	}
}

// producer
func genFib() <-chan int {
	ch := make(chan int)
	go func() {
		for x, y, count, i := 0, 1, rand.Intn(50), 0; i < count; i++ {
			ch <- x
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		}
		close(ch)
	}()
	return ch
}

func genWords() <-chan string {
	str := "Lorem sint eiusmod veniam ullamco incididunt consequat et id sunt consectetur ad elit Dolore quis aute proident pariatur duis nisi aliquip cupidatat Lorem dolor Magna ea sunt sunt deserunt magna Cupidatat occaecat cupidatat excepteur culpa ipsum excepteur consectetur laboris duis id aute voluptate Lorem Voluptate laboris culpa deserunt cupidatat est esse ad ipsum Consectetur sunt enim nostrud mollit sint Ipsum fugiat incididunt voluptate reprehenderitConsequat adipisicing reprehenderit voluptate enim nisi ipsum in veniam Quis aute ea ex sit cupidatat consectetur consequat non consectetur aliqua magna est Excepteur adipisicing et pariatur sint eiusmod duis ipsum Non ad quis amet dolor sunt esse laboris eiusmod ex eiusmod est laborum non Consequat mollit aliqua aliquip duis duis consequat consequat consectetur cupidatat quis laboris nisi adipisicing Dolor incididunt dolor elit anim do eu et cillum duis et est elit Fugiat reprehenderit adipisicing duis laboris labore"
	ch := make(chan string)
	go func() {
		words := strings.Split(str, " ")
		for _, word := range words {
			ch <- word
			time.Sleep(300 * time.Millisecond)
		}
	}()
	return ch
}
