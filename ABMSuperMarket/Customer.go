package main

import(
	"fmt"
)

// Customer is a person buying food
type customer struct {
	name string
	patience   int
	isAntiMask bool
	items      int
	queNumber  int
}

func(c *customer)addCustomer (setItems int, setQueNumber int) *customer {
	temp := customer{items: setItems}
	temp.queNumber = setQueNumber

	return &temp
}

func main() {
	var customerSlice = make([]customer, 20,20)

	customerSlice[0].name = "John"
	/*var newCustomer = customer{name: "John"}
	newCustomer.items = 5
	newCustomer.queNumber = 3
	customerSlice=append(newCustomer)*/
	//customerSlice.append(addCustomer(5,3))
	//customerSlice.addCustomer(5,3)

	fmt.Println("customerSlice", customerSlice[0])

}
