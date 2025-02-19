package main

import "fmt"

func main() {
	var input int
	fmt.Println("Enter your input from range 1-7: ")
	fmt.Scan(&input)

	switch input {
	case 1, 2, 3:
		res1 := input * 5
		fmt.Printf("The output is: %d\n", res1)
	case 4, 5:
		res2 := input * 10
		fmt.Printf("The output is: %d\n", res2)
	case 6, 7:
		res3 := input * 15
		fmt.Printf("The output is: %d\n", res3)
	default:
		fmt.Println("Input not in range 1-7")
	}
}