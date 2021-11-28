package main

//a.  this is ok
type Fruit struct {
	Name       string
	Maturinty  int
	Yield      int
	Priceperkg int
}

// c. ok
func (f *Fruit) incomePerHarvest() int {
	return f.Priceperkg * f.Yield
}
func (f *Fruit) incomeIn400Days() int {
	return f.incomePerHarvest() * 400 / f.Maturinty
}

func main() {
	var mangoes = Fruit{

	}
	var apples = Fruit {

	}
	var oranges = Fruit {

	}
	fruits := []Fruit{
		mangoes, apples, oranges,
		// {Name: "Mangoes", Maturinty: 100, Yield: 1000, Priceperkg: 12},
		// {Name: "Apples", Maturinty: 200, Yield: 200, Priceperkg: 56},
		// {Name: "Oranges", Maturinty: 50, Yield: 3000, Priceperkg: 47},
	}

}
func mostProfitable(fruits []Fruit) {
	highest := fruits[0].incomeIn400Days()
	
}
