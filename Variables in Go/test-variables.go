package main

import "fmt"

//blind write this😅
func mathsOps() {
	fmt.Println(10 + 4)
	fmt.Println(10 - 4)
	fmt.Println(10 * 4)
	fmt.Println(10 / 4)
}

//using  const to calculate area of circle

func careaOfCircle() {
	const pi = 3.14
	var radius float64 = 5
	area := pi * radius * radius
	fmt.Println("Area of a circle = ", area)
}

func dataTypes() {
	var name string = "maneeesha"
	var a int = 2
	var b int = 10
	var isTrue bool = true
	fmt.Println("----------------------------------------")

	fmt.Println("Name:", name)
	fmt.Println("a + b =", a+b)
	fmt.Println("Is True?", isTrue)
	fmt.Println("Is True?", isTrue)
	fmt.Println("----------------------------------------")

	//to check data types
	//🔴🔴SHOULD USE printf INSTEAD OF PRINTLN
	fmt.Printf("Data type of name: %T\n", name)
	fmt.Printf("Data type of a: %T\n", a)
	fmt.Printf("Data type of isTrue: %T\n", isTrue)
	fmt.Println("----------------------------------------")

	//assigning values without var keyword
	c := 15.5
	fmt.Println("Value of c:", c)
	fmt.Printf("Data type of c: %T\n", c)
	fmt.Println("----------------------------------------")

}

func main() {
	mathsOps()
	dataTypes()
	careaOfCircle()
}
