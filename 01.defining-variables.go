package main

import "fmt"

func main() {
	/*

		Recap:
		In formatted string,
		%d stands for integers
		%s strings
		%b  boolean (this is for Go)
		%f float


		 Arithmetics
		 -----------
		1. Addition
		2. Subtraction
		3. Multiplication
		4. Division
		5. Modulus
		6. Increment (pre-increment and post-increment) ** go does not have pre
		7. Decrement (pre-decrement and post-decrement)

		-- If you do any of these operations to values with the same data type, the answer will be of the same data type
		-- To avoid confusion, always use values of the same data type
		-- If you have to use different data type, cast the convenient one
		-- If you cast 4 to a float, it will be 4.0

		-- Compatible casting
		1. Float and Integer (it is prefered to cast integers to floats, not the other way round. to avoid losing data)
		2. Strings and floats/integers (casting strings to integers/float, the string has to contain valid numbers only)
		e.g you can cast "45" to 45 or "54.32" to 54.32 but you cannot cast 45x
		- note that, in Go there are special functions to convert between integers/floats and strings
		- *** Read on : - fmt.Sprint, fmt.Sprintf (integers to strigns), strconv.Atoi, strconv.ParseFloat (strings to integers/floats to strings)
	*/

	// 1. Addition
	var a = 65
	var b = 43
	var c = a + b
	fmt.Println("a + b = ", c)
	fmt.Printf("4 + 7 = %d\n", 4+7)

	var s1 = "No"
	var s2 = "Yes"
	var s = s1 + s2
	fmt.Printf("adding strings : %s\n", s)

	// this is the same for subtraction, multiplication and division
	// a * b
	// a - b
	// a / b

	// Modulus gives the remainder of a division. Mostly used for divisibility tests (if the answer is 0 in a % b, a is divisible by b)
	// e.g
	var j = 102
	var k = 3
	var i = j % k
	fmt.Println("j % k =", i)

	// shorthand for these
	// a = a + b (a += b)
	// a = a * b (a *= b)
	// a = a - b (a -= b) note that a = b - a remains that way.
	// a = a / b (a /= b)
	/*
		numberOfPeople := 50
		newPeople := 10
		numberOfPeople = numberOfPeople + newPeople // numberOfPeople += newPeople
	*/

	// Increment
	// post-increment
	a = 10
	b = 10

	fmt.Println("The value of a + 1 is", a+1)
	fmt.Println("The value of a is", a)

	a++
	fmt.Println("The value of a is", a)

}
