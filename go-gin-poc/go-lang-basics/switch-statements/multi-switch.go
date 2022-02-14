package main

import "fmt"

func main() {
	today := 3

	switch today {
	case 1, 3, 5:
		fmt.Println("Odd week day")
	case 2, 4:
		fmt.Println("even week day")
	case 6, 0:
		fmt.Println("weekends")
	default:
		fmt.Println("invalid week day")
	}
}
