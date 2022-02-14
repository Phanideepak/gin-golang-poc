package main

import (
	"fmt"
)

// := is for declaration + assignment
// For example, var foo int = 10 is the same as foo := 10
// assignment operator-     =
// := can only be used inside function.

// static type language.
func main() {
	var a int = 32
	var b int = 45
	var c string = "Welcome"
	var d float32 = 34.90
	var e bool = false
	var f float32 = d + 90
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	f = f + 91
	fmt.Println(f)
}
