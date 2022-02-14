package main

import "fmt"

func mathematicalOperations(a int, b int) (s int, d int, p int, q int) {
	s = a + b
	d = a - b
	p = a * b
	if a < b {
		temp := b
		b = a
		a = temp
	}
	q = a / b
	return
}
func main() {
	A, B, C, D := mathematicalOperations(34, 97)
	fmt.Println(A, B, C, D)
}
