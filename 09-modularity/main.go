package main

import (
	"fmt"

	"github.com/tkmagesh/synechron-go-dec-2023/09-modularity/calculator"
	/* "github.com/tkmagesh/synechron-go-dec-2023/09-modularity/calculator/utils" */
	ut "github.com/tkmagesh/synechron-go-dec-2023/09-modularity/calculator/utils"
)

func main() {
	fmt.Println("application [module] executed!")
	fmt.Println("Application name :", app_name)
	fmt.Println("Application version :", app_version)

	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println(calculator.Multiply(100, 200))
	fmt.Println(calculator.OpCount())

	// fmt.Println("Is 21 even ?:", utils.IsEven(21))
	fmt.Println("Is 21 even ?:", ut.IsEven(21))
}
