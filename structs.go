package main

import (
	"encoding/json"
	"fmt"
)

// we use the keyword type when we want to define our own data type
type s string
type a []int

// a struct holds data elements of different data types
// Person is our data type
// Elements of a struct can begin with capital/ small letters. The difference is
// Anything that starts with a small letter can only be accessed locally (methods)
// Anything that starts with a Capital letter can be accessed anywhere

// A struct can have methods:
// a function is usually independent, a method is usually dependent on another element e.g a struct
type Person struct {
	Name   string
	Age    int
	Cars   []string
	House  House
	salary float32
}

type House struct {
	Address string
	Price   float64
}

func printNameAndAge(name string, age int) {
	fmt.Printf("----Name: %s, Age : %d\n", name, age)
}

func (person *Person) PrintNameAndAge() {
	fmt.Printf("----Method -- Name: %s, Age : %d\n", person.Name, person.Age)
}

func (p *Person) ModifyAge(newAge int) {
	p.Age = newAge
}

func (p *Person) SetSalary(salary float32) {
	p.salary = salary
}

func (p *Person) GetSalary() float32 {
	return p.salary
}

func main() {
	//var name s
	//name = "Sam"
//	var ages a
	//ages = a{1, 2}
	fmt.Println("hello world ", name, ages)
	// ---- let's define a person
	var person1 Person
	person1 = Person{
		Name: "Habiba",
		Age:  25,
	}
	//person1.Age = 43
	//printNameAndAge("Kim Jong Un", 56) // independent function

	person1.PrintNameAndAge() // depends on the person1 data
	person1.ModifyAge(45)
	person1.PrintNameAndAge()

	person1.SetSalary(456.87)
	s := person1.GetSalary()
	//s = 0.0 // - this won't affect person1
	fmt.Println(s, person1.salary)

	// if we want to access elements, we do person1.Name e.g
	fmt.Println(person1.Name, person1.Age)
	person1.Name = "Yvonne"
	fmt.Println(person1.Name, person1.Age)

	// the default value of a struct is not nil, but each element takes a default value
	var person2 Person
	fmt.Println(person2.Name, person2.Age)
	person2.PrintNameAndAge()

	// if you want to define and assign a variable at the same time
	// e.g var x = 67
	// or x := 67
	var person3 = Person{
		Name: "Samson",
		Age:  25,
		Cars: []string{"Mazda"},
		House: House{
			Address: "Ruaka",
			Price:   34.43,
		},
	}

	fmt.Println(person3)

	// when printing, make sure to print individual elements.
	fmt.Printf("The name is %s and the age is %d\n", person3.Name, person3.Age)

	// the other alternative is, convert the data to a JSON byte and then print the string version of the JSON byte
	person3JSON, err := json.Marshal(person3)
	if err != nil {
		fmt.Printf("The data is invalid")
		return
	}
	fmt.Println(string(person3JSON))

}
