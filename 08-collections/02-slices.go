package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
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

	// adding new items
	nos = append(nos, 10)
	fmt.Println(nos)

	nos = append(nos, 20, 30, 40, 50)
	fmt.Println(nos)

	hundreds := []int{100, 200, 300}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	// slicing
	fmt.Println("nos[2:5] : ", nos[2:5])
	fmt.Println("nos[2:] : ", nos[2:])
	fmt.Println("nos[:5] : ", nos[:5])

	nos2 := nos
	nos2[0] = 9999
	fmt.Printf("nos : %v, nos2 : %v\n", nos, nos2)
}
