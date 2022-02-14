package main

import "fmt"

func main() {
	var array_name = [3]int{1, 2, 3}
	fmt.Println(array_name)
	var array_name1 = [5]int{45, 67, 89, 94, 12}
	fmt.Println(array_name1)

	var prices = [3]int{11, 22, 33}
	fmt.Println(prices)
	prices[1] = 66
	prices[2] = 89
	fmt.Println(prices)

	fmt.Println(len(prices))

}
