package main

import "fmt"

func main() {
	var a = make(map[string]string)
	a["brand"] = "nissan"
	a["color"] = "yellow"
	fmt.Println(a)

	var b = map[string]string{"brand": "audi", "color": "black"}
	fmt.Println(b)

	b["year"] = "1999"
	b["manufactured"] = "adidas"
	b["color"] = "black"

	fmt.Println(b)

	delete(b, "manufactured")

	fmt.Println(b)

	// we can check the existence of key like this: val, isKeyExists :=map_name[key]
	carcolor, isKey1Exists := b["color"]
	fmt.Println(carcolor, isKey1Exists)
	_, isKey2Exists := b["manufactured"]
	fmt.Println(isKey2Exists)

	// iterating over maps.
	for k, v := range b {
		fmt.Println(k, v)
	}

}
