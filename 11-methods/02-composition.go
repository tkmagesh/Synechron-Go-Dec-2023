package main

import (
	"fmt"
)

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
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

// Overriding the "Format()" method of the base struct (Product)
func (pp *PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expirty = %q", pp.Product.Format(), pp.Expiry)
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

	fmt.Println(milk.Format())
	milk.ApplyDiscount(10)
	fmt.Println(milk.Format())
}
