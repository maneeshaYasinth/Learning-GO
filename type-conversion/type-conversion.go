package main

import "fmt"

func main() {
	var intVar int = 42
	var floatVar float64 = 3.14

	fmt.Println("Integer:", intVar)
	fmt.Println("Float:", floatVar)

	// Type conversion...................................
	var newFloatVar float64 = float64(intVar)
	var newIntVar int = int(floatVar)

	fmt.Println("Converted Integer:", newIntVar)
	fmt.Println("Converted Float:", newFloatVar)
}
