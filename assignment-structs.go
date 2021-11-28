/*
Habiba is a business woman who sells handbags for a living. Each bag has the following characteristics:
- name
- brand
- price.

Given the following snippet, complete the 'process' function so that it returns the bag with the lowest price.
*/

package main

import "fmt"

type Bag struct {
	Name  string
	Brand string
	Price float32
}

func (b *Bag) PrintDetails() {
	details := fmt.Sprintf("Name : %s\tBrand: %s\t\t\tPrice: $%3.2f\n", b.Name, b.Brand, b.Price)
	fmt.Println(details)
}

func process(bags []Bag) (cheapest Bag) {
	cheapest = bags[0]

	/*
		In Go we can loop through an array in two ways

		1.
		for i := 0; i < len(bags); i++ {
			fmt.Println(bags[i])
		}

		2.
		for i, bag := range bags {
			fmt.Println(i, bag) // here. bag is equivalent to bags[i]
		}
	*/
	for _, bag := range bags {
		if bag.Price < cheapest.Price {
			cheapest = bag
		}
	}

	// for i := 0; i < len(bags); i++ {
	// 	if bags[i].Price < cheapest.Price {
	// 		cheapest = bags[i]
	// 	}
	// }

	// code here
	return cheapest
}

func main() {
	bags := []Bag{
		{Name: "Clutch Bag", Brand: "Chanele", Price: 45.4},
		{Name: "Handbag", Brand: "Gucci", Price: 34.4},
		{Name: "Slingbag", Brand: "Prada", Price: 56.78},
		{Name: "Backpack", Brand: "King's Collection", Price: 30.98},
		{Name: "Handbag", Brand: "Chanele", Price: 67.65},
	}
	cheapestBag := process(bags)
	cheapestBag.PrintDetails()
}
