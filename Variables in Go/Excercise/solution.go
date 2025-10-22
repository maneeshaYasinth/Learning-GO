package main

import "fmt"

func buyItems() {
	//values of item 1
	var item1Name string = "laptop"
	var item1Price float64 = 999.99
	var item1Qty int = 2
	var isItem1Available bool = true

	//item1 details print
	fmt.Println("Item 1 Name:", item1Name)
	fmt.Println("Item 1 Price:", item1Price)
	fmt.Println("Item 1 Quantity:", item1Qty)

	//total of item1
	totalOfItem1 := item1Price * float64(item1Qty)

	//values of item
	var item2Name string = "Headphones"
	var item2Price float64 = 199.99
	var item2Qty int = 0
	var isItem2Available bool = false

	//item1 check
	if isItem1Available {
		fmt.Printf("total cost is %.2f\n", totalOfItem1)
	} else {
		fmt.Println(item1Name, "is not available")
	}
	fmt.Println()
	fmt.Println()

	//item2 details print
	fmt.Println("Item 2 Name:", item2Name)
	fmt.Println("Item 2 Price:", item2Price)
	fmt.Println("Item 2 Quantity:", item2Qty)

	//total of item1
	totalOfItem2 := item2Price * float64(item2Qty)

	//item2 check
	if isItem2Available {
		fmt.Printf("total cost is %.2f\n", totalOfItem2)
	} else {
		fmt.Println(item2Name, "is not available")
	}

	fmt.Println()
	fmt.Println()
}

func main() {
	buyItems()
}
