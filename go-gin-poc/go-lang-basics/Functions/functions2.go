package main

import "fmt"

func updateAccountBalance(cname string, current_balance float32, withdrawl_amount float32) {
	current_balance = current_balance - withdrawl_amount
	fmt.Printf("Dear %v , your transaction of %v is successfull\n", cname, withdrawl_amount)
	fmt.Printf("your current balance is %v\n", current_balance)
}
func main() {
	updateAccountBalance("Phani Deepak", 71843.340, 13159)
	updateAccountBalance("Punju Anna", 50000, 14903.34)
}
