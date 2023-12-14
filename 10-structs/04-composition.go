package main

import (
	"fmt"
)

type Product struct {
	Id   int
	Name string
	Cost float32
}

type Dummy struct {
	Name string
}

type PerishableProduct struct {
	// Dummy
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {
	// milk := PerishableProduct{Product{100, "Milk", 50}, "2 Days"} // Not advisable
	milk := PerishableProduct{
		Product: Product{
			Id:   100,
			Name: "Milk",
			Cost: 50,
		},
		Expiry: "2 Days",
	}
	/*
		fmt.Println(milk.Product.Id)
		fmt.Println(milk.Product.Name)
		fmt.Println(milk.Product.Cost)
	*/
	fmt.Println(milk.Id)
	fmt.Println(milk.Name)
	fmt.Println(milk.Cost)
	fmt.Println(milk.Expiry)

	eggs := NewPerishableProduct(102, "Eggs", 75, "1 week")
	fmt.Println(eggs.Id)
	fmt.Println(eggs.Name)
	fmt.Println(eggs.Cost)
	fmt.Println(eggs.Expiry)

}
