package main

import "fmt"

func main() {
	a := 45
	b := 562
	c := 67
	if a > b && a > c {
		fmt.Printf("largest is %v\n", a)
	} else if b > c {
		fmt.Printf("largest is %v\n", b)
	} else {
		fmt.Printf("largest is %v\n", c)
	}
}
