package main

import "fmt"

func main() {
	var x int = 100
	var xPtr *int
	xPtr = &x // address of the data
	fmt.Println(xPtr)

	//accessing the value using the pointer (dereferencing)
	val := *xPtr
	fmt.Println(val)
	fmt.Println(x == *(&x))

	fmt.Println("Before incrementing x :", x)
	fmt.Println("[main] &x :", &x)
	increment(&x)
	fmt.Println("After incrementing x :", x)

	n1, n2 := 100, 200
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap( /* ? */ )
	fmt.Printf("After swapping, n1 = %d and n2 = %d\n", n1, n2)
}

func increment(n *int) {
	fmt.Println("[increment] n :", n)
	*n++
}

func swap( /* ? */ ) /* no return result */ {
	/* ? */
}
