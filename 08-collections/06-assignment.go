/*
Find the "word size" that occurs the most in the given string (which size words occurs the most)

Print the "word size" and "# of occurences"

ex output:

	4 letter words occurs the most with 32 occurrances

Hint:

	use strings.Split() to split the string into words
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Pariatur nisi nostrud id ipsum esse minim velit Lorem eiusmod est reprehenderit sint Esse non proident labore dolore labore nisi dolore dolor proident voluptate minim laborum Ad irure ut Lorem exercitation Ipsum reprehenderit consequat veniam non reprehenderit ut eiusmod magna magna aliquip ut sint Mollit irure fugiat nostrud non excepteur aliquip cillum Labore dolore deserunt irure ea tempor Consequat ea quis ipsum esse minim reprehenderit amet Qui do dolor do anim occaecat culpa commodo sunt mollit excepteur laboris tempor Ullamco labore quis culpa laborum sunt Culpa incididunt ad consectetur ut deserunt tempor proident ut Qui quis aliqua fugiat sunt dolore commodo reprehenderit veniam tempor Voluptate sit laboris sunt et do sint ea irure veniam duis laborum Quis irure id officia inNisi velit proident cupidatat laborum velit enim qui deserunt consectetur Ad ea quis veniam cupidatat tempor sunt elit velit Sit anim irure sunt ut exercitation excepteur elit laboris occaecat dolor Duis ad commodo ut quis nisi pariatur anim Fugiat ad id anim velit labore"

	var maxCount, maxSize int
	wordsCountBySize := make(map[int]int)
	words := strings.Split(str, " ")
	for _, word := range words {
		key := len(word)
		wordsCountBySize[key]++
	}

	for size, count := range wordsCountBySize {
		if count > maxCount {
			maxCount, maxSize = count, size
		}
	}
	fmt.Printf("%d letter words occurs the most with %d occurrances\n", maxSize, maxCount)
}
