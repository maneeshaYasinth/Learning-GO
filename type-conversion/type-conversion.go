package main

import "fmt"

func calc() {
	var a int = 10
	var b float64 = 5.5

	var sum float64 = float64(a) + b
	fmt.Println("Sum as float:", sum)

	var sum2 int = a + int(b)
	fmt.Println("sum as int = ", sum2)

}

func main() {
	var intVar int = 42
	var floatVar float64 = 3.14

	fmt.Println("Integer:", intVar)
	fmt.Println("Float:", floatVar)

	// Type conversion...................................
	var newFloatVar float64 = float64(intVar)
	var newIntVar int = int(floatVar)
	// ................................................

	fmt.Println("Converted Integer:", newIntVar)
	fmt.Println("Converted Float:", newFloatVar)

	calc()
}
