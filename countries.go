package main

import "fmt"

type Country struct {
	Name         string
	Independence int
	Population   int
}

func (c *Country) yearsofindependence() int {
	return 2021 - c.Independence
}

func main() {

	country1 := Country{
		Name:         "Austria",
		Independence: 1800,
		Population:   123,
	}
	fmt.Println(country1.Name, country1.Independence, country1.Population)
	x := country1.yearsofindependence()
	fmt.Println(x)
	country2 := Country{
		Name:         "Kenya",
		Independence: 1963,
		Population:   41,
	}
	fmt.Println(country2.Name, country2.Independence, country2.Population)
	//country2.yearsofindependence()
	country3 := Country{
		Name:         "Ghana",
		Independence: 1967,
		Population:   51,
	}
	fmt.Println(country3.Name, country3.Independence, country3.Population)
	//country3.yearsofindependence()

	longestIndependent := getLongestIndependent(country1, country2, country3)
	fmt.Printf("Longest independent country is %s with %d years\n", longestIndependent.Name, longestIndependent.yearsofindependence())
	/*

	 */
}

func getLongestIndependent(country1, country2, country3 Country) (longestIndependent Country) {
	// 1. Assume that country1 is the longest independent
	longestIndependent = country1

	// 2. Compare all countries against longestIndependent. Note: don't use else if here, use independent ifs
	if country2.yearsofindependence() > longestIndependent.yearsofindependence() {
		longestIndependent = country2
	}
	if country3.yearsofindependence() > longestIndependent.yearsofindependence() {
		longestIndependent = country3
	}

	return longestIndependent

}

// func longestindependence() {
// 	var x := "longest"
// 	"longest" = country1
// 	if country2 > "longest" {
// 		return country2.Name
// 	}
// 	if country3 > "longest" {
// 		return  country3.Name
// 	}else {
// 		return country1.Name
// 	}
// }
