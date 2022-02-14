package main

import "fmt"

func main() {
	sizes := [2]string{"big", "small"}
	fruits := [3]string{"apple", "grapes", "banana"}
	for i := 0; i < len(sizes); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(sizes[i], fruits[j])
		}
	}
}
