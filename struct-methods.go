package main

import "fmt"

type person struct {
	Name string
}

/*
When we create a method with a struct (without pointer),
it creates a copy of the struct and then updates the copy of the struct.
In the example below, it makes a copy of person2,
makes a change on the copy and the original value of person2 is not changed.

-- this is ideal in cases where you just want to read data from the struct without updating it.
*/
func (p person) changeNameToPatrick() {
	p.Name = "patrick"
	fmt.Println("change without pointer", p.Name)
}

/*
When we create a method with a pointer to a struct,
it actually makes changes to the actual struct.
In the example below, it actually changes the name of person1 and person2

--- This is the ideal option in most cases especially when you want to update the values of the struct
*/

func (p *person) changeNameToJane() {
	p.Name = "jane"
	fmt.Println("change with pointer", p.Name)

}

func main() {
	var person1 person
	person1.changeNameToJane()
	fmt.Println(person1.Name)

	var person2 person
	person2.changeNameToJane()
	person2.changeNameToPatrick()
	fmt.Println(person2.Name)

}
