package main

import "fmt"

func main() {
	/*
		var productRanks map[string]int
		productRanks = make(map[string]int)
		productRanks["pen"] = 3
		productRanks["pencil"] = 1
		productRanks["marker"] = 5
	*/

	// var productRanks map[string]int = map[string]int{"pen": 3, "pencil": 1, "marker": 5}
	// var productRanks = map[string]int{"pen": 3, "pencil": 1, "marker": 5}
	// productRanks := map[string]int{"pen": 3, "pencil": 1, "marker": 5}
	productRanks := map[string]int{
		"pen":          3,
		"pencil":       1,
		"marker":       5,
		"scribble-pad": 4,
	}
	fmt.Println(productRanks)
	fmt.Println("len(productRanks) :", len(productRanks))

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	productNames := make([]string, 0, 0)
	for key, _ := range productRanks {
		productNames = append(productNames, key)
	}
	fmt.Println("productNames :", productNames)

	// removing an item
	// keyToRemove := "pencil" //existing product
	keyToRemove := "stappler" //non-existing product
	fmt.Printf("After removing '%s'\n", keyToRemove)
	delete(productRanks, keyToRemove)
	fmt.Println(productRanks)

	// check for the existense of a key
	keyToCheck := "pencil" //existing product
	// keyToCheck := "stappler" //non-existing product
	if val, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, val)
	} else {
		fmt.Printf("Key %q does not exist\n", keyToCheck)
	}

	dupProductRanks := productRanks // NOT creating a copy
	dupProductRanks["pen"] = 9999
	fmt.Println("productRanks :", productRanks)
	fmt.Println("dupProductRanks :", dupProductRanks)

}
