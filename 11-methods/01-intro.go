package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

// pointer receiver - when the state of the object has to be updated
func (p *Product) ApplyDiscount(discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

func main() {
	pen := Product{Id: 100, Name: "Pen", Cost: 10}
	// fmt.Println(Format(pen))
	fmt.Println(pen.Format())
	fmt.Println("After applying 10% discount")
	// ApplyDiscount(&pen, 10)
	// (&pen).ApplyDiscount(10) //=> NOT needed
	pen.ApplyDiscount(10)
	// fmt.Println(Format(pen))
	fmt.Println(pen.Format())
}
