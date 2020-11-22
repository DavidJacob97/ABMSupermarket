package main

import (
   "fmt"
   "time"
)

// Customer is a person buying food
type customer struct {
	items      int
	queNumber  int
}

func addCustomer (setItems int, setQueNumber int) *customer {
	temp := customer{items: setItems}
	temp.queNumber = setQueNumber

	return temp
}

func main() {
	customerSlice := make([]customer, 20)


}
