package main

import "fmt"

func factorial(x float64) (y float64) {
	if x <= 0 {
		return 1
	}
	return x * factorial(x-1)
}
func main() {

	var num float64 = 7
	if num < 0 {
		fmt.Println("negative number ki factorial untunda bro...")
	} else {
		fmt.Println(factorial(num))
	}

}
