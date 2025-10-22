package main

import "fmt"

func main() {
	//bitwise AND operator '&'
	fmt.Println("Bitwise AND Operator &")
	a := 5 //binary: 0101
	b := 3 //binary: 0011
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("a & b =", a&b) //binary: 0001 = decimal 1

	//bitwise OR operator '|'
	fmt.Println("\nBitwise OR Operator |")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("a | b =", a|b) //binary: 0111 = decimal 7

}
