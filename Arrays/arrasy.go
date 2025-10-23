package main

import "fmt"

func main() {

	var numbers = [5]int{34, 53, 65, 66, 42}
	fmt.Println(numbers[2])

	var fruits [3]string
	fruits[0] = "banana"
	fruits[1] = "apple"
	fruits[2] = "orange"

	fmt.Println("fruit array:", fruits)

}
