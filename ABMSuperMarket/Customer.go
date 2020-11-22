package main

import (
   "fmt"
   "time"
)

// Customer is a person buying food
type Customer struct {
	name string
	patience   int
	isAntiMask bool
	items      int
	queueNumber  int
}

func addCustomer (setItems int, setQueNumber int) *customer {
	temp := customer{items: setItems}
	temp.queNumber = setQueNumber

	return temp
}

func main() {
	customerSlice := make([]customer, 20)


}
