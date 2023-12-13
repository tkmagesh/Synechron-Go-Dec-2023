package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4}
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}

	// type inference
	// var nos = [5]int{3, 1, 4, 2, 5}

	// Using :=
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("len(nos) :", len(nos))

	fmt.Println("Iterating using indexer")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Printf("nos[%d] : %d\n", idx, nos[idx])
	}

	fmt.Println("Iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] : %d\n", idx, val)
	}

	/*
		nos2 := nos // creating a copy
		nos2[0] = 9999
		fmt.Printf("nos : %v, nos2 : %v\n", nos, nos2)
	*/

	nos2 := &nos
	// fmt.Println((*nos2)[0])
	fmt.Println(nos2[0])

	fmt.Printf("Before sorting, nos : %v\n", nos)
	sort(&nos)
	fmt.Printf("After sorting, nos : %v\n", nos)
}

func sort(list *[5]int) {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
