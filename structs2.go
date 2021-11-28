package main

import "fmt"

type Car struct {
	Model string
	Brand string
	Price float64
}

func (c *Car) setcardetails(model, brand string, price float64) {
	c.Model = model
	c.Brand = brand
	c.Price = price
	//fmt.Printf("...Method ...Model...%s\n,Brand....%s\n, Price....%f\n", c.Model, c.Brand, c.Price)

}

func Boogie(car1, car2 Car) {
	if car1.Price > car2.Price {
		fmt.Println(car1.Model)
	} else {
		fmt.Println(car2.Model)

	}

}

func main() {
	var car1 Car
	car1.setcardetails("toyota", "lexus", 3549.9)

	// 	car1 := Car{
	// Model: "Toyota",
	// Brand: "Lexus",
	// Price: 235.56,
	// }
	// fmt.Println("hello ")
	// car1.setcardetails()
	var car2 Car
	car2.setcardetails("jeep", "wrangler", 1243)

	// car2 := Car{
	// 	Model: "Nissan",
	// 	Brand: "Juke",
	// 	Price: 235.56,
	// 	}
	// fmt.Println()
	// car2.setcardetails()

	Boogie(car1, car2)

}	
