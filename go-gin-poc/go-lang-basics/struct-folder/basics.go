package main

import "fmt"

type Customer struct {
	name    string
	age     int
	balance float32
	salary  float32
}

func updateCustomerSalary(old_customer Customer, hike float32) Customer {
	old_customer.salary = old_customer.salary * (100.0 + hike) / 100
	return old_customer
}

func main() {
	var deepak Customer
	deepak.name = "Doddaka Sai Phani Deepak"
	deepak.age = 22
	deepak.balance = 100000.00
	deepak.salary = 72034.56
	fmt.Println(deepak)

	deepak = updateCustomerSalary(deepak, 25.0)

	fmt.Println(deepak)
}
