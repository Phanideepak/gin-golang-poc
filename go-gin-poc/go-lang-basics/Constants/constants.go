package main

import "fmt"

// const keyword declares the variable as "constant", which means that it is unchangeable and read-only
// The value of a constant must be assigned when you declare it.
// multiple consts can be declared in const () block
func main() {
	const a int = 40
	const b = 30

	var d = a + b
	fmt.Println(d, a, b)

	const (
		m  int = 40
		n      = 30
		pi     = "3.1414"
	)
	fmt.Println(m*n*m*2, m, n, pi)

}
