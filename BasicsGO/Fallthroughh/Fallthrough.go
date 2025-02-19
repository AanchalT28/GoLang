package main

import "fmt"

func main() {

	//user input
	var errorlvl int
	fmt.Println("Enter the error level(1-4): ")
	fmt.Scan(&errorlvl)

	switch errorlvl {
	case 1:
		fmt.Println("Level 1: Info")
		fallthrough //fallthrough to next case even if condition doesnt match
	case 2:
		fmt.Println("Level 2: Warning")
		fallthrough
	case 3:
		fmt.Println("Level 3: Error")
		fallthrough
	case 4:
		fmt.Println("Level 4: Critical error")
	case 5:
		fmt.Println("Invalid Input :[")
	}
}
