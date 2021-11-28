package main

import (
	"fmt"
)

type CarDetail struct {
	Price     int
	EngineCap int
}
type Car struct {
	Name    string
	Details CarDetail
}

func (c *Car) GetPrice() int {
	return c.Details.Price
}

func main() {
	var ages = []int{9, 5, 6, 4, 3, 2, 3, 5}

	smallestAge := getSmallestAge(ages)
	biggestAge := getBiggestAge(ages)

	fmt.Printf("Smallest : %d, Biggest : %d\n", smallestAge, biggestAge)
	//  -----------------------------------------------------------------------

	cars := []Car{
		{Name: "Vitz", Details: CarDetail{Price: 500, EngineCap: 1000}},
		{Name: "Peugeot 208", Details: CarDetail{Price: 950, EngineCap: 1500}},
		{Name: "Volvo S40", Details: CarDetail{Price: 1500, EngineCap: 2000}},
		{Name: "Passo", Details: CarDetail{Price: 800, EngineCap: 1000}},
		{Name: "Demio", Details: CarDetail{Price: 670, EngineCap: 1100}},
	}

	cheapestCar := getCheapestCar(cars)
	mostExpensiveCar := getExpensiveCar(cars)

	fmt.Printf("Cheapest Car : %s, price: %d\n", cheapestCar.Name, cheapestCar.GetPrice())
	fmt.Printf("Expensive Car : %s, price: %d\n", mostExpensiveCar.Name, mostExpensiveCar.GetPrice())
	Searchbyprice(cars, 700)

}

// 1

func getSmallestAge(ages []int) int {
	var smallestAge int
	smallestAge = ages[0]
	for _, age := range ages {
		if age < smallestAge {
			smallestAge = age
		}
	}
	return smallestAge
}

// 2.
func getBiggestAge(ages []int) int {
	var biggestAge int
	biggestAge = ages[0]
	for _, age := range ages {
		if age > biggestAge {
			biggestAge = age
		}
	}

	return biggestAge
}

func getCheapestCar(cars []Car) Car {
	var cheapest Car

	cheapest = cars[0]
	for _, car := range cars {
		if car.Details.Price < cheapest.Details.Price {
			cheapest = car
		}
	}
	return cheapest
}

func getExpensiveCar(cars []Car) Car {
	var expensive Car

	// here, write code that will return the most expensive car
	expensive = cars[0]
	for _, car := range cars {
		if car.Details.Price > expensive.Details.Price {
			expensive = car
		}
	}

	return expensive

}

func Searchbyprice(cars []Car, x int) {
	found := false

	for _, car := range cars {
		if car.Details.Price <= x {
			fmt.Printf("name of automobile :%s\n", car.Name)
			found = true
		}
	}
	if !found {
		fmt.Println("do better")
	}

}

/* this function Searchby price takes an array of cars belonging to the struct car and an interger x
it prints out the names of cars whose values are less than equal to x if no car is found it prints out "do better" */
