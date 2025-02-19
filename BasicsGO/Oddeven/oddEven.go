package main

import "fmt"

func main() {

	var n int
	fmt.Println("Enter the number of elements in array: ")
	fmt.Scan(&n)

	// Declare slice based on user input size
	numbers := make([]int, n)

	fmt.Println("Enter the numbers in the array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Element %d: ", i+1)
		fmt.Scan(&numbers[i]) // Use & to pass the address
	}

	// Declare slices for odd and even numbers
	var oddnum []int
	var evennum []int

	for _, num := range numbers {
		switch {
		case num%2 == 0:
			evennum = append(evennum, num)
		default:
			oddnum = append(oddnum, num)
		}
	}
	fmt.Println("Even Numbers: ", evennum)
	fmt.Println("Odd Numbers: ", oddnum)
}
