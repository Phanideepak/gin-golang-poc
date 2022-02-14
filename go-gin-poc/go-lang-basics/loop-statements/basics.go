package main

import "fmt"

func main() {
	num := 20
	for i := 0; i <= 100; i = i + 10 {
		fmt.Println(num + i)
	}
	fmt.Println("Print odd numbers upto 20")

	for i := 1; i <= 20; i = i + 1 {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
