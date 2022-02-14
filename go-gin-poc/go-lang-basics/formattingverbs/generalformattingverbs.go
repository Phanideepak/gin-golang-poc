package main

import (
	"fmt"
)

func main() {
	a := 20
	b := 90.67
	c := "Deadad"

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%#v\n", c)

	fmt.Printf("%v%%\n", a)
	fmt.Printf("%v%%\n", b)
	fmt.Printf("%v%%\n", c)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}
