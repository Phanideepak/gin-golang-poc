package main

import "fmt"

func main() {
	age := 13
	if age > 21 {
		fmt.Println("Eligible for marriage")
	} else {
		if age > 18 {
			fmt.Println("Eligible for vote")
		} else {
			fmt.Println("Minor")
		}

	}
}
