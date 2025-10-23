package main

// pointer is a variable that stores the memory address of another variable.

import "fmt"

//using pointers to increment a value of a variable

func incremetValue(value *int) {
	*value++
	fmt.Println("incremented value is ", *value)
}

func main() {
	var num int = 42
	var ptr *int = &num
	fmt.Println("Value is: ", num)
	fmt.Println("Address is: ", ptr)

	//dereferencing the pointer
	fmt.Println("Dereferenced Value is:", *ptr)
	incremetValue(ptr)

}
