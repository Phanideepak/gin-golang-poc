package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int = 23, 56, 67, 90
	fmt.Println(a, b, c, d)
	var x, y, z = 45, 89.802, "Welcome"
	fmt.Println(x, y, z)

	// variable block to increase readibility.
	var (
		m int
		n int     = 90
		p float32 = 90.87
		q float32
	)
	fmt.Println("Variable block")
	fmt.Println(m, n, p)
	r := 90.79
	fmt.Println(r, q)

}
